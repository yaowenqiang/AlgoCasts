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

/*
func(n *Node)Remove(data int) {
    for i :=0; i <  len(n.children); i ++ {
        if n.children[i].data == data {
            n.children = append(n.children[:i], n.children[i+1:]...)
        }
    }
}

*/

func(n *Node)Remove(data int) {
    for  {
        hit := 0
        length := len(n.children);
        for i := 0; i < length; i ++ {
            if n.children[i].data == data {
                n.children = append(n.children[:i], n.children[i+1:]...)
                hit = 1
                break
            }
        }
        if hit == 0 {
            break
        }
    }
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
    n.Add(3)
    n.Add(3)
    n.Add(3)
    n.Add(3)
    n.Add(3)
    n.Add(3)
    n.Add(4)
    fmt.Printf("%+v\n", n.children)
    n.Remove(3)
    n.Add(5)
    n.Add(6)
    n.Add(7)
    n.Add(8)
    fmt.Printf("%+v\n", n.children)

}

