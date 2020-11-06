// Package e contains all of the custom errors
//
// Additionally it implements a few 'marker' interfaces whereby errors can indicate how they should
// be displayed:
//
// ShowToUserError indicates that the error is intended to be shown to the user
// ShowAllHelpError indicates that the ui should show full help after handling the error
// ShowHelpError indicates that the ui should show command specific help after handling the error
// ShowAdjacentError indicates that the ui should show an adjacency list after handling the error
package e
