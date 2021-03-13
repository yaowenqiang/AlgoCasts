package main

import (
    "strings"
    "fmt"
)

func main() {
    fmt.Println(capitalize("hello world 你好    你也好"))
}

func capitalize(s string) string {
    words := strings.Split(s, " ")
    for i, k :=  range words {
        if  (len(k) > 0) && (k[0] >= byte('A') ) && ( k[0] <= byte('z')) {
            words[i] = strings.ToUpper(string(k[0])) + k[1:]
        }
    }

    return strings.Join(words, " ")
}
