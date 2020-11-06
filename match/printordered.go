package match

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
	var full string
	if node.FullWord {
		full = "*"
	}
	text := strings.Join([]string{node.Letter, node.Word, full}, " ")
	println(text, depth)
}
