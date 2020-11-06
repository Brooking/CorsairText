package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLetter(t *testing.T) {
	testCases := []struct {
		name   string
		letter string
		word   string
		final  bool
		root   *node
		assert func(*node)
	}{
		{
			name:   "add single letter to empty",
			letter: "m",
			word:   "ma",
			root:   nil,
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "m", root.Letter, "root.Letter")
				assert.Equal(t, "ma", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")
			},
		},
		{
			name:   "add single lesser letter to single",
			letter: "m",
			word:   "ma",
			root: &node{
				Letter: "n",
				Word:   "no",
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.NotNil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "n", root.Letter, "root.Letter")
				assert.Equal(t, "no", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")

				assert.Nil(t, root.Less.Less, "root.Less.Less")
				assert.Nil(t, root.Less.More, "root.Less.More")
				assert.Nil(t, root.Less.Next, "root.Less.Next")
				assert.Equal(t, "m", root.Less.Letter, "root.Less.Letter")
				assert.Equal(t, "ma", root.Less.Word, "root.Less.Word")
				assert.Equal(t, false, root.Less.FullWord, "root.Less.FullWord")
			},
		},
		{
			name:   "add single greater letter to single",
			letter: "p",
			word:   "pa",
			root: &node{
				Letter: "n",
				Word:   "no",
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.NotNil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "n", root.Letter, "root.Letter")
				assert.Equal(t, "no", root.Word, "root.Sub")
				assert.Equal(t, false, root.FullWord, "root.FullWord")

				assert.Nil(t, root.More.Less, "root.More.Less")
				assert.Nil(t, root.More.More, "root.More.More")
				assert.Nil(t, root.More.Next, "root.More.Next")
				assert.Equal(t, "p", root.More.Letter, "root.More.Letter")
				assert.Equal(t, "pa", root.More.Word, "root.More.Word")
				assert.Equal(t, false, root.More.FullWord, "root.Less.FullWord")
			},
		},
		{
			name:   "add single letter that exists",
			letter: "m",
			word:   "my",
			root: &node{
				Letter: "m",
				Word:   "ma",
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "m", root.Letter, "root.Letter")
				assert.Equal(t, "", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")
			},
		},
		{
			name:   "add single letter word",
			letter: "i",
			word:   "i",
			final:  true,
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "i", root.Letter, "root.Letter")
				assert.Equal(t, "i", root.Word, "root.Word")
				assert.Equal(t, true, root.FullWord, "root.FullWord")
			},
		},
		{
			name:   "add word to single letter word",
			letter: "i",
			word:   "it",
			final:  false,
			root: &node{
				Letter:   "i",
				Word:     "i",
				FullWord: true,
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "i", root.Letter, "root.Letter")
				assert.Equal(t, "i", root.Word, "root.Word")
				assert.Equal(t, true, root.FullWord, "root.FullWord")
			},
		},
		{
			name:   "add single letter word to existing word",
			letter: "i",
			word:   "i",
			final:  true,
			root: &node{
				Letter:   "i",
				Word:     "it",
				FullWord: false,
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "i", root.Letter, "root.Letter")
				assert.Equal(t, "i", root.Word, "root.Word")
				assert.Equal(t, true, root.FullWord, "root.FullWord")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root = testCase.root

			// act
			addLetter(testCase.letter, testCase.word, testCase.final, root, &root)

			// assert
			testCase.assert(root)

		})
	}
}

