package action

// Description gives human readable information about an action
type Description struct {
	Type        Type
	Description string
	ShortUsage  string
	Usage       string
	NameRegex   string
	Parameters  []ParameterType
}

// DescriptionTable is the complete list of action descriptions
var DescriptionTable = []Description{
	{
		Type:        TypeHelp,
		Description: "List possible commands",
		ShortUsage:  "(H)elp",
		Usage:       "(H)elp",
		NameRegex:   "h(elp)?",
		Parameters:  []ParameterType{},
	},
	{
		Type:        TypeQuit,
		Description: "Leave the game",
		ShortUsage:  "(Q)uit",
		Usage:       "(Q)uit",
		NameRegex:   "q(uit)?",
	},
	{
		Type:        TypeLook,
		Description: "Look around",
		ShortUsage:  "(L)ook",
		Usage:       "(L)ook",
		NameRegex:   "l(ook)?",
	},
	{
		Type:        TypeGo,
		Description: "Travel",
		ShortUsage:  "(G)o",
		Usage:       "(G)o (destination)",
		NameRegex:   "g(o)?",
		Parameters:  []ParameterType{ParameterTypeAny},
	},
	{
		Type:        TypeMine,
		Description: "Dig for ore",
		ShortUsage:  "(M)ine",
		Usage:       "(M)ine",
		NameRegex:   "m(ine)?",
	},
	{
		Type:        TypeBuy,
		Description: "Purchase a commodity",
		ShortUsage:  "(B)uy",
		Usage:       "(B)uy (number) (commodity)",
		NameRegex:   "b(uy)?",
		Parameters:  []ParameterType{ParameterTypeNumber, ParameterTypeAny},
	},
	{
		Type:        TypeSell,
		Description: "Sell a commodity",
		ShortUsage:  "(S)ell",
		Usage:       "(S)ell (number) (commodity)",
		NameRegex:   "s(ell)?",
		Parameters:  []ParameterType{ParameterTypeNumber, ParameterTypeAny},
	},
}

// Describe returns a complete description of a Type
func Describe(actionType Type) Description {
	for _, description := range DescriptionTable {
		if description.Type != actionType {
			continue
		}
		return description
	}
	return Description{}
}

// ParameterType describes a parameter
type ParameterType int

const (
	ParameterTypeNone   ParameterType = 0
	ParameterTypeNumber ParameterType = 1
	ParameterTypeAny    ParameterType = 2
)

// parameterToString is a mapping from ParameterType to a string describing it
var parameterTypeToString = map[ParameterType]string{
	ParameterTypeNone:   "ParameterTypeNone",
	ParameterTypeNumber: "ParameterTypeNumber",
	ParameterTypeAny:    "ParameterTypeAny",
}

// String returns a textual represention of a ParameterType
func (p ParameterType) String() string {
	s, ok := parameterTypeToString[p]
	if !ok {
		return ""
	}
	return s
}
