package spot

import "corsairtext/universe/action"

// Spot is a location in the universe
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Spot interface {
	Actions() []action.ActionDescription
	Act(action.Action)
	AddChild(child Spot)
	Description() string
	Path() string
}

// Init is the Spot initializer
type Init struct {
	Description string
	Base        bool
	Name        string
	Parent      Spot
}

// NewSpot creates a Spot
func NewSpot(init Init) Spot {
	s := &spot{
		description: init.Description,
		base:        init.Base,
		name:        init.Name,
		parent:      init.Parent,
	}
	if init.Parent != nil {
		init.Parent.AddChild(s)
	}
	return s
}

// spot implements the Spot interface
type spot struct {
	actionList  []action.ActionDescription
	children    []Spot
	description string
	base        bool
	name        string
	parent      Spot
}

// Actions returns a slice of actions for this spot
func (s *spot) Actions() []action.ActionDescription {
	return []action.ActionDescription{}
}

// Act applies the action to this spot
func (s *spot) Act(action.Action) {
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
	return s.name + ", " + s.description
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
