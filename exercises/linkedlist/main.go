package main

import (
    "fmt"
)


type Node struct {
    data string
    next *Node
}

func NewNode(data string) *Node {
    return  &Node{
        data: data,
    }
}

type LinkedList struct {
    head *Node
    tail *Node
}

func (ls *LinkedList)InsertFirst(data string) {
    if ls.head == nil {
        ls.head = NewNode(data)
    } else {
        oldHead :=  ls.head
        ls.head = NewNode(data)
        ls.head.next = oldHead
    }
}

func (ls *LinkedList)Size() int {
    size := 0
    current  := ls.head
    for {
        if current == nil {
            break;
        } else {
            size ++
            current = current.next
        }
    }
    return size
}

func (ls *LinkedList)GetFirst() *Node {
    return ls.head
}

func (ls *LinkedList)RemoveFirst()  {
    head := ls.head
    if head != nil {
        ls.head = head.next
        head = nil
    }
}

func (ls *LinkedList)RemoveLast() {
    current := ls.head
    if current == nil {
        return
    }
    if current.next == nil {
        current = nil
        return
    }
    for {
        if current.next != nil && current.next.next == nil {
           current.next  = nil
           break;
        }
        current = current.next
    }
}

func (ls *LinkedList)GetLast() *Node {
    current  := ls.head
    for {
        if current == nil  || current.next == nil {
            break;
        } else {
            current = current.next
        }
    }
    return current
}

func (ls *LinkedList)InsertLast(data string)  {
    last  := ls.GetLast()
    newNode := NewNode(data)
    if last == nil {
        ls.head = newNode
    } else {
        last.next = newNode
    }
}

func (ls *LinkedList)GetAt(index int) *Node {
    count := 0
    current := ls.head
    for {
        if current == nil || count == index || current.next == nil {
            break
        }
        current = current.next
        count++
    }
    if count < index {
        return nil
    }
    return current
}

func (ls *LinkedList)RemoveAt(index int) {
    if (ls.head == nil ) {
        return
    } else {
        if index == 0 {
            ls.head = ls.head.next
            return
        }
    }
    node := ls.GetAt(index -1)
    if node == nil {
        return
    }
    if node.next != nil && node.next.next != nil {
        node.next = node.next.next
    } else {
        node.next = nil
    }
}

func (ls *LinkedList)Clear() {
    current  := ls.head
    ls.head = nil
    for {
        if current.next == nil {
            break;
        } else {
            next := current.next
            current = nil
            current = next
        }
    }
}

func main() {
    ls := NewLinkedList()
    ls.InsertFirst("last Node")
    ls.InsertFirst("first Node")
    ls.InsertFirst("second Node")
    ls.InsertFirst("second Node")
    ls.InsertFirst("second Node")
    ls.InsertFirst("second Node")
    ls.InsertFirst("first head  Node")
    //node := NewNode("hello")
    //ls.head = node
    //fmt.Printf("%+v", ls.head)
    //fmt.Printf("%+v", node)
    fmt.Println(ls.head)
    fmt.Println(ls.head.next)
    fmt.Println(ls.Size())
    fmt.Println(ls.GetFirst())
    ls.RemoveFirst()
    fmt.Println(ls.GetFirst())

    fmt.Println(ls.GetLast())
    fmt.Println("RemoveLast")
    ls.RemoveLast()
    fmt.Println(ls.GetLast())
    ls.Clear()
    fmt.Println(ls.Size())
    fmt.Println(ls.GetFirst())
    fmt.Println(ls.GetLast())
    ls.RemoveLast()
    ls.RemoveLast()
    ls.RemoveLast()
    fmt.Println(ls.GetLast())
    ls.InsertFirst("new last")
    fmt.Println(ls.GetLast())
    ls.Clear()
    ls.InsertLast("latest new last")
    fmt.Println(ls.GetLast())
    ls.InsertLast("1")
    ls.InsertLast("2")
    fmt.Println(ls.GetFirst())
    fmt.Println(ls.GetAt(1))

    ls.Clear()
    ls.RemoveAt(0)
    ls.InsertLast("2")
    fmt.Printf("%+v\n", ls.GetAt(0))
    ls.RemoveAt(0)
    fmt.Printf("%+v\n", ls.GetAt(0))
    ls.RemoveAt(1)
    ls.InsertFirst("first Node")
    ls.InsertFirst("second Node")
    ls.InsertFirst("second Node")
    ls.RemoveAt(1)
    ls.InsertFirst("first Node")
    ls.Clear()
    ls.InsertFirst("1")
    ls.InsertFirst("2")
    ls.InsertFirst("3")
    ls.RemoveAt(2)
    fmt.Println(ls.head)
    fmt.Println(ls.head.next)
    fmt.Println(ls.head.next.next)

}



func NewLinkedList() *LinkedList {
    return &LinkedList{
    }
}
