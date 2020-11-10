package universe

// Information is the interface to the universe's information
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Information interface {
	ListCommands() []string
	ListLocalCommands() []string

	ListLocations() []string
	ListAdjacentLocations() []string
	LocalLocation() *View

	ListItems() []string

	Inventory() []string
}

// ListCommands returns a list of all possible commands
func (u *universe) ListCommands() []string {
	return nil
}

// ListLocalCommands returns a list of commands valid at the current spot
func (u *universe) ListLocalCommands() []string {
	u.current.Actions()
	return nil
}

// ListLocations returns a list of all locations
func (u *universe) ListLocations() []string {
	return nil
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

// Inventory returns a description of the ship
func (u *universe) Inventory() []string {
	return nil
}
