package textui

import (
	"corsairtext/universe"
	"strings"
)

// printMap handles a map command
func (t *textUI) printMap(target *string) error {
	root := t.i.Map(target)
	t.mapDownWalker(root, 0)
	t.s.Out.Println("")
	t.mapUpWalker(root.Parent, root, 1)
	return nil
}

const indentBase = " "

// mapDownWalker is the recursive downward traverser for printing maps
func (t *textUI) mapDownWalker(node *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	t.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		t.mapDownWalker(child, depth+1)
	}
}

// mapUpWalker is the recursive upward traverser for printing maps
func (t *textUI) mapUpWalker(node *universe.MapNode, origin *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	t.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		if child == origin {
			continue
		}
		t.mapDownWalker(child, depth+1)
	}
	t.mapUpWalker(node.Parent, node, depth+1)
}
