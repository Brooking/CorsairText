package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcher(t *testing.T) {
	testCases := []struct {
		name      string
		matchCase bool
		wordList  []string
		inputWord string
		assert    func([]string)
	}{
		{
			name:      "simple success",
			wordList:  []string{"a"},
			inputWord: "a",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "a", actual[0])
			},
		},
		{
			name:      "fail on case",
			matchCase: true,
			wordList:  []string{"a"},
			inputWord: "A",
			assert: func(actual []string) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:      "simple insensitive success, stored capital",
			wordList:  []string{"A"},
			inputWord: "a",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "A", actual[0])
			},
		},
		{
			name:      "simple insensitive success, stored lower",
			wordList:  []string{"a"},
			inputWord: "A",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "a", actual[0])
			},
		},
		{
			name:      "fails gracefully with no word list",
			inputWord: "a",
			assert: func(actual []string) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:     "fails gracefully with empty input word",
			wordList: []string{"a"},
			assert: func(actual []string) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:      "success on partial",
			wordList:  []string{"specify"},
			inputWord: "Spec",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "specify", actual[0])
			},
		},
		{
			name:      "success on ambiguous",
			wordList:  []string{"at", "atom"},
			inputWord: "a",
			assert: func(actual []string) {
				assert.Equal(t, 2, len(actual))
				assert.Equal(t, "at", actual[0])
				assert.Equal(t, "atom", actual[1])
			},
		},
		{
			name:      "success on ambiguous, but complete word",
			wordList:  []string{"at", "atom"},
			inputWord: "at",
			assert: func(actual []string) {
				assert.Equal(t, 2, len(actual))
				assert.Equal(t, "at", actual[0])
				assert.Equal(t, "atom", actual[1])
			},
		},
		{
			name:      "success with duplicates in dict, only returns one",
			wordList:  []string{"same", "same"},
			inputWord: "same",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "same", actual[0])
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			m := NewMatcher(testCase.wordList, testCase.matchCase)

			// act
			returnedWords := m.Match(testCase.inputWord)

			// assert
			testCase.assert(returnedWords)

		})
	}
}
