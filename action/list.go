package action

// List is a slice of action types
type List []Type

// Append allows two action lists to be concatinated
func (l List) Append(more List) List {
	for _, item := range more {
		l = append(l, item)
	}
	return l
}

// Descriptions returns a slice of descriptions
func (l List) Descriptions() []Description {
	var descriptions []Description
	for _, actionType := range l {
		description, ok := table[actionType]
		if !ok {
			continue
		}
		descriptions = append(descriptions, description)
	}
	return descriptions
}
