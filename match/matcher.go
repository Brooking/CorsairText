package match

import (
	"strings"
)

// Matcher is an interface that allows us to find words in a dictionary that
// a given prefix uniquely identifies .
// It is useful in matching names and commands.
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Matcher interface {
	Add(word string)
	Ingest(wordList []string)
	Match(target string) string

	PrintOrdered()
	PrintTree()
}

// NewMatcher creates a new Matcher
func NewMatcher(wordList []string, matchCase bool) Matcher {
	matcher := &matcher{
		matchCase: matchCase,
	}
	matcher.Ingest(wordList)
	return matcher
}

// matcher is our concrete implementation of Matcher
type matcher struct {
	matchCase bool
	root      *node
}

// Add adds a word to the matcher's dictionary
func (m *matcher) Add(word string) {
	if !m.matchCase {
		word = strings.ToLower(word)
	}
	addWord(word, 0, m.root, &m.root)
}

// Ingest adds a word list to the matcher's dictionary
func (m *matcher) Ingest(wordList []string) {
	for _, word := range wordList {
		m.Add(word)
	}
}

// Match finds the stored word that the given word uniquely identifies
func (m *matcher) Match(word string) string {
	if !m.matchCase {
		word = strings.ToLower(word)
	}
	return findWord(word, m.root)
}
