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

func (ls *LinkedList)InsertAt(index int, data string) {
    if (ls.head == nil ) {
        ls.head = NewNode(data)
        return
    } else {
        if index == 0 {
            if ls.head.next == nil {
                newNode := NewNode(data)
                newNode.next = ls.head
                ls.head = newNode
            } else {
                newNode := NewNode(data)
                newNode.next = ls.head
                ls.head = newNode
            }
        } else {
            indexNode := ls.GetAt(index-1)
            if indexNode != nil && indexNode.next != nil {
                newNode := NewNode(data)
                newNode.next = indexNode.next
                indexNode.next = newNode
            } else {
                ls.InsertLast(data)
            }
        }
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

func (ls *LinkedList)Midpoint() *Node {
    if ls.head == nil {
        return nil
    }
    slow := ls.head
    fast := ls.head
    for {
        if fast.next != nil && fast.next.next != nil {
            fast = fast.next.next
            slow = slow.next
        } else {
            break
        }
    }
    return slow
}

func (ls *LinkedList)Circular() bool {
    if ls.head == nil {
        return false
    }
    slow := ls.head
    fast := ls.head
    for {
        if fast.next != nil && fast.next.next != nil {
            fast = fast.next.next
            slow = slow.next
            if slow == fast {
                return true
            }
        } else {
            return false
        }
    }
}
func main() {
    ls := NewLinkedList()
    ls.InsertAt(0,"a")
    ls.InsertAt(1,"b")
    ls.InsertAt(2,"c")
    ls.InsertAt(3,"d")

    last := ls.GetLast()

    last.next = ls.GetAt(2)

    fmt.Println(ls.Circular())

}



func NewLinkedList() *LinkedList {
    return &LinkedList{
    }
}
