package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "hello world"
    fmt.Println(vowel(s))
}

// return vowel(元音字母) character numbers of a given string
func vowel(s string) int{
    vowels := "aeiouAEIOU"
    count := 0
    for _, i := range s {
        if strings.Index(vowels, string(i)) >= 0 {
            count++
        }
    }
    return count
}
