package github

import (
	"fmt"
	// "iter"
	"log"
	"time"

	"github.com/zdrgeo/osmium/internal/analysis"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type PullRequestAnalysisSource struct {
	client          *api.GraphQLClient
	repositoryOwner string
	repositoryName  string
}

func NewPullRequestAnalysisSource(client *api.GraphQLClient, repositoryOwner, repositoryName string) *PullRequestAnalysisSource {
	return &PullRequestAnalysisSource{client: client, repositoryOwner: repositoryOwner, repositoryName: repositoryName}
}

func (source *PullRequestAnalysisSource) Query() *analysis.Analysis {
	pullRequests, err := source.pullRequests()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pull Requests count: %d\n", len(pullRequests))

	modules := map[string]*analysis.Module{}
	spans := map[string]*analysis.Span{}
	changes := map[string]*analysis.Change{}
	nodes := map[string]*analysis.Node{}

	for _, pullRequest := range pullRequests {
		files, err := source.files(pullRequest.ID)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pull Request %d files count: %d\n", pullRequest.Number, len(files))

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

	return &analysis.Analysis{Modules: modules, Spans: spans}
}

type (
	RateLimit struct {
		Cost      int
		Remaining int
		ResetAt   time.Time
	}

	PageInfo struct {
		// StartCursor     string
		EndCursor string
		// HasPreviousPage bool
		HasNextPage bool
	}

	PullRequestChangedFile struct {
		Path string
	}

	Files struct {
		Nodes    []PullRequestChangedFile
		PageInfo PageInfo
	}

	PullRequest struct {
		ID     string
		Number int
		Title  string
		URL    string
		Files  Files `graphql:"files(first: 100)"`
	}

	PullRequests struct {
		Nodes    []PullRequest
		PageInfo PageInfo
	}

	Repository struct {
		Name         string
		PullRequests PullRequests `graphql:"pullRequests(first: 100, after: $pullRequestsCursor)"`
	}

	PullRequestFragment struct {
		Files Files `graphql:"files(first: 100, after: $filesCursor)"`
	}

	Node struct {
		PullRequest PullRequestFragment `graphql:"... on PullRequest"`
	}
)

func (source *PullRequestAnalysisSource) pullRequests() ([]PullRequest, error) {
	var query struct {
		RateLimit  RateLimit
		Repository Repository `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]any{
		"owner":              graphql.String(source.repositoryOwner),
		"name":               graphql.String(source.repositoryName),
		"pullRequestsCursor": graphql.String(""),
	}

	pullRequests := []PullRequest{}

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

func (source *PullRequestAnalysisSource) files(pullRequestID string) ([]PullRequestChangedFile, error) {
	var query struct {
		RateLimit RateLimit
		Node      Node `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id":          graphql.ID(pullRequestID),
		"filesCursor": graphql.String(""),
	}

	files := []PullRequestChangedFile{}

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

/*
func (source *PullRequestAnalysisSource) pullRequests() iter.Seq[PullRequest] {
	var query struct {
		RateLimit  RateLimit
		Repository Repository `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]any{
		"owner":              graphql.String(source.repositoryOwner),
		"name":               graphql.String(source.repositoryName),
		"pullRequestsCursor": graphql.String(""),
	}

	return func(yield func(PullRequest) bool) {
		for {
			if err := source.client.Query("", &query, variables); err != nil {
				log.Fatal(err)
			}

			for _, pullRequest := range query.Repository.PullRequests.Nodes {
				if !yield(pullRequest) {
					return
				}
			}

			if !query.Repository.PullRequests.PageInfo.HasNextPage {
				break
			}

			variables["pullRequestsCursor"] = graphql.String(query.Repository.PullRequests.PageInfo.EndCursor)
		}
	}
}

func (source *PullRequestAnalysisSource) files(pullRequestID string) iter.Seq[PullRequestChangedFile] {
	var query struct {
		RateLimit RateLimit
		Node      Node `graphql:"node(id: $id)"`
	}

	variables := map[string]any{
		"id":          graphql.ID(pullRequestID),
		"filesCursor": graphql.String(""),
	}

	return func(yield func(PullRequestChangedFile) bool) {
		for {
			if err := source.client.Query("", &query, variables); err != nil {
				log.Fatal(err)
			}

			for _, file := range query.Node.PullRequest.Files.Nodes {
				if !yield(file) {
					return
				}
			}

			if !query.Node.PullRequest.Files.PageInfo.HasNextPage {
				break
			}

			variables["filesCursor"] = graphql.String(query.Node.PullRequest.Files.PageInfo.EndCursor)
		}
	}
}
*/

/*
actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
batchSize := 3
batches := make([][]int, 0, (len(actions) + batchSize - 1) / batchSize)

for batchSize < len(actions) {
    actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)
*/
