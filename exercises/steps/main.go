package main

import (
    "fmt"
)

func main() {
    //step(10)
    fmt.Println(stepv2(10))
}


func step(s int) {
    for i := 0; i < s; i++ {
      rp("#", i)
      rp(" ", s -i)
      rp("\n", 1)
    }
}


// step v2, only just make the string, one print call

func stepv2(s int) string {
    var matrix []byte
    for i := 0; i < s; i++ {
        for j := 0; j < s; j++ {
            if j <=i {
                matrix = append(matrix,'#')
            } else {
               matrix = append(matrix, ' ')
            }
        }
        matrix = append(matrix,'\n')
    }
    return string(matrix)
}

func rp(s string, count int) {
    for i := 0; i < count; i++  {
        fmt.Printf("%s", s)
    }
}
