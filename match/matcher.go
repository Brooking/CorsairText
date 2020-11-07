package match

// Matcher is an interface that allows us to find words in a dictionary that
// a given prefix uniquely identifies .
// It is useful in matching names and commands.
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Matcher interface {
	Add(word string)
	Match(target string) []string

	PrintOrdered()
	PrintTree()
}

// NewMatcher creates a new Matcher
func NewMatcher(wordList []string, matchCase bool) Matcher {
	return NewTreeMatcher(wordList, matchCase)
}

// Ingest adds a word list to the matcher's dictionary
func Ingest(m Matcher, wordList []string) {
	for _, word := range wordList {
		m.Add(word)
	}
}
