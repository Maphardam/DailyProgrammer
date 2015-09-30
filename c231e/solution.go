package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var offset byte = 48

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //scan input
	state := scanner.Bytes()

	displayState(state)
	for i := 1; i <= 25; i++ {
		state = updateState(state)
		displayState(state)
	}
}

func updateState(state []byte) []byte {
	nState := make([]byte, len(state))

	nState[0] = state[1]
	nState[len(state)-1] = state[len(state)-2]
	for i := 1; i <= len(state)-2; i++ {
		if (state[i-1]-offset == 0 && state[i+1]-offset == 1) ||
			(state[i-1]-offset == 1 && state[i+1]-offset == 0) {
			nState[i] = 1 + offset
		} else {
			nState[i] = 0 + offset
		}
	}

	return nState
}

func displayState(state []byte) {
	for _, el := range state {
		if el-offset == 1 {
			fmt.Print("x")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("\n")
}
