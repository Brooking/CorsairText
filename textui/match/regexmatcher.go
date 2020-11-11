package match

import (
	"regexp"
	"sort"

	"github.com/pkg/errors"
)

// NewRegexMatcher is a temporary constructor for a regex based matcher
func NewRegexMatcher(wordList []string, matchCase bool) Matcher {
	matcher := &regexmatcher{
		MatchCase: matchCase,
		list:      make(map[string]interface{}),
	}
	Ingest(matcher, wordList)
	return matcher
}

// regexmatcher is a concrete matcher based on regex
type regexmatcher struct {
	MatchCase bool
	//todo make into map to avoid duplicates
	list map[string]interface{}
}

func (m *regexmatcher) Add(word string) {
	_, exists := m.list[word]
	if exists {
		// todo what about duplicates with instance data
		return
	}
	m.list[word] = nil
}

func (m *regexmatcher) Match(target string) []string {
	if target == "" {
		return []string{}
	}

	var (
		result sort.StringSlice
		regex  string
	)

	switch m.MatchCase {
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
func (m *regexmatcher) PrintOrdered() {}

// PrintOrdered not implemented for regex matcher
func (m *regexmatcher) PrintTree() {}
