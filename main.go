package main

import (
	"github.com/tufanbarisyildirim/radix-tree/radix"
)

func main() {

	tree := radix.NewTree()

	tree.Add("f")
	tree.Add("foo")
	tree.Add("foobar")
	tree.Add("barfoo")
	tree.Add("baro")
	tree.Add("fobar")
	tree.Add("tufan")
	tree.Add("baris")
	tree.Add("barid")

	tree.PrintTree(0)

	//fmt.Println(tree.Search("fo"),tree.Search("foo"),tree.Search("tufan"))

}
