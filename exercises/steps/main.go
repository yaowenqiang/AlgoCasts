package main

import (
    "fmt"
)

func main() {
    step(10)
}


func step(s int) {
    for i := 0; i < s; i++ {
      rp("#", i)
      rp(" ", s -i)
      rp("\n", 1)
    }
}

func rp(s string, count int) {
    for i := 0; i < count; i++  {
        fmt.Printf("%s", s)
    }
}
