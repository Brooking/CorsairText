package tree

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
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "m", root.Letter)
				assert.Equal(t, 1, len(root.Words))
				assert.Equal(t, "ma", root.GetWords()[0])
			},
		},
		{
			name:   "add single lesser letter to single",
			letter: "m",
			word:   "ma",
			root: &node{
				Letter: "n",
				Words: map[string]interface{}{
					"no": nil,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.NotNil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "n", root.Letter)
				assert.Equal(t, 1, len(root.Words))
				assert.Equal(t, "no", root.GetWords()[0])

				assert.Nil(t, root.Less.Less)
				assert.Nil(t, root.Less.More)
				assert.Nil(t, root.Less.Next)
				assert.Equal(t, "m", root.Less.Letter)
				assert.Equal(t, 1, len(root.Less.Words))
				assert.Equal(t, "ma", root.Less.GetWords()[0])
			},
		},
		{
			name:   "add single greater letter to single",
			letter: "p",
			word:   "pa",
			root: &node{
				Letter: "n",
				Words: map[string]interface{}{
					"no": nil,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.NotNil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "n", root.Letter)
				assert.Equal(t, 1, len(root.Words))
				assert.Equal(t, "no", root.GetWords()[0])

				assert.Nil(t, root.More.Less)
				assert.Nil(t, root.More.More)
				assert.Nil(t, root.More.Next)
				assert.Equal(t, "p", root.More.Letter)
				assert.Equal(t, 1, len(root.More.Words))
				assert.Equal(t, "pa", root.More.GetWords()[0])
			},
		},
		{
			name:   "add single letter that exists",
			letter: "m",
			word:   "my",
			root: &node{
				Letter: "m",
				Words: map[string]interface{}{
					"ma": nil,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "m", root.Letter)
				assert.Equal(t, 2, len(root.Words))
				assert.Equal(t, "ma", root.GetWords()[0])
				assert.Equal(t, "my", root.GetWords()[1])
			},
		},
		{
			name:   "add single letter word",
			letter: "i",
			word:   "i",
			final:  true,
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "i", root.Letter)
				assert.Equal(t, 1, len(root.Words))
				assert.Equal(t, "i", root.GetWords()[0])
			},
		},
		{
			name:   "add word to single letter word",
			letter: "i",
			word:   "it",
			final:  false,
			root: &node{
				Letter: "i",
				Words: map[string]interface{}{
					"i": nil,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "i", root.Letter)
				assert.Equal(t, 2, len(root.Words))
				assert.Equal(t, "i", root.GetWords()[0])
				assert.Equal(t, "it", root.GetWords()[1])
			},
		},
		{
			name:   "add single letter word to existing word",
			letter: "i",
			word:   "i",
			final:  true,
			root: &node{
				Letter: "i",
				Words: map[string]interface{}{
					"it": nil,
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.Nil(t, root.Next)
				assert.Equal(t, "i", root.Letter)
				assert.Equal(t, 2, len(root.Words))
				assert.Equal(t, "i", root.GetWords()[0])
				assert.Equal(t, "it", root.GetWords()[1])
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
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.NotNil(t, root.Next)
				assert.Equal(t, "y", root.Letter)
				assert.Equal(t, 1, len(root.Words))
				assert.Equal(t, "yes", root.GetWords()[0])

				assert.Nil(t, root.Next.Less)
				assert.Nil(t, root.Next.More)
				assert.NotNil(t, root.Next.Next)
				assert.Equal(t, "e", root.Next.Letter)
				assert.Equal(t, 1, len(root.Next.Words))
				assert.Equal(t, "yes", root.Next.GetWords()[0])

				assert.Nil(t, root.Next.Next.Less)
				assert.Nil(t, root.Next.Next.More)
				assert.Nil(t, root.Next.Next.Next)
				assert.Equal(t, "s", root.Next.Next.Letter)
				assert.Equal(t, 1, len(root.Next.Next.Words))
				assert.Equal(t, "yes", root.Next.Next.GetWords()[0])
			},
		},
		{
			name: "add long word over short",
			word: "yes",
			root: &node{
				Letter: "y",
				Words: map[string]interface{}{
					"ya": nil,
				},
				Next: &node{
					Letter: "a",
					Words: map[string]interface{}{
						"ya": nil,
					},
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.NotNil(t, root.Next)
				assert.Equal(t, "y", root.Letter)
				assert.Equal(t, 2, len(root.Words))
				assert.Equal(t, "ya", root.GetWords()[0])
				assert.Equal(t, "yes", root.GetWords()[1])

				assert.Nil(t, root.Next.Less)
				assert.NotNil(t, root.Next.More)
				assert.Nil(t, root.Next.Next)
				assert.Equal(t, "a", root.Next.Letter)
				assert.Equal(t, 1, len(root.Next.Words))
				assert.Equal(t, "ya", root.Next.GetWords()[0])

				assert.Nil(t, root.Next.More.Less)
				assert.Nil(t, root.Next.More.More)
				assert.NotNil(t, root.Next.More.Next)
				assert.Equal(t, "e", root.Next.More.Letter)
				assert.Equal(t, 1, len(root.Next.More.Words))
				assert.Equal(t, "yes", root.Next.More.GetWords()[0])

				assert.Nil(t, root.Next.More.Next.Less)
				assert.Nil(t, root.Next.More.Next.More)
				assert.Nil(t, root.Next.More.Next.Next)
				assert.Equal(t, "s", root.Next.More.Next.Letter)
				assert.Equal(t, 1, len(root.Next.More.Next.Words))
				assert.Equal(t, "yes", root.Next.More.Next.GetWords()[0])
			},
		},
		{
			name: "add short word over long",
			word: "ya",
			root: &node{
				Letter: "y",
				Words: map[string]interface{}{
					"yes": nil,
				},
				Next: &node{
					Letter: "e",
					Words: map[string]interface{}{
						"yes": nil,
					},
					Next: &node{
						Letter: "s",
						Words: map[string]interface{}{
							"yes": nil,
						},
					},
				},
			},
			assert: func(root *node) {
				assert.NotNil(t, root)
				assert.Nil(t, root.Less)
				assert.Nil(t, root.More)
				assert.NotNil(t, root.Next)
				assert.Equal(t, "y", root.Letter)
				assert.Equal(t, 2, len(root.Words))
				assert.Equal(t, "ya", root.GetWords()[0])
				assert.Equal(t, "yes", root.GetWords()[1])

				assert.NotNil(t, root.Next.Less)
				assert.Nil(t, root.Next.More)
				assert.NotNil(t, root.Next.Next)
				assert.Equal(t, "e", root.Next.Letter)
				assert.Equal(t, 1, len(root.Next.Words))
				assert.Equal(t, "yes", root.Next.GetWords()[0])

				assert.Nil(t, root.Next.Next.Less)
				assert.Nil(t, root.Next.Next.More)
				assert.Nil(t, root.Next.Next.Next)
				assert.Equal(t, "s", root.Next.Next.Letter)
				assert.Equal(t, 1, len(root.Next.Next.Words))
				assert.Equal(t, "yes", root.Next.Next.GetWords()[0])

				assert.Nil(t, root.Next.Less.Less)
				assert.Nil(t, root.Next.Less.More)
				assert.Nil(t, root.Next.Less.Next)
				assert.Equal(t, "a", root.Next.Less.Letter)
				assert.Equal(t, 1, len(root.Next.Less.Words))
				assert.Equal(t, "ya", root.Next.Less.GetWords()[0])
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
				assert.Nil(t, actual)
			},
		},
		{
			name: "fails when short word is not in tree",
			word: "i",
			assert: func(actual []string) {
				assert.Nil(t, actual)
			},
		},
		{
			name: "fails when long word is not in tree",
			word: "spider",
			assert: func(actual []string) {
				assert.Nil(t, actual)
			},
		},
		{
			name: "fails on found but ambiguous word",
			word: "t",
			assert: func(actual []string) {
				assert.Equal(t, 3, len(actual))
				assert.Equal(t, "tap", actual[0])
				assert.Equal(t, "to", actual[1])
				assert.Equal(t, "top", actual[2])
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
					Words: map[string]interface{}{
						"to":  nil,
						"top": nil,
						"tap": nil,
					},
					Next: &node{
						Letter: "o",
						Words: map[string]interface{}{
							"to":  nil,
							"top": nil,
						},
						Less: &node{
							Letter: "a",
							Words: map[string]interface{}{
								"tap": nil,
							},
							Next: &node{
								Letter: "p",
								Words: map[string]interface{}{
									"tap": nil,
								},
							},
						},
						Next: &node{
							Letter: "p",
							Words: map[string]interface{}{
								"top": nil,
							},
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
