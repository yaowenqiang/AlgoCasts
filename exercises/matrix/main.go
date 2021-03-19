package main

import (
    "fmt"
)

func main() {
    m := matrix(10)

    for _, v := range m {
        for _, kv := range v {
            fmt.Printf("%3d ", kv)
        }
        fmt.Print("\n")
    }
}

func matrix(n int) [][]int {
    var m = make([][]int, n)
    for i := 0; i < n; i++ {
        m[i] = make([]int, n)
    }
    start_column := 0;
    start_row := 0;
    end_row := n-1;
    end_column := n-1;
    counter := 1
    for {
        // top row
        if counter >= n * n {
            break;
        }
        for i := start_column; i <= end_column; i++ {
            m[start_row][i] = counter
            counter++
        }
        start_row = start_row +1

        //right column
        for i := start_row; i <= end_row; i++ {
            m[i][end_column] = counter
            counter++
        }

        end_column = end_column-1

        //bottom row

        for i := end_column; i >= start_column; i-- {
            m[end_row][i] = counter
            counter++
        }

        end_row = end_row-1



        // start column

        for i := end_row; i >= start_row; i-- {
            m[i][start_column] = counter
            counter++
        }
        start_column = start_column+1

        /*
        if ((start_column <= end_column) && (start_row <= end_row) ) {
            break;
        }
        */
    }
    return m
}
