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

	// Description returns a string describing this spot
	Description() string

	// Name returns the name of the spot
	Name() string

	// Path returns a hierarchical location of this spot
	Path() string

	// ListAdjacent returns a list of adjacent spots
	ListAdjacent() []Spot

	// AddChild adds a child to this spot
	AddChild(child Spot)

	// Parent returns the parent of this spot
	Parent() Spot

	// Children lists the children of this spot
	Children() []Spot
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
		actions:     spotConstActions,
		description: init.Description,
		base:        base.NewBase(init.BaseType),
		name:        init.Name,
		parent:      init.Parent,
	}

	if s.base != nil {
		s.actions = s.actions.Append(s.base.Actions())
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
	return s.actions
}

// AddChild adds a child spot
func (s *spot) AddChild(child Spot) {
	if s == nil {
		return
	}
	s.children = append(s.children, child)
}

func (s *spot) Parent() Spot {
	return s.parent
}

func (s *spot) Children() []Spot {
	return s.children
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
	var current Spot = s
	var path string
	for {
		if current == nil {
			break
		}
		path = current.Name() + ">" + path
		if current.Parent() == nil {
			break
		}
		current = current.Parent()
	}
	return path
}

// ListAdjacent returns a list of adjacent spots
// todo, we should not do this every time (cache it somewhere)
func (s *spot) ListAdjacent() []Spot {
	var list []Spot
	if s.Parent() != nil {
		list = append(list, s.Parent())
		siblings := s.Parent().Children()
		for _, sibling := range siblings {
			if sibling == s {
				continue
			}
			list = append(list, sibling)
		}
	}
	for _, child := range s.Children() {
		list = append(list, child)
	}
	return list
}

// spotConstActions lists this spots actions
var spotConstActions = action.List{
	action.TypeGo,
	action.TypeHelp,
	action.TypeLook,
	action.TypeQuit,
}
