package radix

import (
	"fmt"
	"strings"
)

type Node struct {
	Char   rune
	IsLeaf bool
	Word   string
	Tree   map[rune]*Node
}

func NewTree() *Node {
	return &Node{
		Char:   0,
		IsLeaf: false,
		Word:   "",
		Tree:   map[rune]*Node{},
	}
}

func (node *Node) Add(word string, ) {
	node.AddNode([]rune(word), word)
}

func (node *Node) AddNode(runes []rune, w string) {
	r := runes[0]
	if _, ok := node.Tree[r]; !ok {
		node.Tree[r] = NewTree()
		node.Tree[r].Char = r
		node.Tree[r].Word = string(r)
	}
	if len(runes) == 1 {
		node.Tree[r].Word = w
		node.Tree[r].Char = r
		node.Tree[r].IsLeaf = true
	} else {
		node.Tree[r].AddNode(runes[1:], w)
	}
}

func (node *Node) PrintTree(depth int) {
	fmt.Printf("%s|  +-%s\n", strings.Repeat(" ", depth), node.Word)

	if len(node.Tree) > 0 {
		for _, sn := range node.Tree {
			sn.PrintTree(depth + len(node.Word))
		}
	}
}
