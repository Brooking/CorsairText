package match

// Matcher is an interface that allows us to find words in a dictionary that
// a given prefix uniquely identifies .
// It is useful in matching names and commands.
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Matcher interface {
	Add(MatchEntry)
	Match(target string) []MatchEntry

	PrintOrdered()
	PrintTree()
}

// NewMatcher creates a new Matcher
func NewMatcher(wordList []MatchEntry, matchCase bool) Matcher {
	return NewRegexMatcher(wordList, matchCase)
}

// Ingest adds a word list to the matcher's dictionary
func Ingest(m Matcher, wordList []MatchEntry) {
	for _, entry := range wordList {
		m.Add(entry)
	}
}

type MatchEntry struct {
	Word    string
	Context interface{}
}
