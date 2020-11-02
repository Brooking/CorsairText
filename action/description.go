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

// table is the complete list of action descriptions
var table = map[Type]Description{
	TypeHelp: {
		Type:        TypeHelp,
		Description: "List possible commands",
		ShortUsage:  "(H)elp",
		Usage:       "(H)elp",
		NameRegex:   "h(elp)?",
		Parameters:  []ParameterType{},
	},
	TypeLook: {
		Type:        TypeLook,
		Description: "Look around",
		ShortUsage:  "(L)ook",
		Usage:       "(L)ook",
		NameRegex:   "l(ook)?",
	},
	TypeGo: {
		Type:        TypeGo,
		Description: "Travel",
		ShortUsage:  "(G)o",
		Usage:       "(G)o (destination)",
		NameRegex:   "g(o)?",
		Parameters:  []ParameterType{ParameterTypeAny},
	},
	TypeMine: {
		Type:        TypeMine,
		Description: "Dig for ore",
		ShortUsage:  "(M)ine",
		Usage:       "(M)ine",
		NameRegex:   "m(ine)?",
	},
	TypeBuy: {
		Type:        TypeBuy,
		Description: "Purchase a commodity",
		ShortUsage:  "(B)uy",
		Usage:       "(B)uy (number) (commodity)",
		NameRegex:   "b(uy)?",
		Parameters:  []ParameterType{ParameterTypeNumber, ParameterTypeAny},
	},
	TypeSell: {
		Type:        TypeSell,
		Description: "Sell a commodity",
		ShortUsage:  "(S)ell",
		Usage:       "(S)ell (number) (commodity)",
		NameRegex:   "s(ell)?",
		Parameters:  []ParameterType{ParameterTypeNumber, ParameterTypeAny},
	},
	TypeQuit: {
		Type:        TypeQuit,
		Description: "Leave the game",
		ShortUsage:  "(Q)uit",
		Usage:       "(Q)uit",
		NameRegex:   "q(uit)?",
	},
}

type ParameterType int

const (
	ParameterTypeNone   ParameterType = 0
	ParameterTypeNumber ParameterType = 1
	ParameterTypeAny    ParameterType = 2
)

var parameterTypeToString = map[ParameterType]string{
	ParameterTypeNone:   "ParameterTypeNone",
	ParameterTypeNumber: "ParameterTypeNumber",
	ParameterTypeAny:    "ParameterTypeAny",
}

func (p ParameterType) String() string {
	s, ok := parameterTypeToString[p]
	if !ok {
		return ""
	}
	return s
}
