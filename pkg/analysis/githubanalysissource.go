package analysis

import "fmt"

type GitHubAnalysisSource interface {
	Query(repositoryOwner, repositoryName string) (*Analysis, error)
}

type GitHubAnalysisProgress struct {
	PullRequestCount      int
	PullRequestTotalCount int
}

type GitHubAnalysisProgressFunc func(progress *GitHubAnalysisProgress)

func RenderGitHubAnalysisProgress(progress *GitHubAnalysisProgress) {
	fmt.Print("\x1B[F\x1B[K")
	fmt.Printf("Pull Requests %d of %d\n", progress.PullRequestCount, progress.PullRequestTotalCount)
}
