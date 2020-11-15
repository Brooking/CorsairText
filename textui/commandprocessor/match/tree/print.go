package tree

import (
	"fmt"
	"strings"
)

const indentBase = " "

// PrintOrdered displays the tree in ordinal order
func (m *Matcher) PrintOrdered() {
	fmt.Println("Ordered:")
	printOrdered(m.root, 0)
}

var trees []layer

// PrintTree displays the tree in storage order
func (m *Matcher) PrintTree() {
	fmt.Println("root:")
	trees = append(trees, layer{root: m.root})
	for {
		if len(trees) == 0 {
			return
		}
		var tree layer
		tree, trees = trees[0], trees[1:]
		printTree(tree.root, tree.name, 0)
	}
}

func println(text string, depth int) {
	indent := strings.Repeat(indentBase, depth)
	fmt.Println(indent, text)
}
