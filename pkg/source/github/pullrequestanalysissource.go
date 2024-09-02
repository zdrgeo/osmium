package github

import (
	"fmt"
	"log"
	"time"

	"github.com/zdrgeo/osmium/pkg/analysis"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type PullRequestAnalysisSource struct {
	client *api.GraphQLClient
}

func NewPullRequestAnalysisSource(client *api.GraphQLClient) *PullRequestAnalysisSource {
	return &PullRequestAnalysisSource{client: client}
}

func (source *PullRequestAnalysisSource) Query(repositoryOwner, repositoryName string) (*analysis.Analysis, error) {
	pullRequests, err := source.getPullRequests(repositoryOwner, repositoryName)

	if err != nil {
		log.Panic(err)
	}

	modules := map[string]*analysis.Module{}
	spans := map[string]*analysis.Span{}
	changes := map[string]*analysis.Change{}
	nodes := map[string]*analysis.Node{}

	for _, pullRequest := range pullRequests {
		var files []pullRequestChangedFile

		if !pullRequest.Files.PageInfo.HasNextPage {
			files = pullRequest.Files.Nodes
		} else {
			files, err = source.getFiles(pullRequest.ID)

			if err != nil {
				log.Panic(err)
			}
		}

		changeName := fmt.Sprint(pullRequest.Number)

		changes[changeName] = &analysis.Change{Name: changeName}

		for _, nodeFile := range files {
			if node, ok := nodes[nodeFile.Path]; ok {
				for _, edgeFile := range files {
					if edge, ok := node.Edges[edgeFile.Path]; ok {
						edge.ChangeNames = append(edge.ChangeNames, changeName)
					} else {
						node.Edges[edgeFile.Path] = &analysis.Edge{ChangeNames: []string{changeName}}
					}
				}
			} else {
				edges := map[string]*analysis.Edge{}

				for _, edgeFile := range files {
					edges[edgeFile.Path] = &analysis.Edge{ChangeNames: []string{changeName}}
				}

				nodes[nodeFile.Path] = &analysis.Node{Edges: edges}
			}
		}
	}

	span := &analysis.Span{Changes: changes, Nodes: nodes}

	spans[span.Name] = span

	return &analysis.Analysis{Modules: modules, Spans: spans}, nil
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

	for {
		if err := source.client.Query("", &query, variables); err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, query.Repository.PullRequests.Nodes...)

		if !query.Repository.PullRequests.PageInfo.HasNextPage {
			break
		}

		variables["pullRequestsCursor"] = graphql.String(query.Repository.PullRequests.PageInfo.EndCursor)
	}

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
