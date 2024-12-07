package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var leftList []int
    rightMap := make(map[int]int)

    var list string = "left"
    var number int

    for scanner.Scan() {
    	number, err = strconv.Atoi(scanner.Text())
    	if err != nil {
    		panic(err)
    	}

    	if list == "left" {
    		leftList = append(leftList, number)
    		list = "right"
    	} else {
    		rightMap[number] += 1
    		list = "left"
    	}
    }

    similarityScore := 0
    for i := 0; i < len(leftList); i++ {
        similarityScore += leftList[i] * rightMap[leftList[i]]
    }

    fmt.Println(similarityScore)
}