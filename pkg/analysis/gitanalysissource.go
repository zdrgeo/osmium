package analysis

import "fmt"

type GitAnalysisSource interface {
	Query(repositoryURL, repositoryPath string) (*Analysis, error)
}

type GitAnalysisProgress struct {
	CommitCount      int
	CommitTotalCount int
}

type GitAnalysisProgressFunc func(progress *GitAnalysisProgress)

func RenderGitAnalysisProgress(progress *GitAnalysisProgress) {
	fmt.Print("\x1B[F\x1B[K")
	fmt.Printf("Commits %d of %d\n", progress.CommitCount, progress.CommitTotalCount)
}
