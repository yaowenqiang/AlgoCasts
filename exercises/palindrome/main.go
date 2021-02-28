package main

import (
    "fmt"
)

func main() {
    s := "abba"
    fmt.Println( palindrome(s))
}

func palindrome(s string) bool {
    l := len(s)
    for i:= 0; i <= l/2; i++ {
        if s[i] != s[l-i-1] {
            return false
        }
    }
    return true
}
