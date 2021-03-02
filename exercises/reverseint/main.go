package main

import(
    "fmt"
    "math"
)

func main() {
    reverseInt(12345678)
}

func reverseInt(s int64) int64 {
    fmt.Println(s)
    var arr []int64
    for {
        if s == 0 {
            break
        }
        r := s % 10
        arr = append(arr, r)
        s = s / 10
    }

    var ret int64
    var l = len(arr)
    //fmt.Println(arr)
    for k := 0; k <=l-1;k++ {
        tens := math.Pow(10.0, float64(l-k-1))
        ret += int64(float64(arr[k]) * tens)
    }
    fmt.Println(ret)

    return ret
}
