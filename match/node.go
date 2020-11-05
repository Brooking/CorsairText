package match

// newnode creates a new node structure
func newnode(letter string, word string, full bool) *node {
	return &node{
		Letter:   letter,
		Word:     word,
		FullWord: full,
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

	// Letter is the letter held here
	Letter string

	// Word is the word that this letter uniquely leads to
	Word string

	// FullWord indicates that this is the end of a known word
	FullWord bool
}

// addLetter finds the spot for and adds a letter (and returns that new node)
// if the letter already existed, then it just returns that node
func addLetter(letter string, word string, finalLetter bool, root *node, trailer **node) *node {
	if len(letter) > 1 {
		panic("addLetter passed multiple letters")
	}
	if root == nil {
		if trailer == nil {
			panic("addLetter passed a nil trailer")
		}
		*trailer = newnode(letter, word, finalLetter)
		return *trailer
	}

	switch {
	case letter == root.Letter:
		switch {
		case finalLetter:
			root.Word = word
			root.FullWord = true
		case !root.FullWord:
			root.Word = ""
		}
		return root
	case letter < root.Letter:
		return addLetter(letter, word, finalLetter, root.Less, &root.Less)
	case letter > root.Letter:
		return addLetter(letter, word, finalLetter, root.More, &root.More)
	}
	panic("addLetter letter comparison failed")
}

// addWord distributes the letters of the word into the search tree
func addWord(word string, index int, root *node, trailer **node) error {
	if index == len(word) {
		return nil
	}
	letter := word[index : index+1]
	finalLetter := index+1 == len(word)
	this := addLetter(letter, word, finalLetter, root, trailer)
	addWord(word, index+1, this.Next, &this.Next)
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

// findWord walks the tree and finds out what stored word the given word uniquely leads to
func findWord(word string, root *node) string {
	node := &node{}
	for i := 0; i < len(word); i++ {
		if root == nil {
			return ""
		}
		letter := word[i : i+1]
		node = findLetter(letter, root)
		if node == nil {
			// did not find the letter at this level
			return ""
		}
		root = node.Next
	}

	return node.Word
}
