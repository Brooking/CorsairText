package universe

import (
	"corsairtext/action"

	"github.com/pkg/errors"
)

// Actions returns the legal actions at the current spot
func (u *universe) Actions() action.List {
	return u.current.Actions()
}

// Act requests that an action be taken
func (u *universe) Act(actionRequest action.Request) (interface{}, error) {
	switch actionRequest.Type {
	case action.TypeLook:
		return u.doLook(actionRequest.Parameters), nil
	}
	return nil, errors.Errorf("unknown action type %v", actionRequest.Type)
}

func (u *universe) doLook([]interface{}) action.LookResponse {
	return action.LookResponse{
		Description: u.current.Description(),
		Path:        u.current.Path(),
		Name:        u.current.Name(),
	}
}
