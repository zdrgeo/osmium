package analysis

type AnalysisSource interface {
	Query() *Analysis
}
