package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Scan() //scan line number line
	ln, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	
	var s string
	for i := 1; i <= ln; i++ {
	    scanner.Scan()
	    tmp := scanner.Text()
	    tmp  = strings  .ToLower(tmp)
	    tmp  = strings.Replace(tmp, " ", "", -1)
	    tmp  = strings.Replace(tmp, ".", "", -1)
	    tmp  = strings.Replace(tmp, ",", "", -1)
	    tmp  = strings.Replace(tmp, ":", "", -1)
	    tmp  = strings.Replace(tmp, ";", "", -1)
	    tmp  = strings.Replace(tmp, "!", "", -1)
	    tmp  = strings.Replace(tmp, "?", "", -1)
	    tmp  = strings.Replace(tmp, "\n", "", -1)
	    s   += tmp
	}
	
	if s == reverse(s) {
	    fmt.Println("Palindrome")
	} else {
	    fmt.Println("Not a Palindrome")
	}     
}

func reverse(s string) string {
    n := 0
    rune := make([]rune, len(s))
    for _, r := range s { 
            rune[n] = r
            n++
    } 
    rune = rune[0:n]
    // Reverse 
    for i := 0; i < n/2; i++ { 
            rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
    } 
     
    return string(rune)   
}
