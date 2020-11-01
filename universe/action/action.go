package action

type ActionType string

const (
	Look ActionType = "Look"
)

type ActionDescription struct {
	Type        ActionType
	Description string
	Usage       string
	Shortcut    string
}

type Action struct {
	Type   ActionType
	Number int
	Target string
}
