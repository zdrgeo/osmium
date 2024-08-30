package analysis

type GitAnalysisSource interface {
	Query(repositoryURL, repositoryPath string) (*Analysis, error)
}
