package main

import (
    "fmt"
)

type Stack struct {
    items []int
}

func (q *Stack)Push(i int) {
    q.items = append(q.items, i)
}

func (q *Stack)Pop() int {
    if (len(q.items) > 0) {
        last := q.items[len(q.items)-1]
        q.items = q.items[0:len(q.items)-1]
        return last
    } else {
        //TODO return nil
        return 0
    }
}

func (q *Stack)Peek() interface{} {
    if (len(q.items) > 0) {
        last := q.items[len(q.items)-1]
        return last
    } else {
        return nil
    }
}

func (q *Stack)Len() int {
    return len(q.items)
}

func NewStack() *Stack{
    return &Stack{
    }
}


func main() {
    q := NewStack()
    q.Push(1)
    q.Push(2)
    q.Push(3)
    q.Push(4)
    q.Push(5)
    q.Push(6)
    fmt.Println(q)
    result := q.Pop()
    fmt.Println(result)
    fmt.Println(q)
}


