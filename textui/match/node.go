package match

import "sort"

// newnode creates a new node structure
func newnode(letter string, word string) *node {
	words := make(map[string]interface{}, 0)
	words[word] = nil
	return &node{
		Letter: letter,
		Words:  words,
	}
}

// node contains a single letter in our search tree
type node struct {
	// Less is a pointer to letters at this level that are closer to 'a'
	Less *node

	// More is a pointer to letters at this level that are closer to 'z'
	More *node

	// Next is a pointer to the root of then next level
	Next *node

	// Letter is the comparison letter held here
	Letter string

	// Words is the array of original words that this letter uniquely leads to
	Words map[string]interface{}
}

// GetWords returns an alphabetical list of words in this node
func (n *node) GetWords() []string {
	if len(n.Words) == 0 {
		return nil
	}
	var result sort.StringSlice
	for key := range n.Words {
		result = append(result, key)
	}
	result.Sort()
	return result
}

// addLetter finds the spot for and adds a letter (and returns that new node)
// if the letter already existed, then it just returns that node
func addLetter(letter string, word string, root *node, trailer **node) *node {
	if len(letter) > 1 {
		panic("addLetter passed multiple letters")
	}
	if root == nil {
		if trailer == nil {
			panic("addLetter passed a nil trailer")
		}
		*trailer = newnode(letter, word)
		return *trailer
	}

	switch {
	case letter == root.Letter:
		// Todo what to do about duplicates with instance data?
		_, exists := root.Words[word]
		if exists {
			return root
		}
		root.Words[word] = nil
		return root
	case letter < root.Letter:
		return addLetter(letter, word, root.Less, &root.Less)
	case letter > root.Letter:
		return addLetter(letter, word, root.More, &root.More)
	}
	panic("addLetter letter comparison failed")
}

// addWord distributes the letters of the word into the search tree
func addWord(originalWord string, comparisonWord string, index int, root *node, trailer **node) error {
	if index == len(comparisonWord) {
		return nil
	}
	letter := comparisonWord[index : index+1]
	this := addLetter(letter, originalWord, root, trailer)
	addWord(originalWord, comparisonWord, index+1, this.Next, &this.Next)
	return nil
}

// findLetter searches this level's search tree for the letter
func findLetter(letter string, root *node) *node {
	if len(letter) > 1 {
		panic("findLetter passed a multiple letters")
	}
	if root == nil {
		return nil
	}
	switch {
	case letter == root.Letter:
		return root
	case letter < root.Letter:
		return findLetter(letter, root.Less)
	case letter > root.Letter:
		return findLetter(letter, root.More)
	}
	panic("findLetter letter comparison failed")
}

// findWord walks the tree and finds out what stored words the given word uniquely leads to
func findWord(word string, root *node) []string {
	node := &node{}
	for i := 0; i < len(word); i++ {
		if root == nil {
			return nil
		}
		letter := word[i : i+1]
		node = findLetter(letter, root)
		if node == nil {
			// did not find the letter at this level
			return nil
		}
		root = node.Next
	}

	return node.GetWords()
}
