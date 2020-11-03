package universe

import (
	"corsairtext/action"

	"github.com/pkg/errors"
)

// Actions returns the legal actions for the current spot
func (u *universe) Actions() action.List {
	actions := universalActions
	return actions.Append(u.current.Actions())
}

// Act requests that an action be taken
func (u *universe) Act(actionRequest action.Request) (interface{}, error) {
	if !u.Actions().Includes(actionRequest.Type) {
		return nil, errors.Errorf("action type %v not legal here", actionRequest.Type)
	}

	actionHandler := map[action.Type]func([]interface{}) (interface{}, error){
		action.TypeHelp: u.doHelp,
		action.TypeLook: u.doLook,
		action.TypeQuit: u.doNothing,
	}
	handler, ok := actionHandler[actionRequest.Type]
	if !ok {
		return nil, errors.Errorf("unknown action type %v", actionRequest.Type.String())
	}

	return handler(actionRequest.Parameters)
}

// doLook actually implements the look action
func (u *universe) doLook([]interface{}) (interface{}, error) {
	return action.LookResponse{
		Description: u.current.Description(),
		Path:        u.current.Path(),
		Name:        u.current.Name(),
	}, nil
}

// doHelp actually implements the help action
func (u *universe) doHelp([]interface{}) (interface{}, error) {
	return u.Actions().Descriptions(), nil
}

// doLook actually implements the quit action
func (u *universe) doNothing([]interface{}) (interface{}, error) {
	return nil, nil
}

// universalActions lists the universal actions
var universalActions = action.List{
	action.TypeHelp,
	action.TypeQuit,
}
