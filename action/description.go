package action

// Description gives human readable information about an action
type Description struct {
	Type        Type
	Description string
	ShortUsage  string
	Usage       string
	Regex       []string
}

// table is the complete list of action descriptions
var table = map[Type]Description{
	TypeHelp: {
		Type:        TypeHelp,
		Description: "List possible commands",
		ShortUsage:  "(H)elp",
		Usage:       "(H)elp",
		Regex:       []string{"h(elp)?"},
	},
	TypeLook: {
		Type:        TypeLook,
		Description: "Look around",
		ShortUsage:  "(L)ook",
		Usage:       "(L)ook",
		Regex:       []string{"l(ook)?"},
	},
	TypeGo: {
		Type:        TypeGo,
		Description: "Travel",
		ShortUsage:  "(G)o",
		Usage:       "(G)o (destination)",
		Regex:       []string{"g(o)?", ".+"},
	},
	TypeMine: {
		Type:        TypeMine,
		Description: "Dig for ore",
		ShortUsage:  "(M)ine",
		Usage:       "(M)ine",
		Regex:       []string{"m(ine)?"},
	},
	TypeBuy: {
		Type:        TypeBuy,
		Description: "Purchase a commodity",
		ShortUsage:  "(B)uy",
		Usage:       "(B)uy (number) (commodity)",
		Regex:       []string{"b(uy)?", "[[:digit:]]+", ".+"},
	},
	TypeSell: {
		Type:        TypeSell,
		Description: "Sell a commodity",
		ShortUsage:  "(S)ell",
		Usage:       "(S)ell (number) (commodity)",
		Regex:       []string{"s(ell)?", "[[:digit:]]+", ".+"},
	},
}
