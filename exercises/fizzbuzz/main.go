package main

import (
    "fmt"
)

func main() {
    fuzzbuzz(100)
}

func fuzzbuzz (in int) {
    for i:= 1; i <= in; i++ {
        if i % 3 == 0 {
            if i % 5 == 0 {
                fmt.Println("fuzzbuzz")
            } else {
                fmt.Println("fuzz")
            }
        } else {
            if i % 5 == 0 {
                fmt.Println("buzz")
            } else {
                fmt.Println(i)
            }
        }
    }
}
