package startgg

type Entrant struct {
	Id   int
	Name string
}

type Slot struct {
	Id      string
	Entrant Entrant
}

type Node struct {
	Id            int    `json:"id"`
	LPlacement    int    `json:"lPlacement"`
	FullRoundText string `json:"fullRoundText"`
	DisplayScore  string `json:"displayScore"`
}

type Set struct {
	Id       int
	Nodes    []Node
	PageInfo struct {
		Total int
	}
}

type Event struct {
	Id   int
	Name string `json:"name"`
	Sets Set
}

type Tournament struct {
	Id   int
	Name string
}
