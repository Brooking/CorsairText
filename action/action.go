package action

// Request contains an action request from the ui
type Request struct {
	Type       Type
	Parameters []interface{}
}

// LookResponse is returned by a look action
type LookResponse struct {
	Name        string
	Description string
	Path        string
}

// HelpResponse is returned by a help action
type HelpResponse []Description
