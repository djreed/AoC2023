package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	Id      string
	Winning []int
	Numbers []int
}

func main() {
	if len(os.Args) <= 1 {
		panic("Must pass in input file to be processed, e.x. `day4 example_p1`")
	}

	filename := os.Args[1]
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cards := make([]ScratchCard, 0)
	// Read each line of the file and process
	for fileScanner.Scan() {
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		lineStr := fileScanner.Text()
		card := ParseCard(lineStr)
		cards = append(cards, card)
	}

	// Part 1
	pointTotal := 0
	for _, card := range cards {
		matches := MatchingNumbers(card)
		points := CardValue(card)
		pointTotal += points

		fmt.Printf("%s: %d matching numbers => %d points\n", card.Id, matches, points)
	}
	fmt.Println("Total card values: ", pointTotal)
	// End of Part 1

	// Part 2
	totalCards := 0
	cardCopyCount := make(map[int]int)
	for cardIdx, card := range cards {
		copiesOfCard := 1 + cardCopyCount[cardIdx]
		totalCards += copiesOfCard
		numMatches := MatchingNumbers(card)

		for extraCardIdx := cardIdx + 1; extraCardIdx <= cardIdx+numMatches; extraCardIdx++ {
			cardCopyCount[extraCardIdx] += copiesOfCard
		}

		fmt.Printf("Processed card [%d] %d times; %d matches\n", cardIdx+1, 1+cardCopyCount[cardIdx], numMatches)
	}

	fmt.Printf("Processed %d total cards\n", totalCards)
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func ParseCard(str string) ScratchCard {
	var card ScratchCard

	// "Card N" and "41 48 83 ..."
	idAndValues := strings.Split(str, ":")
	card.Id = idAndValues[0]

	// "41 48 83 ..." and "83 86  6 ..."
	winningAndNumbers := strings.Split(idAndValues[1], "|")

	// "41" "48" "84" ...
	winningNumbers := strings.Split(winningAndNumbers[0], " ")

	// "83" "86" "" "6" ...
	cardNumbers := strings.Split(winningAndNumbers[1], " ")

	for _, win := range winningNumbers {
		if win == "" {
			continue
		}
		asInt, _ := strconv.Atoi(win)
		card.Winning = append(card.Winning, asInt)
	}

	for _, num := range cardNumbers {
		if num == "" {
			continue
		}
		asInt, _ := strconv.Atoi(num)
		card.Numbers = append(card.Numbers, asInt)
	}

	return card
}

func MatchingNumbers(card ScratchCard) int {
	matches := 0
	for _, cardVal := range card.Numbers {
		for _, winner := range card.Winning {
			if cardVal == winner {
				matches++
			}
		}
	}
	return matches
}

func CardValue(card ScratchCard) int {
	matches := MatchingNumbers(card)
	if matches == 0 {
		return 0
	} else {
		score := 1
		for i := 1; i < matches; i++ {
			score *= 2
		}
		return score
	}
}
