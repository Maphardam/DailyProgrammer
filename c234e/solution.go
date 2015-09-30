package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var nv int

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("File not found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	nv, _ = strconv.Atoi(input[0])
	nf, _ := strconv.Atoi(input[1])

	mult(1, 0, 10, make([]int, nf))

}

func mult(v, fi, start int, fs []int) {
	for f := start; f <= 99; f++ {
		fs[fi] = f
		if fi == len(fs)-1 {
			if !tooManyZeros(fs) && numDigits(v*f) == nv &&
				goodDigits(v*f, fs) {
				fmt.Printf("%v=", v*f)
				for _, e := range fs[:len(fs)-1] {
					fmt.Printf("%v*", e)
				}
				fmt.Printf("%v\n", fs[len(fs)-1])
			}
		} else {
			mult(v*f, fi+1, f, fs)
		}
	}
}

func numDigits(n int) (count int) {
	for n != 0 {
		n /= 10
		count++
	}
	return
}

func tooManyZeros(slice []int) bool {
	oneZero := false
	for _, el := range slice {
		if el%10 == 0 && !oneZero {
			oneZero = true
		} else if el%10 == 0 && oneZero {
			return true
		}
	}
	return false
}

func goodDigits(v int, fs []int) bool {
	// vampire digits

	vd := make([]int, 0)
	for v > 0 {
		vd = append(vd, v%10)
		v /= 10
	}
	var is bool
	// check fang 1
	for _, f := range fs {
		for f > 0 {
			vd, is = inVampire(f%10, vd)
			if !is {
				return false
			}
			f /= 10
		}
	}

	return true
}

func inVampire(a int, list []int) ([]int, bool) {
	for i, b := range list {
		if b == a {
			return append(list[:i], list[i+1:]...), true
		}
	}
	return list, false
}
