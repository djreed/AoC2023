package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Must pass in input file to be processed, e.x. `day1 example_p1`")
	}

	filename := os.Args[1]
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	// Read each line of the file and process
	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		converted := ConvertWordsToDigits(lineStr)
		digits := FilterDigits(converted)
		first, last := FirstLastDigits(digits)
		combined := CombineDigits(first, last)

		fmt.Printf("%s => %s || %d + %d = %d\n", lineStr, converted, first, last, combined)

		sum += combined
	}

	fmt.Printf("Sum == %d\n", sum)
}

func ConvertWordsToDigits(str string) string {
	finalStr := str
	finalStr = strings.Replace(finalStr, "one", "o1e", -1)
	finalStr = strings.Replace(finalStr, "two", "t2o", -1)
	finalStr = strings.Replace(finalStr, "three", "th3ee", -1)
	finalStr = strings.Replace(finalStr, "four", "fo4ur", -1)
	finalStr = strings.Replace(finalStr, "five", "fi5ve", -1)
	finalStr = strings.Replace(finalStr, "six", "s6x", -1)
	finalStr = strings.Replace(finalStr, "seven", "se7en", -1)
	finalStr = strings.Replace(finalStr, "eight", "ei8ht", -1)
	finalStr = strings.Replace(finalStr, "nine", "ni9ne", -1)
	return finalStr
}

func FilterDigits(str string) []int {
	// Regex for numbers
	re := regexp.MustCompile("[0-9]")

	strs := re.FindAllString(str, -1)

	digits := make([]int, 0)
	for _, s := range strs {
		d, _ := strconv.Atoi(s)
		digits = append(digits, d)
	}

	return digits
}

func FirstLastDigits(list []int) (int, int) {
	return list[0], list[len(list)-1]
}

func CombineDigits(d1, d2 int) int {
	combined, _ := strconv.Atoi(fmt.Sprintf("%d%d", d1, d2))
	return combined
}
