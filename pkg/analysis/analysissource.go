package analysis

type AnalysisSource interface {
	Query(options map[string]string) (*Analysis, error)
}
