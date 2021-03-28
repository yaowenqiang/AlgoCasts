package main

import (
    "fmt"
)

type Node struct {
    data int
    left *Node
    right *Node
}

type BST struct {
    root *Node
}


func (n *Node)Insert(data int) {
    if data < n.data  {
        if n.left != nil {
            n.left.Insert(data)
        } else {
            n.left = makeNode(data)
        }
    } else if (data > n.data) {
        if n.right != nil {
            n.right.Insert(data)
        } else {
            n.right = makeNode(data)
        }
    }
}

func makeNode(data int) *Node {
    return &Node{
        data: data,
    }
}


func makeBST() *BST {
    return &BST{
    }
}


func main() {
    tree := makeBST()
    fmt.Println(tree)
    n := makeNode(0)
    n.Insert(1)
    n.Insert(10)
    n.Insert(5)
    n.Insert(2)
    n.Insert(-11)
    n.Insert(100)
    tree.root = n
    fmt.Println(tree)
}
