package tree

import (
	"strings"
)

// Matcher is our concrete implementation of Matcher
type Matcher struct {
	matchCase bool
	root      *node
}

// NewTreeMatcher creates a new tree based Matcher
func NewTreeMatcher(wordList []string, matchCase bool) *Matcher {
	matcher := &Matcher{
		matchCase: matchCase,
	}
	for _, word := range wordList {
		matcher.Add(word)
	}
	return matcher
}

// Add adds a word to the matcher's dictionary
func (m *Matcher) Add(originalWord string) {
	comparisonWord := originalWord
	if !m.matchCase {
		comparisonWord = strings.ToLower(originalWord)
	}
	addWord(originalWord, comparisonWord, 0, m.root, &m.root)
}

// Match finds the stored words that the given word uniquely identifies
func (m *Matcher) Match(word string) []string {
	if !m.matchCase {
		word = strings.ToLower(word)
	}

	return findWord(word, m.root)

}
