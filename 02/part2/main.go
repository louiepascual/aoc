// https://adventofcode.com/2023/day/2/
// Part two solution only
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

    redRegex := regexp.MustCompile(`(\d+) red`)
    greenRegex := regexp.MustCompile(`(\d+) green`)
    blueRegex := regexp.MustCompile(`(\d+) blue`)

    totalPower := 0

    for scanner.Scan() {
        maxRed := 1
        maxGreen := 1
        maxBlue := 1
        
        line := scanner.Text()

        redBalls := redRegex.FindAllStringSubmatch(line, -1)
        for _, v := range redBalls {
            reds, _ := strconv.Atoi(v[1])
            if reds > maxRed {
                maxRed = reds
            }
        }

        greenBalls := greenRegex.FindAllStringSubmatch(line, -1)
        for _, v := range greenBalls {
            greens, _ := strconv.Atoi(v[1])
            if greens > maxGreen {
                maxGreen = greens
            }
        }

        blueBalls := blueRegex.FindAllStringSubmatch(line, -1)
        for _, v := range blueBalls {
            blues, _ := strconv.Atoi(v[1])
            if blues > maxBlue {
                maxBlue = blues
            }
        }

        totalPower += (maxRed * maxGreen * maxBlue)    
    }

    fmt.Println(totalPower)
}