package main

import (
	"bufio"
	"time"
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File not found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read number of lines
	scanner.Scan()
	numberOfLines,_ := strconv.Atoi(scanner.Text())
	
    // read asterisks
	_blueprint      := make([]string, numberOfLines)
	for idx := 0; scanner.Scan(); idx++ {
	    _blueprint[idx] = scanner.Text()
	}
	blueprint := to2dSlice(_blueprint)

	house := generateOutline(blueprint)	
    house = placeDoor(blueprint, house)
	house = placeWindows(blueprint, house)
	house = placeRoof(blueprint, house)
	
	for _,line := range house {
	    fmt.Println(string(line))
	}
}

func to2dSlice(_blueprint []string) [][]bool {
    blueprint := make([][]bool, len(_blueprint))
    rowLength := len(_blueprint[len(_blueprint)-1])
    for i:= range blueprint {
        blueprint[i] = make([]bool, rowLength)
        for lineIdx,char := range _blueprint[i] {
            blueprint[i][lineIdx] = string(char) == "*"
        }
    }
    return blueprint
}

func generateOutline(blueprint [][]bool) [][]rune{

    outline := make([][]rune,(2*len(blueprint)+1))
    
    for idx := range outline {
        outline[idx] = make([]rune, 4*len(blueprint[0])+1)
        for jdx := 0; jdx < 4*len(blueprint[0])+1; jdx++ {
            outline[idx][jdx] = ' '
        }      
    }
    
    
    for rIdx, row := range blueprint {
        for eIdx, entry := range row {
            
            if entry { // is asterisk
                if eIdx == 0 || !row[eIdx-1] { // left side
                    if rIdx == 0 || !blueprint[rIdx-1][eIdx] {
                        // nothing above block
                        outline[2*rIdx][4*eIdx] = '+'
                    } else {
                        outline[2*rIdx][4*eIdx] = '|'
                    }
                    outline[2*rIdx+1][4*eIdx] = '|'
                    if rIdx == len(blueprint)-1 {
                        // bottom corner
                        outline[2*rIdx+2][4*eIdx] = '+'
                    } else {
                        outline[2*rIdx+2][4*eIdx] = '|'
                    }
                }
                if eIdx == len(blueprint[0])-1 || !row[eIdx+1] {
                    // right side
                    if rIdx == 0 || !blueprint[rIdx-1][eIdx] {
                        // nothing above block
                        outline[2*rIdx][4*eIdx+4] = '+'
                    } else {
                        outline[2*rIdx][4*eIdx+4] = '|'
                    }
                    outline[2*rIdx+1][4*eIdx+4] = '|'
                    if rIdx == len(blueprint)-1 {
                        // bottom corner
                        outline[2*rIdx+2][4*eIdx+4] = '+'
                    } else {
                        outline[2*rIdx+2][4*eIdx+4] = '|'
                    } 
                }
                if rIdx == 0 || !blueprint[rIdx-1][eIdx] {
                    // nothing above block
                    if (eIdx == 0 || !row[eIdx-1]) || 
                        (rIdx != 0 && blueprint[rIdx-1][eIdx-1]){
                        // nothing left or is left+above
                        outline[2*rIdx][4*eIdx] = '+'
                    } else {
                        outline[2*rIdx][4*eIdx] = '-'
                    }
                    outline[2*rIdx][4*eIdx+1] = '-'
                    outline[2*rIdx][4*eIdx+2] = '-'
                    outline[2*rIdx][4*eIdx+3] = '-'
                    if (eIdx == len(blueprint[0])-1 || 
                        !row[eIdx+1]) || 
                        (rIdx != 0 && blueprint[rIdx-1][eIdx+1]) {
                        // nothing right
                        outline[2*rIdx][4*eIdx+4] = '+'
                    } else {
                        outline[2*rIdx][4*eIdx+4] = '-'
                    }
                }
                if rIdx == len(blueprint)-1 {
                    // floor
                    if eIdx == 0 {
                        outline[2*rIdx+2][eIdx] = '+'
                    } else {
                        outline[2*rIdx+2][eIdx] = '-'
                    }
                    outline[2*rIdx+2][4*eIdx+1] = '-'
                    outline[2*rIdx+2][4*eIdx+2] = '-'
                    outline[2*rIdx+2][4*eIdx+3] = '-'
                    if eIdx == len(blueprint[0])-1 {
                        outline[2*rIdx+2][4*eIdx+4] = '+'
                    } else {
                        outline[2*rIdx+2][4*eIdx+4] = '-'
                    }
                }
            }


        }
    }
    
    return outline
}

