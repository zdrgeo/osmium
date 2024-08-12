package analysis

type Module struct {
	Name      string
	NodeNames []string
}

type Change struct {
	Name string
}

type Edge struct {
	NodeName    string
	ChangeNames []string
}

type Node struct {
	Name  string
	Edges map[string]*Edge
}

type Span struct {
	Name    string
	Size    int
	Changes map[string]*Change
	Nodes   map[string]*Node
}

type Analysis struct {
	Name    string
	Modules map[string]*Module
	Spans   map[string]*Span
}
