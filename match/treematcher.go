package match

import (
	"strings"
)

// NewTreeMatcher creates a new tree based Matcher
func NewTreeMatcher(wordList []MatchEntry, matchCase bool) Matcher {
	matcher := &treematcher{
		matchCase: matchCase,
	}
	Ingest(matcher, wordList)
	return matcher
}

// treematcher is our concrete implementation of Matcher
type treematcher struct {
	matchCase bool
	root      *node
}

// Add adds a word to the matcher's dictionary
func (m *treematcher) Add(original MatchEntry) {
	comparisonWord := original.Word
	if !m.matchCase {
		comparisonWord = strings.ToLower(original.Word)
	}
	addWord(original, comparisonWord, 0, m.root, &m.root)
}

// Match finds the stored words that the given word uniquely identifies
func (m *treematcher) Match(word string) []MatchEntry {
	if !m.matchCase {
		word = strings.ToLower(word)
	}

	return findWord(word, m.root)

}
