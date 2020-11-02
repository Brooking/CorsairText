package spot

import (
	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/universe/base"
)

// Spot is a location in the universe
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Spot interface {
	// Actions returns a list of possible actions at this spot
	Actions() action.List

	// AddChild adds a child to this spot
	AddChild(child Spot)

	// Description returns a string describing this spot
	Description() string

	// Name returns the name of the spot
	Name() string

	// Path returns a hierarchical location of this spot
	Path() string
}

// Init is the Spot initializer
type Init struct {
	Support     support.Support
	Description string
	BaseType    base.Type
	Name        string
	Parent      Spot
}

// NewSpot creates a Spot
func NewSpot(init Init) Spot {
	s := &spot{
		s:           init.Support,
		actions:     spotActions,
		description: init.Description,
		base:        base.NewBase(init.BaseType),
		name:        init.Name,
		parent:      init.Parent,
	}
	return s
}

// spot implements the Spot interface
type spot struct {
	s           support.Support
	actions     action.List
	children    []Spot
	description string
	base        base.Base
	name        string
	parent      Spot
}

// Actions returns a list of possible actions at this spot
func (s *spot) Actions() action.List {
	actions := s.actions
	if s.base == nil {
		return actions
	}
	return actions.Append(s.base.Actions())
}

// AddChild adds a child spot
func (s *spot) AddChild(child Spot) {
	if s == nil {
		return
	}
	s.children = append(s.children, child)
}

// Description returns a textual description of the spot
func (s *spot) Description() string {
	if s == nil {
		return ""
	}
	return s.description
}

// Name returns the name of the spot
func (s *spot) Name() string {
	if s == nil {
		return ""
	}
	return s.name
}

// Path returns the path to the spot
func (s *spot) Path() string {
	var path string
	for {
		if s == nil {
			break
		}
		path = s.name + ">" + path
		if s.parent == nil {
			break
		}
		s = s.parent.(*spot)
	}
	return path
}

var spotActions = action.List{
	action.TypeGo,
	action.TypeHelp,
	action.TypeLook,
}
