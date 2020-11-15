package tree

import (
	"strings"
)

func printOrdered(root *node, depth int) {
	if root == nil {
		return
	}
	printOrdered(root.Less, depth)
	printNode(root, depth)
	printOrdered(root.More, depth)
	printOrdered(root.Next, depth+1)
}

func printNode(node *node, depth int) {
	text := strings.Join([]string{node.Letter}, " ")
	println(text, depth)
}
