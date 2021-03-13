package main

import (
    "strings"
    "unicode"
    "fmt"
)

func main() {
    //fmt.Println(capitalize("hello world 你好    你也好"))
    capitalize2("hello world　你好")
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

func capitalize2(s string) string {
    for i, k :=  range s {
        if i == 0 || (s[i-1] == ' ' ){
            s = s[:i] + string((unicode.ToUpper(k))) + s[i+1:]
        }
    }
    return s
}
