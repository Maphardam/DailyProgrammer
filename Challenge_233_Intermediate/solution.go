package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type Update struct {
	X, Y int
	Char byte
}

func main() {
	canvas := read(os.Args[1])

	ch := make(chan Update)
	for y, _ := range canvas {
		for x, _ := range canvas[y] {
			go update(canvas, x, y, ch)
		}
	}

	nCanvas := make([][]byte, len(canvas))
	for i, _ := range canvas {
		nCanvas[i] = make([]byte, len(canvas[i]))
	}

	upd := Update{}
	for i := 0; i < len(canvas)*len(canvas[0]); i++ {
		upd = <-ch
		nCanvas[upd.Y][upd.X] = upd.Char
	}

	for _, row := range nCanvas {
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// unify lengths
	for i, _ := range canvas {
		tmp := make([]byte, 0, maxLen)
		tmp = append(tmp, canvas[i]...)
		canvas[i] = append(tmp, bytes.Repeat([]byte(" "), maxLen-len(canvas[i]))...)
	}

	return canvas
}

func update(canvas [][]byte, x, y int, ch chan Update) {
	nb := make([]byte, 0)

	nb = append(nb, cell(canvas, x, y+1))
	nb = append(nb, cell(canvas, x, y-1))
	nb = append(nb, cell(canvas, x+1, y))
	nb = append(nb, cell(canvas, x+1, y+1))
	nb = append(nb, cell(canvas, x+1, y-1))
	nb = append(nb, cell(canvas, x-1, y))
	nb = append(nb, cell(canvas, x-1, y+1))
	nb = append(nb, cell(canvas, x-1, y-1))

	nb = removeSpaces(nb)

	switch {
	case canvas[y][x] != ' ':
		switch {
		case len(nb) <= 1: // under-population
			ch <- Update{x, y, ' '}
		case len(nb) <= 3: // lives happily
			ch <- Update{x, y, randEl(nb)}
		default: // overcrowding
			ch <- Update{x, y, ' '}
		}
	case len(nb) == 3: //reproduction
		ch <- Update{x, y, randEl(nb)}
	default:
		ch <- Update{x, y, canvas[y][x]}
	}
}

func randEl(slice []byte) byte {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

func cell(canvas [][]byte, x, y int) byte {
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
