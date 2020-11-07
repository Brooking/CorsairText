package match

import (
	"strings"
)

// NewTreeMatcher creates a new tree based Matcher
func NewTreeMatcher(wordList []string, matchCase bool) Matcher {
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
func (m *treematcher) Add(originalWord string) {
	comparisonWord := originalWord
	if !m.matchCase {
		comparisonWord = strings.ToLower(originalWord)
	}
	addWord(originalWord, comparisonWord, 0, m.root, &m.root)
}

// Match finds the stored words that the given word uniquely identifies
func (m *treematcher) Match(word string) []string {
	if !m.matchCase {
		word = strings.ToLower(word)
	}

	return findWord(word, m.root)

}
