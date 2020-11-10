package action

import "fmt"

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

func (l List) Map() map[string]interface{} {
	result := make(map[string]interface{}, 0)
	for _, entry := range l {
		_, exists := result[entry.String()]
		if exists {
			fmt.Println("internal: duplicate action", entry)
			continue
		}
		result[entry.String()] = nil
	}

	return result
}
