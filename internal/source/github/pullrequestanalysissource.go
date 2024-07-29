package github

import (
	"context"
	"fmt"
	"log"

	"github.com/zdrgeo/osmium/internal/analysis"

	"github.com/shurcooL/githubv4"
)

type PullRequestAnalysisSource struct {
	client          *githubv4.Client
	repositoryOwner string
	repositoryName  string
}

func NewPullRequestAnalysisSource(client *githubv4.Client, repositoryOwner, repositoryName string) *PullRequestAnalysisSource {
	return &PullRequestAnalysisSource{client: client, repositoryOwner: repositoryOwner, repositoryName: repositoryName}
}

func (source *PullRequestAnalysisSource) Query() *analysis.Analysis {
	type PageInfo struct {
		HasPreviousPage bool
		HasNextPage     bool
		StartCursor     string
		EndCursor       string
	}

	type PullRequestChangedFile struct {
		Path string
	}

	type PullRequest struct {
		Number int
		Title  string
		URL    string
		Files  struct {
			Nodes    []PullRequestChangedFile
			PageInfo PageInfo
		} `graphql:"files(first: 100 after: $endCursor)"`
	}

	var query struct {
		Repository struct {
			Name         string
			Description  string
			PullRequests struct {
				Nodes    []PullRequest
				PageInfo PageInfo
			} `graphql:"pullRequests(last: 100, before: $startCursor)"`
		} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
	}

	variables := map[string]any{
		"repositoryOwner": githubv4.String(source.repositoryOwner),
		"repositoryName":  githubv4.String(source.repositoryName),
		"startCursor":     githubv4.String(""),
		"endCursor":       githubv4.String(""),
	}

	err := source.client.Query(context.Background(), &query, variables)

	if err != nil {
		log.Fatal(err)
	}

	modules := map[string]*analysis.Module{}
	spans := map[string]*analysis.Span{}
	changes := map[string]*analysis.Change{}
	nodes := map[string]*analysis.Node{}

	for _, pullRequest := range query.Repository.PullRequests.Nodes {
		changeName := fmt.Sprint(pullRequest.Number)

		changes[changeName] = &analysis.Change{Name: changeName}

		for _, nodeFile := range pullRequest.Files.Nodes {
			if node, ok := nodes[nodeFile.Path]; ok {
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

				nodes[nodeFile.Path] = &analysis.Node{Edges: edges}
			}
		}
	}

	span := &analysis.Span{Changes: changes, Nodes: nodes}

	spans[span.Name] = span

	return &analysis.Analysis{Modules: modules, Spans: spans}
}
