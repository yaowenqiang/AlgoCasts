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

func (q *Queue)Remove() int {
    if (len(q.items) > 0) {
        last := q.items[len(q.items)-1]
        q.items = q.items[0:len(q.items)-1]
        return last
    } else {
        //TODO return nil
        return 0
    }
}

func (q *Queue)Peek() interface{} {
    if (len(q.items) > 0) {
        last := q.items[len(q.items)-1]
        return last
    } else {
        return nil
    }
}

func (q *Queue)Len() int {
    return len(q.items)
}

func NewQueue() *Queue{
    return &Queue{
    }
}

func weave(a, b *Queue) *Queue{
    c := NewQueue()
    for {
        /*
        if a.Len() == 0  &&  b.Len() == 0 {
            break;
        }
        */
        if (a.Peek() == nil)  &&  (b.Peek() == nil) {
            break;
        }
        if (a.Peek() != nil ) {
            c.Add(a.Remove())
        }
        if (b.Peek() != nil ) {
            c.Add(b.Remove())
        }
    }
    return c
}

func main() {
    q := NewQueue()
    //q.Add(1)
    q.Add(2)
    q.Add(3)
    q.Add(4)
    q.Add(5)
    q.Add(6)
    q.Add(11)
    //q.Add(7)
    //q.Add(8)
    //q.Add(9)
    fmt.Println(q)
    //last := q.Remove()
    //fmt.Println(q)
    //fmt.Println(last)

    p := NewQueue()
    p.Add(6)
    p.Add(7)
    p.Add(8)
    p.Add(9)
    p.Add(10)
    fmt.Println(p)
    s := weave(q, p)
    fmt.Println(s)

}


