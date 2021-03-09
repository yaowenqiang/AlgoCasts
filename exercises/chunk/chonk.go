package main

import (
    "fmt"
)

func main() {
    list := []int{1,2,3,4, 5, 6, 7, 8}
    fmt.Println(chunk(list,30))
}

func chunk(list []int, size int) [][]int{
    returnList := [][]int{}
    j := len(list)/ size
    if (len(list) % size > 0) {
        j = j+1
    }
    for i := 1; i <=j;i++ {
        if i == j {
            returnList = append(returnList, list[(i-1)*size:])
        } else {
            returnList = append(returnList, list[(i-1)*size:i*size])
        }
    }
    return returnList
}
