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

func NewSpot(name string, description string, inhabitable bool, parent Spot) Spot {
	s := &spot{
		description: description,
		inhabitable: inhabitable,
		name:        name,
		parent:      parent,
	}
	if parent != nil {
		parent.AddChild(s)
	}
	return s
}

type spot struct {
	actionList  []action.ActionDescription
	children    []Spot
	description string
	inhabitable bool
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

func (s *spot) Description() string {
	if s == nil {
		return ""
	}
	return s.name + ", " + s.description
}

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
