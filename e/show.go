package e

// ShowAllHelpError is the interface that marks an error for full help
type ShowAllHelpError interface {
	IsShowAllHelp()
}

// IsShowAllHelpError indicates whether the error should be followed by a full help screen
func IsShowAllHelpError(err error) bool {
	_, ok := err.(ShowAllHelpError)
	return ok
}

// ShowHelpError is the interface that marks an error for command specific help
type ShowHelpError interface {
	GetCommand() string
}

// IsShowHelpError indicates whether the error should be followed by a command specific help screen
func IsShowHelpError(err error) bool {
	_, ok := err.(ShowHelpError)
	return ok
}

// GetCommandForHelp gets the hidden command within the error to facilitate showing the correct help
func GetCommandForHelp(err error) string {
	var command string
	showHelpError, ok := err.(ShowHelpError)
	if !ok {
		return command
	}
	return showHelpError.GetCommand()
}

// ShowToUserError is the interface that marks an error for display to the user
type ShowToUserError interface {
	IsShowToUser()
}

// IsShowToUserError indicates whether the error should be shown to the user
func IsShowToUserError(err error) bool {
	_, ok := err.(ShowToUserError)
	return ok
}

// ShowAdjacencyError is the interface that marks an error for adjacency help
type ShowAdjacencyError interface {
	IsShowAdjacent()
}

// IsShowAdjacencyError indicates whether the error should be followed by adjacency help screen
func IsShowAdjacencyError(err error) bool {
	_, ok := err.(ShowAdjacencyError)
	return ok
}
