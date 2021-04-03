package main

import (
    "fmt"
)

func main() {
    //arr1 := []int{1,2,10}
    //arr2 := []int{-3,0,7}
    arr3 := []int{4,5,6,5,3,3,2,1,23,4,5,6,10,8}
    //arr3 := []int{}
    //arr3 := []int{4,5}
    //fmt.Println(merge(arr1,arr2))
    fmt.Println(mergeSort(arr3))
}

func mergeSort(arr []int) []int {
    if len(arr) <= 1  {
        return arr
    }

    center := len(arr) / 2

    left := arr[:center]
    right := arr[center:]
    return merge(mergeSort(left), mergeSort(right))

}

func merge(left []int, right []int) []int {
    result  := []int{}
    for {
        if len(left) == 0 || len(right) == 0 {
            break
        }

        if left[0] < right[0] {
            result = append(result,left[0])
            left = left[1:]
        } else {
            result = append(result,right[0])
            right = right[1:]
        }
    }

    return append(append(result, left...), right...)
}
