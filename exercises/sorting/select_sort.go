package main

import (
    "fmt"
)

func main() {
    arr := []int{1,2,3,4,5,6,6,5,4}
    fmt.Println(selectSort(arr))
}

func selectSort (arr []int) []int {
    for i := 0; i < len(arr); i++ {
        minIndex := i
        for j := i+1; j < len(arr); j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        if minIndex != i {
            tmp := arr[i]
            arr[i], arr[minIndex] = arr[minIndex], arr[i]
            arr[minIndex] = tmp
        }
    }
    return arr
}
