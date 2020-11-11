package match

import (
	"fmt"
	"strings"
)

const indentBase = " "

func (m *treematcher) PrintOrdered() {
	fmt.Println("Ordered:")
	printOrdered(m.root, 0)
}

var trees []layer

func (m *treematcher) PrintTree() {
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
