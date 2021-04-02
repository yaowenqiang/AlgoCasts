package main

import (
    "fmt"
)

func bubbleSort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                tmp := arr[j]
                arr[j+1] = arr[j]
                arr[j] = tmp
            }
        }
    }
    return arr
}

func main() {
    arr := []int{1,4,5,7,5,4,4,5,6,8,9,7,5,11}

    fmt.Println(bubbleSort(arr))
}
