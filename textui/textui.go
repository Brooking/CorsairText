package textui

import (
	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/universe"
	"corsairtext/universe/spot"
	"fmt"
)

// TextUI is the entry interface for the text ui
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type TextUI interface {
	// Run starts the text based UI
	Run()
}

// NewTextUI create a new text ui
func NewTextUI(s support.Support, u universe.Universe) TextUI {
	return &textUI{
		s: s,
		u: u,
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	s support.Support
	u universe.Universe
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	for {
		spot := t.u.WhereAmI()
		t.s.Out.Println("You are at", spot.Description())
		t.s.Out.Println(spot.Path())
		t.s.Out.Print(t.composeActions(spot))
		t.s.Out.Print("ready> ")
		text, err := t.s.In.Readln()
		if err != nil {
			return
		}
		if text == "q" {
			return
		}
		t.s.Out.Println("'" + text + "'")
	}
}

func (t *textUI) composeActions(spot spot.Spot) string {
	actionTypes := spot.Actions()
	descriptions := action.ActionDescriptions(actionTypes)

	var result string
	for _, description := range descriptions {
		result += fmt.Sprintf("%s - %s\n", description.Usage, description.Description)
	}
	result += fmt.Sprintf("(Q)uit - exit the game\n")
	return result
}
