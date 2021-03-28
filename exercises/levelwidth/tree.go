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
    root *Node
}

func (t *Tree) traverseBF(fn func(*Node)) {
    arr := []*Node{t.root}
    for {
        if len(arr) == 0 {
            break
        }

        node := arr[0]
        arr = arr[1:]
        arr = append(arr, node.children...)
        fn(node)
    }
}

func (t *Tree) traverseDF(fn func(*Node)) {
    arr := []*Node{t.root}
    for {
        if len(arr) == 0 {
            break
        }

        node := arr[0]
        arr = arr[1:]
        for _, i := range node.children {
            arr = append([]*Node{i}, arr...)
        }
        fn(node)
    }
}

func(t *Tree)LevelWidth() []int {
    root := t.root
    if root == nil {
        return []int{}
    }
    counter := []int{0}
    arr := []*Node{root, MakeNode(-1)}
    for {
        if len(arr) <=1 {
            break
        }

        node := arr[0]
        arr = arr[1:]
        if node.data == -1 {
            counter = append(counter, 0)
            arr = append(arr, node)
        } else {
            arr = append(arr ,node.children...)
            counter[len(counter)-1]++
        }

    }
    return counter

}

func updateNode(n *Node) {
    n.data*= 2
}

func MakeNode(data int) *Node {
    return &Node{
        data: data,
    }
}

func MakeTree() *Tree {
    return &Tree{
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
    n.children[1].Add(2)
    n.children[1].Add(2)
    n.children[1].Add(2)
    fmt.Printf("%+v\n", n.children)
    n.Remove(3)
    n.Add(5)
    n.Add(6)
    n.Add(7)
    n.Add(8)
    n.children[1].Add(2)
    n.children[1].Add(2)
    n.children[1].Add(2)
    fmt.Printf("%+v\n", n.children)


    tree := MakeTree()
    tree.root = n

    fmt.Println(tree)

    tree.traverseBF(updateNode)
    fmt.Println(tree)

    tree.traverseDF(updateNode)
    fmt.Println(tree)

    fmt.Println(tree.LevelWidth())

    b := MakeTree()
    fmt.Println(b.LevelWidth())

}

