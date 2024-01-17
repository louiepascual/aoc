// https://adventofcode.com/2023/day/2/
// Part one solution only
// I only learned how to code golang for 1 project in Pex, and I'm a heavy Python coder. I'm sure there's a cleaner way to code this 💀
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
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	numbersRegex := regexp.MustCompile(`(\d+)`)
	redRegex := regexp.MustCompile(`(\d+) red`)
	greenRegex := regexp.MustCompile(`(\d+) green`)
	blueRegex := regexp.MustCompile(`(\d+) blue`)

	totalGameId := 0

	for scanner.Scan() {
		isPossible := true
		line := scanner.Text()

		numbers := numbersRegex.FindAllStringSubmatch(line, -1)
		gameId, _ := strconv.Atoi(numbers[0][1])

		redBalls := redRegex.FindAllStringSubmatch(line, -1)
		for _, v := range redBalls {
			reds, _ := strconv.Atoi(v[1])
			if reds > maxRed {
				isPossible = false
				// fmt.Println("reds:", reds, maxRed)
			}
		}

		greenBalls := greenRegex.FindAllStringSubmatch(line, -1)
		for _, v := range greenBalls {
			greens, _ := strconv.Atoi(v[1])
			if greens > maxGreen {
				isPossible = false
				// fmt.Println("greens:", greens, maxGreen)

			}
		}

		blueBalls := blueRegex.FindAllStringSubmatch(line, -1)
		for _, v := range blueBalls {
			blues, _ := strconv.Atoi(v[1])
			if blues > maxBlue {
				isPossible = false
				// fmt.Println("blues:", blues, maxBlue)

			}
		}

		if !isPossible {
			continue
		}
		
		// fmt.Println("Possible: ", gameId)
		totalGameId += gameId
		
	}

	fmt.Println(totalGameId)
}