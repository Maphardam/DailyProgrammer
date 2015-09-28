package main

import (
    "log"
    "os"
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "math"
)

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
	nv,_ := strconv.Atoi(input[0])
	nf,_ := strconv.Atoi(input[1])
	
	fmt.Printf("%v %v\n", nv, nf)
	
	for f1 := int(math.Pow(10.0, float64(nf-1))); 
	    f1  < int(math.Pow(10.0, float64(nf))); f1++ {
	    for f2 := f1; 
	        f2 < int(math.Pow(10.0, float64(nf+1))); 
	        f2++ {
	        if !(f1%10 == 0 && f2%10 == 0) && 
	            numDigits(f1*f2)== nv && goodDigits(f1,f2){
	            fmt.Printf("%v=%v*%v\n", f1*f2, f1, f2)
            }
	    }
	}
	
}

func numDigits(n int) (count int){
	for n!=0 {
        n /= 10;
        count++;
    }
    return
}

func goodDigits(f1,f2 int) bool {
    // vampire digits
    v := f1*f2
    vd := make([]int, 0)
    for v > 0 {
        vd = append(vd, v%10)
        v /= 10 
    }
    var is bool
    // check fang 1
    for f1 > 0 {
        vd, is = inVampire(f1%10, vd)
        if !is {
            return false
        }
        f1 /= 10
    }
    
    // check fang 2
    for f2 > 0 {
        vd, is = inVampire(f2%10, vd)
        if !is {
            return false
        }
        f2 /= 10
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
