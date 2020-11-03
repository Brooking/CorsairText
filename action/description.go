package action

// Description gives human readable information about an action
type Description struct {
	Type       Type
	ShortUsage string
	Usage      string
	NameRegex  string
	Parameters []ParameterType
}

// DescriptionTable is the complete list of action descriptions
var DescriptionTable = []Description{
	{
		Type:       TypeHelp,
		Usage:      "(H)elp - List possible commands",
		NameRegex:  "h(elp)?",
		Parameters: []ParameterType{},
	},
	{
		Type:      TypeQuit,
		Usage:     "(Q)uit - Leave the game",
		NameRegex: "q(uit)?",
	},
	{
		Type:      TypeLook,
		Usage:     "(L)ook - Look around",
		NameRegex: "l(ook)?",
	},
	{
		Type:       TypeGo,
		ShortUsage: "(G)o - Travel",
		Usage:      "(G)o <destination> - Travel to destination",
		NameRegex:  "g(o)?",
		Parameters: []ParameterType{ParameterTypeAny},
	},
	{
		Type:      TypeMine,
		Usage:     "(M)ine - Dig for ore",
		NameRegex: "m(ine)?",
	},
	{
		Type:       TypeBuy,
		ShortUsage: "(B)uy - Purchase items",
		Usage:      "(B)uy <amount> <item> - Purchase specified amount of items",
		NameRegex:  "b(uy)?",
		Parameters: []ParameterType{ParameterTypeNumber, ParameterTypeAny},
	},
	{
		Type:       TypeSell,
		ShortUsage: "(S)ell - Sell items",
		Usage:      "(S)ell <amount> <item> - Sell specified amount of items",
		NameRegex:  "s(ell)?",
		Parameters: []ParameterType{ParameterTypeNumber, ParameterTypeAny},
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
