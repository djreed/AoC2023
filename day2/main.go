package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var DIGITS = regexp.MustCompile("[0-9]+")

type MarbleSet struct {
	Blue, Red, Green int
}

func main() {
	if len(os.Args) <= 1 {
		panic("Must pass in input file to be processed, e.x. `day2 example_p1`")
	}

	filename := os.Args[1]
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Part 1 Marble Restrictions
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	Part1Restriction := MarbleSet{Red: 12, Green: 13, Blue: 14}

	sum := 0
	powerSum := 0

	// Read each line of the file and process
	for fileScanner.Scan() {
		// Game {N}: [X blue, Y red, Z green; ...]
		lineStr := fileScanner.Text()
		game, rounds := FilterGameRounds(lineStr)

		// Part 1
		valid := true
		for _, round := range rounds {
			if !ValidRound(Part1Restriction, round) {
				fmt.Printf("Game %d invalid due to round: %d Blue, %d Red, %d Green\n", game, round.Blue, round.Red, round.Green)
				valid = false
				continue
			}
		}

		if valid {
			sum += game
		}
		// End of Part 1

		// Part 2
		minBlue := -1
		minRed := -1
		minGreen := -1

		for _, round := range rounds {
			if round.Blue > minBlue {
				minBlue = round.Blue
			}

			if round.Red > minRed {
				minRed = round.Red
			}

			if round.Green > minGreen {
				minGreen = round.Green
			}
		}

		powerSum = powerSum + (minBlue * minRed * minGreen)
		// End of Part 2
	}

	fmt.Printf("Sum == %d\n", sum)
	fmt.Printf("PowerSum == %d\n", powerSum)
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func FilterGameRounds(str string) (int, []MarbleSet) {
	splitStr := strings.Split(str, ":")
	gameStr := splitStr[0]
	gameDigit, _ := strconv.Atoi(DIGITS.FindString(gameStr))
	roundString := splitStr[1]
	rounds := strings.Split(roundString, ";")
	marbleSets := make([]MarbleSet, 0)

	for _, r := range rounds {
		regBlue := regexp.MustCompile("\\d+ blue")
		regRed := regexp.MustCompile("\\d+ red")
		regGreen := regexp.MustCompile("\\d+ green")

		blueMarbleStr := DIGITS.FindString(regBlue.FindString(r))
		redMarbleStr := DIGITS.FindString(regRed.FindString(r))
		greenMarbleStr := DIGITS.FindString(regGreen.FindString(r))

		blueMarbleInt, _ := strconv.Atoi(blueMarbleStr)
		redMarbleInt, _ := strconv.Atoi(redMarbleStr)
		greenMarbleInt, _ := strconv.Atoi(greenMarbleStr)

		roundMarbles := MarbleSet{
			Blue:  blueMarbleInt,
			Red:   redMarbleInt,
			Green: greenMarbleInt,
		}

		marbleSets = append(marbleSets, roundMarbles)
	}

	return gameDigit, marbleSets
}

func ValidRound(restriction, round MarbleSet) bool {
	return round.Blue <= restriction.Blue &&
		round.Red <= restriction.Red &&
		round.Green <= restriction.Green
}
