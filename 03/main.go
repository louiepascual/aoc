// https://adventofcode.com/2023/day/3
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func isNum(char string) bool {
    if _, err := strconv.Atoi(char); err == nil {
        return true
    }
    return false
}

func checkAdjacent(puzzle *[][]string, xCoord int, yCoord int) bool {
    for i := xCoord-1; i <= xCoord+1; i++ {

        if i < 0 || i >= len((*puzzle)){
            continue
        }

        for j := yCoord-1; j <= yCoord+1; j++ {
            if j < 0 || j >= len((*puzzle)[i]){
                continue
            }

            char := (*puzzle)[i][j]
            // fmt.Println((*puzzle)[xCoord][yCoord], i, j, char)

            if !isNum(char) && char != "." {
                return true
            }
        } 
    }

    return false
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    puzzle := make([][]string, 0)

    for scanner.Scan() {
        line := scanner.Text()
        
        newSlice := make([]string, 0)
        for i := range line {
            newSlice = append(newSlice, string(line[i]))
        }

        puzzle = append(puzzle, newSlice)
    }

    totalValue := 0
    for i := range puzzle {
        for j := 0; j < len(puzzle[i]); j++ {
            char := puzzle[i][j]

            if isNum(char) {
                adjFlag := false
                numberSring := ""

                for j < len(puzzle[i]) {
                    if isNum(puzzle[i][j]) {
                        numberSring += puzzle[i][j]
                        
                        if checkAdjacent(&puzzle, i, j) {
                            adjFlag = true
                        }
                    } else {
                        break
                    }
                    j += 1
                }

                fmt.Println(i, j, numberSring, adjFlag)
                if adjFlag {

                    numberInt, err := strconv.Atoi(numberSring)
                    if err != nil {
                        panic(err)
                    }
                    
                    totalValue += numberInt
                }
            }
        }   
    }
    fmt.Println(totalValue)
}