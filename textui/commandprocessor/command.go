package commandprocessor

import (
	"corsairtext/e"
	"corsairtext/support"
	"corsairtext/textui/commandprocessor/match"
	"corsairtext/universe"
	"strings"

	"github.com/pkg/errors"
)

// CommandProcessor is the interface that the text ui uses to get work done
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type CommandProcessor interface {
	CommandMatcher() match.Matcher

	ShowAdjacency() error
	ShowAllHelp() error
	ShowHelp(command string) error
	ShowLook() error

	Obey(commandLine string) error
}

// NewCommandProcessor creates a new command handler
func NewCommandProcessor(s *support.SupportStruct, a universe.Action, i universe.Information) CommandProcessor {
	cp := &commandProcessor{
		s:            s,
		a:            a,
		i:            i,
		descriptions: newDescriptions(),
	}
	cp.init()
	return cp
}

// commandProcessor is our concrete implimetation of CommandProcessor
type commandProcessor struct {
	s               *support.SupportStruct
	a               universe.Action
	i               universe.Information
	commandMatcher  match.Matcher
	locationMatcher match.Matcher
	descriptions    map[string]*commandDescription
}

func (cp *commandProcessor) CommandMatcher() match.Matcher {
	return cp.commandMatcher
}

// ShowAdjacency shows the names of all adjacent spots
func (cp *commandProcessor) ShowAdjacency() error {
	adjacency := cp.i.ListAdjacentLocations()
	for _, neighbor := range adjacency {
		cp.s.Out.Println(neighbor)
	}
	return nil
}

// ShowAllHelp implements the all help command
func (cp *commandProcessor) ShowAllHelp() error {
	var commands []*commandDescription
	var maxLen int
	legalActions := cp.i.ListLocalActions()
	for _, command := range commandHelpOrder {
		description, ok := cp.descriptions[command]
		if !ok {
			return errors.Errorf("internal: unknown command %v", command)
		}
		if description.Action {
			_, exist := legalActions[command]
			if !exist {
				continue
			}
		}
		commands = append(commands, description)
		if len(description.ShortName) > maxLen {
			maxLen = len(description.ShortName)
		}
	}

	for _, description := range commands {
		name := description.ShortName + strings.Repeat(" ", maxLen-len(description.ShortName))
		cp.s.Out.Println(strings.Join([]string{name, " - ", description.ShortUsage}, ""))
	}
	return nil
}

// ShowHelp implements the specific help command
func (cp *commandProcessor) ShowHelp(command string) error {
	commands := cp.commandMatcher.Match(command)
	switch len(commands) {
	case 0:
		return e.NewUnknownCommandError(command)
	default:
		return e.NewUnknownCommandError(command)
	case 1:
		description, ok := cp.descriptions[command]
		if !ok {
			return e.NewUnknownCommandError(command)
		}

		name := description.LongName
		if name == "" {
			name = description.ShortName
		}
		usage := description.LongUsage
		if usage == "" {
			usage = description.ShortUsage
		}
		cp.s.Out.Println(strings.Join([]string{name, " - ", usage}, ""))
		return nil
	}
}

// ShowLook implements a local look around
func (cp *commandProcessor) ShowLook() error {
	location := cp.i.LocalLocation()
	cp.s.Out.Println(strings.Join([]string{"You are at ", location.Name, ", ", location.Description, "."}, ""))

	var path string
	for _, spot := range location.Path {
		path = path + spot + "/"
	}
	cp.s.Out.Println(path)
	return nil
}

// Obey handles a user's command
func (cp *commandProcessor) Obey(commandString string) error {
	var (
		words       []string = strings.Split(commandString, " ")
		command     string   = strings.ToLower(words[0])
		parameters  []string = words[1:]
		description *commandDescription
		err         error
	)

	// find the command
	description, err = cp.parseCommand(command)
	if err != nil {
		return err
	}

	// handle the command
	return description.Handler(cp, parameters)
}

func (cp *commandProcessor) init() {
	cp.commandMatcher = NewCommandMatcher(cp)
	if cp.i != nil {
		cp.locationMatcher = match.NewMatcher(cp.i.ListLocations(), false)
	}
}