func placeDoor(blueprint [][]bool, house [][]rune) [][]rune {
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(len(blueprint[0]))
    house[2*(len(blueprint)-1)+1][4*n+1] = '|'
    house[2*(len(blueprint)-1)+1][4*n+3] = '|'

    return house
}

func placeWindows(blueprint [][]bool, house [][]rune) [][]rune {
    for rIdx, row := range blueprint {
        for eIdx, entry := range row {
            if entry && house[2*rIdx+1][4*eIdx+1] != '|'{
                rand.Seed(time.Now().UnixNano())
                if rand.Float32() >= 0.5 {
                    house[2*rIdx+1][4*eIdx+2] = 'o'
                }
            }  
        }
    }
    return house
}

func placeRoof(blueprint [][]bool, house [][]rune) [][]rune {
    roofmap := make([][]bool,len(blueprint))
    for idx:= range roofmap {
        roofmap[idx] = make([]bool, len(blueprint[0]))
    }
    for rIdx, row := range blueprint {
        for eIdx, entry := range row {
            if entry && (rIdx == 0 || !blueprint[rIdx-1][eIdx]){
                roofmap[rIdx][eIdx] = true
            }
        }
    }
    
    expandCounter := 0
    for rIdx, row := range roofmap {
        for eIdx := 0; eIdx <= len(row)-1; eIdx++ {
            startIdx := eIdx
            counter  := 0
            
            for eIdx <= len(row)-1 && row[eIdx] {
                counter += 1
                eIdx += 1
            }
            
            // check if space needed

            if len(house) < 2*(len(roofmap)-rIdx)+1+2*counter {

                // expand house map
                newHouse := make([][]rune, 
                                2*(len(roofmap)-rIdx)+1+2*counter)
                expandCounter += (len(newHouse)-len(house))/2
                for idx := len(house)-1; idx >= 0; idx-- {
                    newHouse[idx+len(newHouse)-len(house)] = house[idx]
                }
                for idx := len(newHouse)-len(house)-1; 
                    idx >= 0; idx-- {
                    newHouse[idx] = make([]rune,len(house[0]))
                    for jdx := 0; jdx < len(newHouse[idx]); jdx++ {
                        newHouse[idx][jdx] = ' '
                    }
                }
                house = newHouse
//                for _,line := range house {
//	                fmt.Println(string(line))
//	            }
            }
            
            // place
            if counter > 0 {
                eIdx -= 1
                //fmt.Printf("R: %v, S: %v, E: %v, Exp: %v\n", rIdx, startIdx, eIdx, expandCounter)

                idx := 0
                x1 := 4*startIdx+1
                x2 := 4*eIdx+3
                y := 2*expandCounter-1 + 2*rIdx
                
                for ; idx < eIdx-startIdx; idx++ {

                    house[y-2*idx][x1+2*idx]     = '/'
                    house[y-2*idx-1][x1+2*idx+1] = '/'
                    
                    house[y-2*idx][x2-2*idx]     = '\\'
                    house[y-2*idx-1][x2-2*idx-1] = '\\'
                }
                house[y-2*idx][x1+2*idx] = '/'
                house[y-2*idx][x2-2*idx]     = '\\'
                house[y-2*idx-1][x1+2*idx+1] = 'A'
            }
            
        }
    }
    
    
    return house
                
}
