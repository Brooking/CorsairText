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
		assert    func(string)
	}{
		{
			name:      "simple success",
			wordList:  []string{"a"},
			inputWord: "a",
			assert: func(actual string) {
				assert.Equal(t, "a", actual)
			},
		},
		{
			name:      "fail on case",
			matchCase: true,
			wordList:  []string{"a"},
			inputWord: "A",
			assert: func(actual string) {
				assert.Equal(t, "", actual)
			},
		},
		// {
		// 	name:        "simple insensitive success, stored capital",
		// 	wordList:    []string{"A"},
		// 	inputWord:   "a",
		// 	assert: func(actual string) {
		// 		assert.Equal(t, "A", actual)
		// 	},
		// },
		{
			name:      "simple insensitive success, stored lower",
			wordList:  []string{"a"},
			inputWord: "A",
			assert: func(actual string) {
				assert.Equal(t, "a", actual)
			},
		},
		{
			name:      "fails gracefully with no word list",
			inputWord: "a",
			assert: func(actual string) {
				assert.Equal(t, "", actual)
			},
		},
		{
			name:     "fails gracefully with empty input word",
			wordList: []string{"a"},
			assert: func(actual string) {
				assert.Equal(t, "", actual)
			},
		},
		{
			name:      "success on partial",
			wordList:  []string{"specify"},
			inputWord: "Spec",
			assert: func(actual string) {
				assert.Equal(t, "specify", actual)
			},
		},
		{
			name:      "fail on ambiguous",
			wordList:  []string{"at", "atom"},
			inputWord: "a",
			assert: func(actual string) {
				assert.Equal(t, "", actual)
			},
		},
		{
			name:      "success on ambiguous, but complete word",
			wordList:  []string{"at", "atom"},
			inputWord: "at",
			assert: func(actual string) {
				assert.Equal(t, "at", actual)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			m := NewMatcher(testCase.wordList, testCase.matchCase)

			// act
			returnedWord := m.Match(testCase.inputWord)

			// assert
			testCase.assert(returnedWord)

		})
	}
}
