package main

import (
    "fmt"
)

type Node struct {
    data int
    left *Node
    right *Node
    isNil bool
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
func makeNilNode() *Node {
    return &Node{
        isNil: true,
    }
}

func (n *Node)Contains(data int) bool {
    if n.data < data {
        if n.right != nil  {
            return n.right.Contains(data)
        } else {
            return false
        }
    } else if n.data > data {
        if n.left != nil {
        return n.left.Contains(data)
        } else {
            return false
        }
    } else {
        return true
    }
}

func (n *Node)Validate(min *Node, max *Node) bool {
    if max.isNil == false && n.data > max.data {
        return false
    }

    if min.isNil == false && n.data < min.data {
        return false
    }

    if n.left != nil  && n.left.Validate(min, n)  ==  false {
        return false
    }

    if n.right != nil  && n.right.Validate(n, max)  ==  false {
        return false
    }

    return true


}


func (t *BST)Contains(data int) bool {
    root := t.root
    if root == nil  {
        return false
    } else {
        return root.Contains(data)
    }
}

func (t *BST)Validate() bool {
    root := t.root
    if root == nil {
        return true
    }
    min := makeNilNode()
    max := makeNilNode()
    return root.Validate(min, max)
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
    fmt.Println(tree.Contains(1))
    fmt.Println(tree.Contains(10))
    fmt.Println(tree.Contains(11))


    newTree := makeBST()
    newNode := makeNode(0)
    newNode.left = makeNode(11)
    newTree.root = newNode
    println(newTree.Validate())
}
