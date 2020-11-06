package match

import (
	"strings"
)

func printTree(root *node, name string, depth int) {
	println("", 0)
	println(strings.Join([]string{"(", name, ")"}, ""), 0)
	printTreeWorker(root, name, depth)
}

func printTreeWorker(root *node, name string, depth int) {
	if root == nil {
		return
	}
	printTreeWorker(root.Less, name, depth+1)
	printTreeNode(root, name, depth)
	printTreeWorker(root.More, name, depth+1)
}

func printTreeNode(node *node, name string, depth int) {
	var end string
	if node.Next == nil {
		end = node.Word
	}
	text := strings.Join([]string{node.Letter, end}, " ")
	println(text, depth)

	if node.Next != nil {
		trees = append(trees, layer{
			name: name + node.Letter,
			root: node.Next,
		})
	}
}

type layer struct {
	name string
	root *node
}
