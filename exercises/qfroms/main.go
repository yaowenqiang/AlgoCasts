package main

import (
    "fmt"
)

type Stack struct {
    items []int
}

type Queue struct {
    in *Stack
    out *Stack
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

func NewQueue() *Queue{
    return &Queue{
        in: NewStack(),
        out: NewStack(),
    }
}

func (q *Queue)Remove() int {
    if (q.out.Peek() != nil ) {
        return q.out.Pop()
    } else {
        for {
            if (q.in.Len() > 0) {
                q.out.Push(q.in.Pop())
            } else {
                break
            }
        }
        return q.out.Pop()
    }
}

func (q *Queue)Len() int {
    return q.in.Len() + q.out.Len()
}

func (q *Queue)Add(i int) {
    q.in.Push(i)
}


func main() {
    q := NewQueue()
    q.Add(1)
    q.Add(2)
    q.Add(3)
    q.Add(4)
    q.Add(5)
    q.Add(6)
    fmt.Println(q.in)
    fmt.Println(q.out)
    result := q.Remove()
    result2 := q.Remove()
    fmt.Println(q.in)
    fmt.Println(q.out)
    q.Add(7)
    q.Add(8)
    fmt.Println(q.in)
    fmt.Println(q.out)
    fmt.Println(result)
    fmt.Println(result2)


}


