package github

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/zdrgeo/osmium/pkg/analysis"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type PullRequestAnalysisSource struct {
	client       *api.GraphQLClient
	progressFunc analysis.GitHubAnalysisProgressFunc
}

func NewPullRequestAnalysisSource(client *api.GraphQLClient, progressFunc analysis.GitHubAnalysisProgressFunc) *PullRequestAnalysisSource {
	return &PullRequestAnalysisSource{
		client:       client,
		progressFunc: progressFunc,
	}
}

func (source *PullRequestAnalysisSource) Query(spanSize int, repositoryOwner, repositoryName string) (*analysis.Analysis, error) {
	pullRequests, err := source.getPullRequests(repositoryOwner, repositoryName)

	if err != nil {
		log.Panic(err)
	}

	spans := map[string]*analysis.Span{}

	span := &analysis.Span{Name: strconv.Itoa(0), Size: spanSize, Changes: map[string]*analysis.Change{}, Nodes: map[string]*analysis.Node{}}

	spans[span.Name] = span

	for pullRequestIndex, pullRequest := range pullRequests {
		changeName := fmt.Sprint(pullRequest.Number)

		span.Changes[changeName] = &analysis.Change{Name: changeName}

		for _, nodeFile := range pullRequest.Files.Nodes {
			if node, ok := span.Nodes[nodeFile.Path]; ok {
				for _, edgeFile := range pullRequest.Files.Nodes {
					if edge, ok := node.Edges[edgeFile.Path]; ok {
						edge.ChangeNames = append(edge.ChangeNames, changeName)
					} else {
						node.Edges[edgeFile.Path] = &analysis.Edge{ChangeNames: []string{changeName}}
					}
				}
			} else {
				edges := map[string]*analysis.Edge{}

				for _, edgeFile := range pullRequest.Files.Nodes {
					edges[edgeFile.Path] = &analysis.Edge{ChangeNames: []string{changeName}}
				}

				span.Nodes[nodeFile.Path] = &analysis.Node{Edges: edges}
			}
		}

		if spanSize != 0 && pullRequestIndex%spanSize == 0 {
			span = &analysis.Span{Name: strconv.Itoa(pullRequestIndex), Size: spanSize, Changes: map[string]*analysis.Change{}, Nodes: map[string]*analysis.Node{}}

			spans[span.Name] = span
		}
	}

	return &analysis.Analysis{Modules: map[string]*analysis.Module{}, Spans: spans}, nil
}

type (
	rateLimit struct {
		Cost      int
		Remaining int
		ResetAt   time.Time
	}

	pageInfo struct {
		EndCursor   string
		HasNextPage bool
	}

	pullRequestChangedFile struct {
		Path string
	}

	files struct {
		Nodes      []pullRequestChangedFile
		PageInfo   pageInfo
		TotalCount int
	}

	pullRequest struct {
		ID     string
		Number int
		Title  string
		URL    string
		Files  files `graphql:"files(first: 100)"`
	}

	pullRequests struct {
		Nodes      []pullRequest
		PageInfo   pageInfo
		TotalCount int
	}

	repository struct {
		Name         string
		PullRequests pullRequests `graphql:"pullRequests(first: 100, after: $pullRequestsCursor)"`
	}

	pullRequestFragment struct {
		Files files `graphql:"files(first: 100, after: $filesCursor)"`
	}

	node struct {
		PullRequest pullRequestFragment `graphql:"... on PullRequest"`
	}
)

func (source *PullRequestAnalysisSource) getPullRequests(repositoryOwner, repositoryName string) ([]pullRequest, error) {
	var query struct {
		RateLimit  rateLimit
		Repository repository `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]any{
		"owner":              graphql.String(repositoryOwner),
		"name":               graphql.String(repositoryName),
		"pullRequestsCursor": graphql.String(""),
	}

	pullRequests := []pullRequest{}

	if source.progressFunc != nil {
		progress := &analysis.GitHubAnalysisProgress{
			PullRequestCount:      0,
			PullRequestTotalCount: 0,
		}

		source.progressFunc(progress)
	}

	wg := sync.WaitGroup{}

	for {
		if err := source.client.Query("", &query, variables); err != nil {
			return nil, err
		}

		for _, pullRequest := range query.Repository.PullRequests.Nodes {
			if pullRequest.Files.PageInfo.HasNextPage {
				wg.Add(1)

				go func(pullRequestFiles *files, pullRequestID string) {
					defer wg.Done()

					files, err := source.getFiles(pullRequestID)

					if err != nil {
						return
					}

					pullRequestFiles.Nodes = files
				}(&pullRequest.Files, pullRequest.ID)
			}
		}

		pullRequests = append(pullRequests, query.Repository.PullRequests.Nodes...)

		if source.progressFunc != nil {
			progress := &analysis.GitHubAnalysisProgress{
				PullRequestCount:      len(pullRequests),
				PullRequestTotalCount: query.Repository.PullRequests.TotalCount,
			}

			source.progressFunc(progress)
		}

		if !query.Repository.PullRequests.PageInfo.HasNextPage {
			break
		}

		variables["pullRequestsCursor"] = graphql.String(query.Repository.PullRequests.PageInfo.EndCursor)
	}

	wg.Wait()

	return pullRequests, nil
}

func (source *PullRequestAnalysisSource) getFiles(pullRequestID string) ([]pullRequestChangedFile, error) {
	var query struct {
		RateLimit rateLimit
		Node      node `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id":          graphql.ID(pullRequestID),
		"filesCursor": graphql.String(""),
	}

	files := []pullRequestChangedFile{}

	for {
		if err := source.client.Query("", &query, variables); err != nil {
			return nil, err
		}

		files = append(files, query.Node.PullRequest.Files.Nodes...)

		if !query.Node.PullRequest.Files.PageInfo.HasNextPage {
			break
		}

		variables["filesCursor"] = graphql.String(query.Node.PullRequest.Files.PageInfo.EndCursor)
	}

	return files, nil
}
