package action

// List is a slice of action types
type List []Type

// Append allows two action lists to be concatinated
func (l List) Append(more List) List {
	for _, item := range more {
		if l.Includes(item) {
			continue
		}
		l = append(l, item)
	}
	return l
}

// Includes indicates whether the action is in the list
func (l List) Includes(actionType Type) bool {
	for _, a := range l {
		if a == actionType {
			return true
		}
	}
	return false
}