func TestAddWord(t *testing.T) {
	testCases := []struct {
		name   string
		word   string
		root   *node
		assert func(*node)
	}{
		{
			name: "add word to empty",
			word: "yes",
			root: nil,
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.NotNil(t, root.Next, "root.Next")
				assert.Equal(t, "y", root.Letter, "root.Letter")
				assert.Equal(t, "yes", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")

				assert.Nil(t, root.Next.Less, "root.Next.Less")
				assert.Nil(t, root.Next.More, "root.Next.More")
				assert.NotNil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "e", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, "yes", root.Next.Word, "root.Next.Word")
				assert.Equal(t, false, root.Next.FullWord, "root.Next.FullWord")

				assert.Nil(t, root.Next.Next.Less, "root.Next.Next.Less")
				assert.Nil(t, root.Next.Next.More, "root.Next.Next.More")
				assert.Nil(t, root.Next.Next.Next, "root.Next.Next.Next")
				assert.Equal(t, "s", root.Next.Next.Letter, "root.Next.Next.Letter")
				assert.Equal(t, "yes", root.Next.Next.Word, "root.Next.Next.Word")
				assert.Equal(t, true, root.Next.Next.FullWord, "root.Next.Next.FullWord")
			},
		},
		{
			name: "add long word over short",
			word: "yes",
			root: &node{
				Letter: "y",
				Word:   "ya",
				Next: &node{
					Letter:   "a",
					Word:     "ya",
					FullWord: true,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.NotNil(t, root.Next, "root.Next")
				assert.Equal(t, "y", root.Letter, "root.Letter")
				assert.Equal(t, "", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")

				assert.Nil(t, root.Next.Less, "root.Next.Less")
				assert.NotNil(t, root.Next.More, "root.Next.More")
				assert.Nil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "a", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, "ya", root.Next.Word, "root.Next.Word")
				assert.Equal(t, true, root.Next.FullWord, "root.Next.FullWord")

				assert.Nil(t, root.Next.More.Less, "root.Next.More.Less")
				assert.Nil(t, root.Next.More.More, "root.Next.More.More")
				assert.NotNil(t, root.Next.More.Next, "root.Next.More.Next")
				assert.Equal(t, "e", root.Next.More.Letter, "root.Next.More.Letter")
				assert.Equal(t, "yes", root.Next.More.Word, "root.Next.More.Word")
				assert.Equal(t, false, root.Next.More.FullWord, "root.Next.More.FullWord")

				assert.Nil(t, root.Next.More.Next.Less, "root.Next.More.Next.Less")
				assert.Nil(t, root.Next.More.Next.More, "root.Next.More.Next.More")
				assert.Nil(t, root.Next.More.Next.Next, "root.Next.More.Next.Next")
				assert.Equal(t, "s", root.Next.More.Next.Letter, "root.Next.More.Next.Letter")
				assert.Equal(t, "yes", root.Next.More.Next.Word, "root.Next.More.Next.Word")
				assert.Equal(t, true, root.Next.More.Next.FullWord, "root.Next.More.Next.FullWord")
			},
		},
		{
			name: "add short word over long",
			word: "ya",
			root: &node{
				Letter: "y",
				Word:   "yes",
				Next: &node{
					Letter:   "e",
					Word:     "yes",
					FullWord: false,
					Next: &node{
						Letter:   "s",
						Word:     "yes",
						FullWord: true,
					},
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.NotNil(t, root.Next, "root.Next")
				assert.Equal(t, "y", root.Letter, "root.Letter")
				assert.Equal(t, "", root.Word, "root.Word")
				assert.Equal(t, false, root.FullWord, "root.FullWord")

				assert.NotNil(t, root.Next.Less, "root.Next.Less")
				assert.Nil(t, root.Next.More, "root.Next.More")
				assert.NotNil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "e", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, "yes", root.Next.Word, "root.Next.Word")
				assert.Equal(t, false, root.Next.FullWord, "root.Next.FullWord")

				assert.Nil(t, root.Next.Next.Less, "root.Next.Next.Less")
				assert.Nil(t, root.Next.Next.More, "root.Next.Next.More")
				assert.Nil(t, root.Next.Next.Next, "root.Next.Next.Next")
				assert.Equal(t, "s", root.Next.Next.Letter, "root.Next.Next.Letter")
				assert.Equal(t, "yes", root.Next.Next.Word, "root.Next.Next.Word")
				assert.Equal(t, true, root.Next.Next.FullWord, "root.Next.Next.FullWord")

				assert.Nil(t, root.Next.Less.Less, "root.Next.Less.Less")
				assert.Nil(t, root.Next.Less.More, "root.Next.Less.More")
				assert.Nil(t, root.Next.Less.Next, "root.Next.Less.Next")
				assert.Equal(t, "a", root.Next.Less.Letter, "root.Next.Less.Letter")
				assert.Equal(t, "ya", root.Next.Less.Word, "root.Next.Less.Word")
				assert.Equal(t, true, root.Next.Less.FullWord, "root.Next.More.FullWord")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root = testCase.root

			// act
			addWord(testCase.word, testCase.word, 0, root, &root)

			// assert
			testCase.assert(root)

		})
	}
}

