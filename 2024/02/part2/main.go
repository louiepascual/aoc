package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func IsSafe(r []int) bool {
    flow := "increasing"
    if r[0] > r[1] {
        flow = "decreasing"
    }

    safe := true
    for i := 1; i < len(r); i++ {
        diff := r[i] - r[i-1]

        if diff == 0 {
            safe = false
            break
        }

        if diff > 3 || diff < -3 {
            safe = false
            break
        }

        if diff > 0 && flow == "decreasing" {
            safe = false
            break
        }

        if diff < 0 && flow == "increasing" {
            safe = false
            break
        }
    }

    return safe   
}

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

        if IsSafe(report) {
            safeCount += 1
        } else {
            for i := 0; i < len(report); i++ {
                newReport := make([]int, len(report))
                copy(newReport, report)
                newReport = append(newReport[:i], newReport[i+1:]...)
                if IsSafe(newReport) {
                    safeCount += 1
                    break
                }
            }
        }
    }
        fmt.Println(safeCount)
}