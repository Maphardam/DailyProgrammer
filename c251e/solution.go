package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    table := make([][]bool, 0)

    scanner := bufio.NewScanner(file)
    idx := 0
    for scanner.Scan() {
        table = append(table, make([]bool, 0))

        for _, b := range []byte(scanner.Text()) {
            if b == '*' {
                table[idx] = append(table[idx], true)
            } else {
                table[idx] = append(table[idx], false)
            }
        }
        idx++
    }

    if len(table) == 0 {
        log.Fatal("Empty table.")
    }

    rows := make([][]int, 0, len(table))
    for _, row := range table {
        rows = append(rows, count(row))
    }

    cols := make([][]int, 0, len(table))
    for i := range table[0] {
        cols = append(cols, count(column(table, i)))
    }

    //print user friendly
    m_row := max(rows)
    m_col := max(cols)
    fmt.Println(m_row, m_col)

    //print cols
    for i := m_col-1; i >= 0; i-- {
        for j := 0; j < m_row; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < len(cols); j++ {
            if len(cols[j]) <= i {
                fmt.Print(" ")
            } else {
                fmt.Printf("%d", cols[j][i])
            }
        }
        fmt.Print("\n")
    }

    //print rows
    for _, r := range rows {
        for i := 0; i < m_row - len(r); i++ {
            fmt.Print(" ")
        }
        for i := 0; i < len(r); i++ {
            fmt.Printf("%d", r[i])
        }
        fmt.Print("\n")
    }
}

func count(line []bool) []int {
    result := make([]int, 0)
    is := false
    x := 0
    for i := 0; i < len(line); i++ {
        if line[i] {
            if !is {
                is = true
                x = 1
            } else {
                x++
            }
        } else {
            if is {
                is = false
                result = append(result, x)
            }
        }
    }
    if is {
        result = append(result, x)
    }
    return result
}

func column(m [][]bool, i int) []bool {
    c := make([]bool, 0)
    for _, r := range m {
        if i >= len(r) {
            log.Fatal("Matrix not complete.")
        }
        c = append(c, r[i])
    }
    return c
}

func max(m [][]int) int {
    max := 0
    for _, r := range m {
        if len(r) > max {
            max = len(r)
        }
    }
    return max
}