func TestFindLetter(t *testing.T) {
	testCases := []struct {
		name    string
		nilRoot bool
		letter  string
		assert  func(*node, string)
	}{
		{
			name:    "fails gracefully on nil root",
			nilRoot: true,
			letter:  "q",
			assert: func(node *node, letter string) {
				assert.Nil(t, node, "node")
			},
		},
		{
			name:   "fails when letter is not in tree",
			letter: "q",
			assert: func(node *node, letter string) {
				assert.Nil(t, node, "node")
			},
		},
		{
			name:   "finds letter at root node",
			letter: "g",
			assert: func(node *node, letter string) {
				assert.NotNil(t, node, "node")
				assert.Equal(t, letter, node.Letter)
			},
		},
		{
			name:   "finds letter off to the left",
			letter: "l",
			assert: func(node *node, letter string) {
				assert.NotNil(t, node, "node")
				assert.Equal(t, letter, node.Letter)
			},
		},
		{
			name:   "finds letter deep to the right",
			letter: "z",
			assert: func(node *node, letter string) {
				assert.NotNil(t, node, "node")
				assert.Equal(t, letter, node.Letter)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root *node
			if !testCase.nilRoot {
				root = &node{
					Letter: "m",
					Less: &node{
						Letter: "g",
						Less: &node{
							Letter: "a",
						},
						More: &node{
							Letter: "l",
						},
					},
					More: &node{
						Letter: "t",
						Less: &node{
							Letter: "s",
						},
						More: &node{
							Letter: "u",
							More: &node{
								Letter: "v",
								More: &node{
									Letter: "z",
								},
							},
						},
					},
				}
			}

			// act
			node := findLetter(testCase.letter, root)

			// assert
			testCase.assert(node, testCase.letter)

		})
	}
}

func TestFindWord(t *testing.T) {
	testCases := []struct {
		name    string
		nilRoot bool
		word    string
		match   string
		assert  func(string, string)
	}{
		{
			name:    "fails gracefully on nil root",
			nilRoot: true,
			word:    "top",
			match:   "",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "fails when short word is not in tree",
			word:  "i",
			match: "",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "fails when long word is not in tree",
			word:  "spider",
			match: "",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "fails on found but ambiguous word",
			word:  "t",
			match: "",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "finds word with partial",
			word:  "ta",
			match: "tap",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "finds sub word",
			word:  "to",
			match: "to",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
		{
			name:  "finds super word",
			word:  "top",
			match: "top",
			assert: func(expected string, actual string) {
				assert.Equal(t, expected, actual)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root *node
			if !testCase.nilRoot {
				root = &node{
					Letter:   "t",
					Word:     "",
					FullWord: false,
					Next: &node{
						Letter:   "o",
						Word:     "to",
						FullWord: true,
						Less: &node{
							Letter:   "a",
							Word:     "tap",
							FullWord: false,
							Next: &node{
								Letter:   "p",
								Word:     "tap",
								FullWord: true,
							},
						},
						Next: &node{
							Letter:   "p",
							Word:     "top",
							FullWord: true,
						},
					},
				}
			}

			// act
			foundWord := findWord(testCase.word, root)

			// assert
			testCase.assert(testCase.match, foundWord)

		})
	}
}
