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
				assert.Equal(t, 1, len(root.Words), "len root.Words")
				assert.Equal(t, "ma", root.Words[0], "root.Words[0]")
			},
		},
		{
			name:   "add single lesser letter to single",
			letter: "m",
			word:   "ma",
			root: &node{
				Letter: "n",
				Words:  []string{"no"},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.NotNil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "n", root.Letter, "root.Letter")
				assert.Equal(t, 1, len(root.Words), "len root.Words")
				assert.Equal(t, "no", root.Words[0], "root.Words[0]")

				assert.Nil(t, root.Less.Less, "root.Less.Less")
				assert.Nil(t, root.Less.More, "root.Less.More")
				assert.Nil(t, root.Less.Next, "root.Less.Next")
				assert.Equal(t, "m", root.Less.Letter, "root.Less.Letter")
				assert.Equal(t, 1, len(root.Less.Words), "len root.Words")
				assert.Equal(t, "ma", root.Less.Words[0], "root.Less.Words[0]")
			},
		},
		{
			name:   "add single greater letter to single",
			letter: "p",
			word:   "pa",
			root: &node{
				Letter: "n",
				Words:  []string{"no"},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.NotNil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "n", root.Letter, "root.Letter")
				assert.Equal(t, 1, len(root.Words), "len root.Words")
				assert.Equal(t, "no", root.Words[0], "root.Sub")

				assert.Nil(t, root.More.Less, "root.More.Less")
				assert.Nil(t, root.More.More, "root.More.More")
				assert.Nil(t, root.More.Next, "root.More.Next")
				assert.Equal(t, "p", root.More.Letter, "root.More.Letter")
				assert.Equal(t, 1, len(root.More.Words), "len root.Words")
				assert.Equal(t, "pa", root.More.Words[0], "root.More.Words[0]")
			},
		},
		{
			name:   "add single letter that exists",
			letter: "m",
			word:   "my",
			root: &node{
				Letter: "m",
				Words:  []string{"ma"},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "m", root.Letter, "root.Letter")
				assert.Equal(t, 2, len(root.Words), "len root.Words")
				assert.Equal(t, "ma", root.Words[0], "root.Words[0]")
				assert.Equal(t, "my", root.Words[1], "root.Words[1]")
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
				assert.Equal(t, 1, len(root.Words), "len root.Words")
				assert.Equal(t, "i", root.Words[0], "root.Words[0]")
			},
		},
		{
			name:   "add word to single letter word",
			letter: "i",
			word:   "it",
			final:  false,
			root: &node{
				Letter: "i",
				Words:  []string{"i"},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "i", root.Letter, "root.Letter")
				assert.Equal(t, 2, len(root.Words), "len root.Words")
				assert.Equal(t, "i", root.Words[0], "root.Words[0]")
				assert.Equal(t, "it", root.Words[1], "root.Words[1]")
			},
		},
		{
			name:   "add single letter word to existing word",
			letter: "i",
			word:   "i",
			final:  true,
			root: &node{
				Letter: "i",
				Words:  []string{"it"},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.Nil(t, root.Next, "root.Next")
				assert.Equal(t, "i", root.Letter, "root.Letter")
				assert.Equal(t, 2, len(root.Words), "len root.Words")
				assert.Equal(t, "it", root.Words[0], "root.Words[0]")
				assert.Equal(t, "i", root.Words[1], "root.Words[1]")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root = testCase.root

			// act
			addLetter(testCase.letter, testCase.word, root, &root)

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
				assert.Equal(t, 1, len(root.Words), "len root.Words")
				assert.Equal(t, "yes", root.Words[0], "root.Words[0]")

				assert.Nil(t, root.Next.Less, "root.Next.Less")
				assert.Nil(t, root.Next.More, "root.Next.More")
				assert.NotNil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "e", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, 1, len(root.Next.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.Words[0], "root.Next.Words[0]")

				assert.Nil(t, root.Next.Next.Less, "root.Next.Next.Less")
				assert.Nil(t, root.Next.Next.More, "root.Next.Next.More")
				assert.Nil(t, root.Next.Next.Next, "root.Next.Next.Next")
				assert.Equal(t, "s", root.Next.Next.Letter, "root.Next.Next.Letter")
				assert.Equal(t, 1, len(root.Next.Next.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.Next.Words[0], "root.Next.Next.Words[0]")
			},
		},
		{
			name: "add long word over short",
			word: "yes",
			root: &node{
				Letter: "y",
				Words:  []string{"ya"},
				Next: &node{
					Letter: "a",
					Words:  []string{"ya"},
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.NotNil(t, root.Next, "root.Next")
				assert.Equal(t, "y", root.Letter, "root.Letter")
				assert.Equal(t, 2, len(root.Words), "len root.Words")
				assert.Equal(t, "ya", root.Words[0], "root.Words[0]")
				assert.Equal(t, "yes", root.Words[1], "root.Words[1]")

				assert.Nil(t, root.Next.Less, "root.Next.Less")
				assert.NotNil(t, root.Next.More, "root.Next.More")
				assert.Nil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "a", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, 1, len(root.Next.Words), "len root.Words")
				assert.Equal(t, "ya", root.Next.Words[0], "root.Next.Words[0]")

				assert.Nil(t, root.Next.More.Less, "root.Next.More.Less")
				assert.Nil(t, root.Next.More.More, "root.Next.More.More")
				assert.NotNil(t, root.Next.More.Next, "root.Next.More.Next")
				assert.Equal(t, "e", root.Next.More.Letter, "root.Next.More.Letter")
				assert.Equal(t, 1, len(root.Next.More.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.More.Words[0], "root.Next.More.Words[0]")

				assert.Nil(t, root.Next.More.Next.Less, "root.Next.More.Next.Less")
				assert.Nil(t, root.Next.More.Next.More, "root.Next.More.Next.More")
				assert.Nil(t, root.Next.More.Next.Next, "root.Next.More.Next.Next")
				assert.Equal(t, "s", root.Next.More.Next.Letter, "root.Next.More.Next.Letter")
				assert.Equal(t, 1, len(root.Next.More.Next.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.More.Next.Words[0], "root.Next.More.Next.Words[0]")
			},
		},
		{
			name: "add short word over long",
			word: "ya",
			root: &node{
				Letter: "y",
				Words:  []string{"yes"},
				Next: &node{
					Letter: "e",
					Words:  []string{"yes"},
					Next: &node{
						Letter: "s",
						Words:  []string{"yes"},
					},
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root, "root")
				assert.Nil(t, root.Less, "root.Less")
				assert.Nil(t, root.More, "root.More")
				assert.NotNil(t, root.Next, "root.Next")
				assert.Equal(t, "y", root.Letter, "root.Letter")
				assert.Equal(t, 2, len(root.Words), "len root.Words")
				assert.Equal(t, "yes", root.Words[0], "root.Words[0]")
				assert.Equal(t, "ya", root.Words[1], "root.Words[1]")

				assert.NotNil(t, root.Next.Less, "root.Next.Less")
				assert.Nil(t, root.Next.More, "root.Next.More")
				assert.NotNil(t, root.Next.Next, "root.Next.Next")
				assert.Equal(t, "e", root.Next.Letter, "root.Next.Letter")
				assert.Equal(t, 1, len(root.Next.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.Words[0], "root.Next.Words[0]")

				assert.Nil(t, root.Next.Next.Less, "root.Next.Next.Less")
				assert.Nil(t, root.Next.Next.More, "root.Next.Next.More")
				assert.Nil(t, root.Next.Next.Next, "root.Next.Next.Next")
				assert.Equal(t, "s", root.Next.Next.Letter, "root.Next.Next.Letter")
				assert.Equal(t, 1, len(root.Next.Next.Words), "len root.Words")
				assert.Equal(t, "yes", root.Next.Next.Words[0], "root.Next.Next.Words[0]")

				assert.Nil(t, root.Next.Less.Less, "root.Next.Less.Less")
				assert.Nil(t, root.Next.Less.More, "root.Next.Less.More")
				assert.Nil(t, root.Next.Less.Next, "root.Next.Less.Next")
				assert.Equal(t, "a", root.Next.Less.Letter, "root.Next.Less.Letter")
				assert.Equal(t, 1, len(root.Next.Less.Words), "len root.Words")
				assert.Equal(t, "ya", root.Next.Less.Words[0], "root.Next.Less.Words[0]")
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
		assert  func([]string)
	}{
		{
			name:    "fails gracefully on nil root",
			nilRoot: true,
			word:    "top",
			assert: func(actual []string) {
				assert.Equal(t, []string{}, actual)
			},
		},
		{
			name: "fails when short word is not in tree",
			word: "i",
			assert: func(actual []string) {
				assert.Equal(t, []string{}, actual)
			},
		},
		{
			name: "fails when long word is not in tree",
			word: "spider",
			assert: func(actual []string) {
				assert.Equal(t, []string{}, actual)
			},
		},
		{
			name: "fails on found but ambiguous word",
			word: "t",
			assert: func(actual []string) {
				assert.Equal(t, 3, len(actual))
				assert.Equal(t, "to", actual[0])
				assert.Equal(t, "top", actual[1])
				assert.Equal(t, "tap", actual[2])
			},
		},
		{
			name: "finds word with partial",
			word: "ta",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "tap", actual[0])
			},
		},
		{
			name: "finds sub word",
			word: "to",
			assert: func(actual []string) {
				assert.Equal(t, 2, len(actual))
				assert.Equal(t, "to", actual[0])
				assert.Equal(t, "top", actual[1])
			},
		},
		{
			name: "finds super word",
			word: "top",
			assert: func(actual []string) {
				assert.Equal(t, 1, len(actual))
				assert.Equal(t, "top", actual[0])
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var root *node
			if !testCase.nilRoot {
				root = &node{
					Letter: "t",
					Words:  []string{"to", "top", "tap"},
					Next: &node{
						Letter: "o",
						Words:  []string{"to", "top"},
						Less: &node{
							Letter: "a",
							Words:  []string{"tap"},
							Next: &node{
								Letter: "p",
								Words:  []string{"tap"},
							},
						},
						Next: &node{
							Letter: "p",
							Words:  []string{"top"},
						},
					},
				}
			}

			// act
			foundWords := findWord(testCase.word, root)

			// assert
			testCase.assert(foundWords)

		})
	}
}
