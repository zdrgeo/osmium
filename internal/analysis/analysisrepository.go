package analysis

type AnalysisRepository interface {
	Add(name string, analysis *Analysis)
	Set(name string, analysis *Analysis)
	Remove(name string)

	Get(name string) *Analysis
}
