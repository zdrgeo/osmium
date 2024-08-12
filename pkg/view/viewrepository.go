package view

type ViewRepository interface {
	Add(analysisName, name string, view *AnalysisView)
	Set(analysisName, name string, view *AnalysisView)
	Remove(analysisName, name string)

	Get(analysisName, name string) *AnalysisView
}
