package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
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
    var rightList []int
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
    		rightList = append(rightList, number)
    		list = "left"
    	}
    }
    sort.Ints(leftList)
    sort.Ints(rightList)

    var totalDistance int = 0
    for i := 0; i < len(leftList); i++ {
    	if leftList[i] < rightList[i] {
    		totalDistance += rightList[i] - leftList[i]
    	} else {
    		totalDistance += leftList[i] - rightList[i]
    	}
    }

    fmt.Println(totalDistance)

}