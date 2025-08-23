package main

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{Value: val}
	}
	if val < n.Value {
		n.left = n.left.Insert(val)
	} else if val > n.Value {
		n.right = n.right.Insert(val)
	}
	return n
}



func (n *Node) Inorder() {
	if n != nil {
		n.left.Inorder()
		fmt.Print(n.Value, " ")
		n.right.Inorder()
	}
}

func (n *Node) Preorder() {
	if n != nil {
		fmt.Print(n.Value, " ")
		n.left.Inorder()
		n.right.Inorder()
	}
}

func (n *Node) PostOrder() {
	if n != nil {
		n.left.PostOrder()
		n.right.PostOrder()
		fmt.Print(n.Value, " ")
	}
}

func (n *Node) Delete(val int) *Node {
	if n == nil {
		return nil
	}
	if val < n.Value {
		n.left = n.left.Delete(val)
	} else if val > n.Value {
		n.right = n.right.Delete(val)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		
		minNode := n.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		n.Value = minNode.Value
		n.right = n.right.Delete(minNode.Value)
	}
	return n
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.Value==val{
		return true
	}else if val<n.Value{
		return n.left.Search(val)
	}else{
		return n.right.Search(val)
	}
}

func main() {
	var root *Node
	val := []int{50, 30, 90, 40, 70, 30}
	for _, num := range val {
		root = root.Insert(num)
	}

	fmt.Println("PreOrder :")
	root.Preorder()
	fmt.Println()

	fmt.Println("Post Order :")
	root.PostOrder()
	fmt.Println()

	fmt.Println("Search 30",root.Search(30))

	root.Delete(90)
	fmt.Println("After delete :")
	fmt.Println("Inorder :")
	root.Inorder()
	fmt.Println()


}
