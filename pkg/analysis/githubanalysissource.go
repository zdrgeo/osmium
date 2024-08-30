package analysis

type GitHubAnalysisSource interface {
	Query(repositoryOwner, repositoryName string) (*Analysis, error)
}
