package commandprocessor

import (
	"corsairtext/universe"
	"strings"
)

// printMap handles a map command
func (cp *commandProcessor) printMap(target *string) error {
	root := cp.i.Map(target)
	cp.mapDownWalker(root, 0)
	cp.s.Out.Println("")
	cp.mapUpWalker(root.Parent, root, 1)
	return nil
}

const indentBase = " "

// mapDownWalker is the recursive downward traverser for printing maps
func (cp *commandProcessor) mapDownWalker(node *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	cp.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		cp.mapDownWalker(child, depth+1)
	}
}

// mapUpWalker is the recursive upward traverser for printing maps
func (cp *commandProcessor) mapUpWalker(node *universe.MapNode, origin *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	cp.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		if child == origin {
			continue
		}
		cp.mapDownWalker(child, depth+1)
	}
	cp.mapUpWalker(node.Parent, node, depth+1)
}
