package main

import (
    "fmt"
)

func main() {
    str := []byte("apple")
    str2 := string(reverse(str))
    fmt.Printf("before: %+v\n",string(str))
    fmt.Printf("after : %+v\n",str2)

}

// reverse a string
func reverse(str []byte) []byte  {
    var returnArr  []byte
    for i := len(str)-1;i>=0;i-- {
        returnArr = append(returnArr, str[i])
    }
    return returnArr
}
