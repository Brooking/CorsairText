package regex

import (
	"regexp"
	"sort"

	"github.com/pkg/errors"
)

// Matcher is a concrete matcher based on regex
type Matcher struct {
	matchCase bool
	list      map[string]interface{}
}

// NewRegexMatcher is the constructor for a regex based matcher
func NewRegexMatcher(wordList []string, matchCase bool) *Matcher {
	matcher := &Matcher{
		matchCase: matchCase,
		list:      make(map[string]interface{}),
	}
	for _, word := range wordList {
		matcher.Add(word)
	}
	return matcher
}

// Add adds a word to the matcher
func (m *Matcher) Add(word string) {
	_, exists := m.list[word]
	if exists {
		// todo what about duplicates with instance data
		return
	}
	m.list[word] = nil
}

// Match takes a string and finds its matches
func (m *Matcher) Match(target string) []string {
	if target == "" {
		return []string{}
	}

	var (
		result sort.StringSlice
		regex  string
	)

	switch m.matchCase {
	case true:
		regex = `^` + target
	default:
		regex = `(?i)^` + target
	}

	for word := range m.list {
		matched, err := regexp.MatchString(regex, word)
		if err != nil {
			panic(errors.Wrapf(err, "regexmatcher failed on MatchString with %v", target))
		}
		if matched {
			result = append(result, word)
		}
	}

	result.Sort()
	return result
}

// PrintOrdered not implemented for regex matcher
func (m *Matcher) PrintOrdered() {}

// PrintOrdered not implemented for regex matcher
func (m *Matcher) PrintTree() {}
