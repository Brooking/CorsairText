package match

import (
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// NewRegexMatcher is a temporary constructor for a regex based matcher
func NewRegexMatcher(wordList []MatchEntry, matchCase bool) Matcher {
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
	list      map[string]interface{}
}

func (m *regexmatcher) Add(entry MatchEntry) {
	_, exists := m.list[entry.Word]
	if exists {
		// todo what about duplicates with contexts
		return
	}
	m.list[entry.Word] = entry.Context
}

func (m *regexmatcher) Match(target string) []MatchEntry {
	if target == "" {
		return nil
	}

	var (
		result []MatchEntry
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

			result = append(result, MatchEntry{
				Word:    word,
				Context: m.list[word],
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].Word, result[j].Word) < 0
	})
	return result
}

// PrintOrdered not implemented for regex matcher
func (m *regexmatcher) PrintOrdered() {}

// PrintOrdered not implemented for regex matcher
func (m *regexmatcher) PrintTree() {}
