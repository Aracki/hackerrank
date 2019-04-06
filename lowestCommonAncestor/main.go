package main

import "fmt"

func main() {

	//tree := []int{4, 2, 3, 1, 7, 6}

	// create a tree
	n1 := &Node{data: 1}
	n3 := &Node{data: 3}
	n2 := &Node{data: 2}
	n2.left = n1
	n2.right = n3

	n6 := &Node{data: 6}
	n7 := &Node{data: 7}
	n7.left = n6

	root := &Node{data: 4}
	root.left = n2
	root.right = n7

	//lastCommonAncestor(root, n1, n7)
	fmt.Println("lca:", lowestCommonAncestor(root, n1, n2).data)
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func lowestCommonAncestor(root, v1, v2 *Node) (lca *Node) {

	fmt.Println(root.data)

	if (root.data-v1.data) * (root.data-v2.data) <= 0 {
		return root
	} else if root.data > v1.data {
		return lowestCommonAncestor(root.left, v1, v2)
	} else {
		return lowestCommonAncestor(root.right, v1, v2)
	}
}
