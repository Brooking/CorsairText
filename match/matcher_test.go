package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContext struct {
	Token string
}

func TestMatcher(t *testing.T) {
	testCases := []struct {
		name      string
		matchCase bool
		wordList  []MatchEntry
		inputWord string
		assert    func([]MatchEntry)
	}{
		{
			name:      "simple success",
			wordList:  []MatchEntry{{Word: "a", Context: &testContext{Token: "token"}}},
			inputWord: "a",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "a", actual[0].Word)
				assert.IsType(t, &testContext{}, actual[0].Context)
				context, _ := actual[0].Context.(*testContext)
				assert.Equal(t, "token", context.Token)
			},
		},
		{
			name:      "fail on case",
			matchCase: true,
			wordList:  []MatchEntry{{Word: "a"}},
			inputWord: "A",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:      "simple insensitive success, stored capital",
			wordList:  []MatchEntry{{Word: "A"}},
			inputWord: "a",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "A", actual[0].Word)
			},
		},
		{
			name:      "simple insensitive success, stored lower",
			wordList:  []MatchEntry{{Word: "a"}},
			inputWord: "A",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "a", actual[0].Word)
			},
		},
		{
			name:      "fails gracefully with no word list",
			inputWord: "a",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:     "fails gracefully with empty input word",
			wordList: []MatchEntry{{Word: "a"}},
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 0, len(actual))
			},
		},
		{
			name:      "success on partial",
			wordList:  []MatchEntry{{Word: "specify"}},
			inputWord: "Spec",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "specify", actual[0].Word)
			},
		},
		{
			name:      "success on ambiguous",
			wordList:  []MatchEntry{{Word: "at"}, {Word: "atom"}},
			inputWord: "a",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 2, len(actual))
				assert.Equal(t, "at", actual[0].Word)
				assert.Equal(t, "atom", actual[1].Word)
			},
		},
		{
			name: "success on ambiguous, but complete word",
			wordList: []MatchEntry{
				{
					Word: "at",
					Context: &testContext{
						Token: "at token",
					},
				},
				{
					Word: "atom",
					Context: &testContext{
						Token: "atom token",
					},
				},
			},
			inputWord: "at",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 2, len(actual))

				assert.Equal(t, "at", actual[0].Word)
				assert.IsType(t, &testContext{}, actual[0].Context)
				context := actual[0].Context.(*testContext)
				assert.Equal(t, "at token", context.Token)

				assert.Equal(t, "atom", actual[1].Word)
				assert.IsType(t, &testContext{}, actual[1].Context)
				context = actual[1].Context.(*testContext)
				assert.Equal(t, "atom token", context.Token)
			},
		},
		{
			name:      "success with duplicates in dict, only returns one",
			wordList:  []MatchEntry{{Word: "same"}, {Word: "same"}},
			inputWord: "same",
			assert: func(actual []MatchEntry) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "same", actual[0].Word)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			m := NewTreeMatcher(testCase.wordList, testCase.matchCase)

			// act
			returnedWords := m.Match(testCase.inputWord)

			// assert
			testCase.assert(returnedWords)

		})
	}
}
