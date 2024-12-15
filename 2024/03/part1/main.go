package main

import (
	"bufio"
    "fmt"
	"os"
    "regexp"
    "strconv"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        matches :=  r.FindAllStringSubmatch(line, -1)

        for i := range matches {
            left, _ := strconv.Atoi(matches[i][1])
            right, _ := strconv.Atoi(matches[i][2])
            total += left * right
        }
    }
    fmt.Println(total)
}