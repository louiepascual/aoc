package main

import (
	"fmt"
	"os"
	"bufio"
    "strings"
    "strconv"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    safeCount := 0
    for scanner.Scan() {
        line := scanner.Text()
        
        lineScanner := bufio.NewScanner(strings.NewReader(line))
        lineScanner.Split(bufio.ScanWords)

        var report []int
        for lineScanner.Scan() {
            number, err := strconv.Atoi(lineScanner.Text())
            if err != nil {
                panic(err)
            }

            report = append(report, number)
        }

        flow := "increasing"
        if report[0] > report[1] {
            flow = "decreasing"
        }

        isSafe := true
        for i := 1; i < len(report); i++ {
            diff := report[i] - report[i-1]

            if diff == 0 {
                isSafe = false
                break
            }

            if diff > 3 || diff < -3 {
                isSafe = false
                break
            }

            if diff > 0 && flow == "decreasing" {
                isSafe = false
                break
            }

            if diff < 0 && flow == "increasing" {
                isSafe = false
                break
            }
        }

        if isSafe {
            safeCount += 1
        }
    }
        fmt.Println(safeCount)
}