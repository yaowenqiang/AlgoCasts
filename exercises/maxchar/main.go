package main

import (
    "fmt"
)

func main() {
    str := "hello world,你好，你好，你好，你好，你好"
    fmt.Println(str)
    hit, num := maxChar(str)
    fmt.Println(string(hit), num)
}


func maxChar( s string ) (rune, int64) {
    maps := make(map[rune]int64)
    var max int64
    var hit rune
    for _,i := range s {
        maps[i]+=1;
        if max < maps[i] {
            max = maps[i]
            hit = i
        }
    }
    return hit, max
}
