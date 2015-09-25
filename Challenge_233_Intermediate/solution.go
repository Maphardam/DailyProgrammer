package main

import (
    "os"
    "log"
    "bufio"
    "bytes"
    "fmt"
    "math"
    "math/rand"
    "time"
)

type Update struct {
    X,Y int
    Char  byte
}

func main() {
    canvas := read(os.Args[1])
    
    for _,row := range canvas {
        fmt.Println(string(row))
    }
    fmt.Println()
    ch := make(chan Update)
    for y, _ := range canvas {
        for x, _ := range canvas[y] {
            go update(canvas, x, y, ch)
        }
    }  
    
    nCanvas := make([][]byte, len(canvas))
    for i,_ := range canvas {
        nCanvas[i] = make([]byte, len(canvas[i]))
    }
    
    upd := Update{}
    for i := 0; i < len(canvas)*len(canvas[0]); i++{
        upd = <- ch
        nCanvas[upd.Y][upd.X] = upd.Char     
    }
    
    for _,row := range nCanvas {
        fmt.Println(string(row))
    }
    
    
}

func read(fn string) (canvas [][]byte) {
    canvas = make([][]byte, 0)
    file, err := os.Open(fn)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var maxLen int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        canvas = append(canvas, scanner.Bytes())
        
        maxLen = int(math.Max(float64(maxLen), 
                              float64(len(canvas[len(canvas)-1]))))
    }
    fmt.Println(cap(canvas))
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    for _,row := range canvas {
        fmt.Printf("%v %v\n", len(row), string(row))
    }
    // unify lengths
    for i,_ := range canvas {
        fmt.Printf("%v Before: %v\n", len(canvas[i]), string(canvas[i]))
        fmt.Printf("cap: %v\n", cap(canvas[i]))
//        for len(canvas[i]) < maxLen {
//            fmt.Printf("%v During: %v\n", len(canvas[i]), string(canvas[i]))
//            canvas[i] = append(canvas[i], ' ')
//        }
    }
    fmt.Println("After:")
    for _,row := range canvas {
        fmt.Println(string(row))
    }
    return canvas
}

func update(canvas [][]byte, x,y int, ch chan Update) {
    nb := make([]byte, 0)

    nb = append(nb, cell(canvas, x  , y+1))
    nb = append(nb, cell(canvas, x  , y-1))
    nb = append(nb, cell(canvas, x+1, y  ))
    nb = append(nb, cell(canvas, x+1, y+1))
    nb = append(nb, cell(canvas, x+1, y-1))
    nb = append(nb, cell(canvas, x-1, y  ))
    nb = append(nb, cell(canvas, x-1, y+1))
    nb = append(nb, cell(canvas, x-1, y-1))
    
    nb = removeSpaces(nb)
    
    // TODO test double switch
    if canvas[y][x] != ' ' {
        switch {
            case len(nb) <= 1: // under-population
                ch <- Update{x,y,' '}
            case len(nb) <= 3: // lives happily
                ch <- Update{x,y,randEl(nb)}
            default:  // overcrowding
                ch <- Update{x,y,' '}          
        }
    } else if len(nb) == 3 { //reproduction
        ch <- Update{x,y,randEl(nb)}
    } else {
        ch <- Update{x,y,canvas[y][x]}
    }
}

func randEl(slice []byte) byte {
    rand.Seed(time.Now().UnixNano())
    return slice[rand.Intn(len(slice))]
}

func cell(canvas [][]byte, x,y int) byte {
    if y >= 0 && y < len(canvas) && x >= 0 && x < len(canvas[y]) {
        return canvas[y][x]
    }
    return ' '
}

func removeSpaces(slice []byte) []byte {
    for bytes.IndexByte(slice, ' ') != -1 {
        i := bytes.IndexByte(slice, ' ')
        slice = append(slice[:i], slice[i+1:]...)
    }
    return slice
}
