package main

import (
    "fmt"
)

type Node struct {
    data int
    children []*Node
}

func(n *Node)Add(data int) {
    node := MakeNode(data)
    n.children = append(n.children, node)
}

type Tree struct {

}


func MakeNode(data int) *Node {
    return &Node{
        data: data,
    }
}

func main() {
    n := MakeNode(1)
    n.Add(2)
    fmt.Printf("%+v\n", n.children)

}

