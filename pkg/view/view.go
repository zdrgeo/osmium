package view

type SpanView struct {
	Name     string
	Size     int
	Values   [][]int
	MinValue int
	MaxValue int
}

type AnalysisView struct {
	Name      string
	NodeNames []string
	SpanViews map[string]*SpanView
}
