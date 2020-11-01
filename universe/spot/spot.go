package spot

import "corsairtext/universe/action"

// Spot is a location in the universe
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Spot interface {
	// Actions returns a list of possible actions at this spot
	Actions() []action.ActionDescription

	// Act is a command to take an action
	Act(action.Action)

	// AddChild adds a child to this spot
	AddChild(child Spot)

	// Description returns a string describing this spot
	Description() string

	// Path returns a hierarchical location of this spot
	Path() string
}

// BaseType describes what sort of base is at this spot
type BaseType int

const (
	// BaseTypeNone indicates that there is no base here
	BaseTypeNone BaseType = 0

	// BaseTypeDirt indicates that there is just a spot to land here
	BaseTypeDirt BaseType = 1

	// BaseTypeFull indicates that there is a full base here
	BaseTypeFull BaseType = 2
)

// Init is the Spot initializer
type Init struct {
	Description string
	BaseType    BaseType
	Name        string
	Parent      Spot
}

// NewSpot creates a Spot
func NewSpot(init Init) Spot {
	s := &spot{
		description: init.Description,
		baseType:    init.BaseType,
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
	baseType    BaseType
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
