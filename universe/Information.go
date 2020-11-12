package universe

import (
	"corsairtext/universe/spot"
	"sort"
	"strings"
)

// Information is the interface to the universe's information
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Information interface {
	// ListLocalCommands returns a list of commands that are legal at the current location
	ListLocalCommands() map[string]interface{}

	// ListLocations returns a list of all locations
	ListLocations() []string

	// ListAdjacentLocations returns a list of all adjacent locations
	ListAdjacentLocations() []string

	// LocalLocation returns a view of the current location
	LocalLocation() *View

	// ListItems returns a list of all commodities
	ListItems() []string

	// Inventory returns a view of the ship
	Inventory() Ship

	// Map returns a map (centered on name if specified)
	Map(name string) *MapNode
}

// ListLocalCommands returns a list of commands valid at the current spot
func (u *universe) ListLocalCommands() map[string]interface{} {
	return u.current.Actions().Map()
}

// ListLocations returns a list of all locations
func (u *universe) ListLocations() []string {
	var result []string
	for _, spot := range u.index {
		result = append(result, spot.Name())
	}
	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(strings.ToLower(result[i]), strings.ToLower(result[j])) < 0
	})
	return result
}

// ListAdjacentLocations returns a list of all locations adjacent to the current spot
func (u *universe) ListAdjacentLocations() []string {
	var list []string
	for _, adjacent := range u.current.ListAdjacent() {
		list = append(list, adjacent.Name())
	}
	return list
}

// View is what you get when you look
type View struct {
	Name        string
	Description string
	Path        []string
}

// LocalLocation views the current spot
func (u *universe) LocalLocation() *View {
	return &View{
		Description: u.current.Description(),
		Name:        u.current.Name(),
		Path:        u.current.Path(),
	}
}

// ListItems returns a list of all items
func (u *universe) ListItems() []string {
	return nil
}

// ItemLot is a homogenous bundle of items
type ItemLot struct {
	Count    int
	UnitCost int
	Origin   string
}

// Ship contains your inventory
type Ship struct {
	Money        int
	ItemCapacity int
	Items        map[string]*ItemLot
}

// Load returns how much cargo is on board
func (s *Ship) Load() int {
	var result int
	for _, lot := range s.Items {
		result += lot.Count
	}
	return result
}

// Inventory returns a description of the ship
func (u *universe) Inventory() Ship {
	return *u.ship
}

// MapNode is the basis for a printable map
type MapNode struct {
	Name     string
	Parent   *MapNode
	Children []*MapNode
}

// Map returns a printable map
func (u *universe) Map(name string) *MapNode {
	root, target := mapWorker(u.root, nil, name)
	if target != nil {
		return target
	}
	return root
}

// mapWorker traverses the spot grid and produces a map
func mapWorker(root spot.Spot, parent *MapNode, name string) (*MapNode, *MapNode) {
	node := &MapNode{
		Name:   root.Name(),
		Parent: parent,
	}

	var target *MapNode
	for _, spot := range root.Children() {
		child, possibleTarget := mapWorker(spot, node, name)
		node.Children = append(node.Children, child)

		if possibleTarget != nil {
			target = possibleTarget
		}
	}

	if root.Name() == name {
		target = node
	}
	return node, target
}
