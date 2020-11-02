package textui

import (
	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/universe"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
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
		actions := spot.Actions()
		t.s.Out.Println("You are at", spot.Description())
		t.s.Out.Println(spot.Path())
		t.s.Out.Print(t.composeActions(actions))
		t.s.Out.Print("ready> ")
		text, err := t.s.In.Readln()
		if err != nil {
			return
		}
		if text == "q" {
			return
		}
		action, err := t.parseAction(text, actions)
		t.s.Out.Println("'" + action.String() + "'")
	}
}

func (t *textUI) composeActions(actions action.List) string {
	var result string
	for _, description := range actions.Descriptions() {
		result += fmt.Sprintf("%s - %s\n", description.ShortUsage, description.Description)
	}
	result += fmt.Sprintf("(Q)uit - exit the game\n")
	return result
}

func (*textUI) parseAction(input string, actions action.List) (action.Type, error) {
	var matchedAction *action.Description
	words := strings.Split(input, " ")

	// find the command
	command := strings.ToLower(words[0])
	for _, description := range actions.Descriptions() {
		match, err := regexp.MatchString("\\b"+description.Regex[0]+"\\b", command)
		if err != nil {
			continue
		}
		if !match {
			continue
		}
		matchedAction = &description
		break
	}

	if matchedAction == nil {
		return action.TypeNone, errors.New("failed to match the command")
	}

	// validate the parameters
	if len(words) > len(matchedAction.Regex) {
		return matchedAction.Type, errors.New("missing parameters")
	}

	if len(words) < len(matchedAction.Regex) {
		return matchedAction.Type, errors.New("too many parameters")
	}

	for i := 1; i < len(matchedAction.Regex); i++ {
		match, err := regexp.MatchString("\\b"+matchedAction.Regex[i]+"\\b", words[i])
		if err != nil {
			return matchedAction.Type, errors.Wrapf(err, "malformed parameter #%v", i)
		}
		if !match {
			return matchedAction.Type, errors.Errorf("malformed parameter %v", i)
		}
	}

	return matchedAction.Type, nil
}
