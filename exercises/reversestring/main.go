package main

import (
    "fmt"
)

func main() {
    str := []rune("apple你好")
    str2 := string(reverse(str))
    fmt.Printf("before: %+v\n",string(str))
    fmt.Printf("after : %+v\n",str2)

}

// reverse a string
func reverse(str []rune) []rune  {
    var returnArr  []rune
    for i := len(str)-1;i>=0;i-- {
        returnArr = append(returnArr, str[i])
    }
    return returnArr
}
