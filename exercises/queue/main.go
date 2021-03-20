package main

import (
    "fmt"
)

type Queue struct {
    items []int
}

func (q *Queue)Add(i int) {
    if len(q.items) == 0 {
        q.items = append(q.items, i)
    } else {
        q.items = append([]int{i}, q.items[0:]...)
    }
}

func (q *Queue)Remove() {
    if (len(q.items) > 0) {
        q.items = q.items[0:len(q.items)-1]
    }
}

func NewQueue() *Queue{
    return &Queue{
    }
}

func main() {
    q := NewQueue()
    q.Add(1)
    q.Add(2)
    q.Add(3)
    q.Add(4)
    q.Add(5)
    q.Add(6)
    q.Add(7)
    q.Add(8)
    q.Add(9)
    fmt.Println(q)
    q.Remove()
    fmt.Println(q)
}


