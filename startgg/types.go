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
	Id    int
	Slots []Slot
}

type Set struct {
	Id       int
	Nodes    []Node
	PageInfo struct {
		Total int
	}
}

type EventResponse struct {
	Event struct {
		Id   int
		Name string `json:"name"`
		Sets Set
	}
}
