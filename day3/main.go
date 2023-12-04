package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"golang.org/x/exp/maps"
)

type CellType int
type Loc struct {
	Row, Col int
}

const (
	NUMBER CellType = iota
	SYMBOL
	// ASTERISK
	EMPTY
)

type Cell struct {
	Type     CellType
	Contents rune
	Location Loc
}

type Number struct {
	Val   int
	Start Cell
	End   Cell
}

type Symbol struct {
	Cell Cell
}

/*
Grid Layout: R x C
[0,0] [0,1] [0,2]
[1,0] [1,1] [1,2]
[2,0] [2,1] [2,2]
*/

type Grid [][]Cell

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

	// Read each line of the file and process
	row := 0
	grid := make(Grid, 0)
	for fileScanner.Scan() {
		/*
		   467..114..
		   ...*......
		   ..35..633.
		*/
		lineStr := fileScanner.Text()
		grid = append(grid, make([]Cell, len(lineStr)))
		for col, char := range lineStr {
			cell := CharToCell(char, row, col)
			grid[row][col] = cell
		}
		row++
	}

	numbers, symbols := grid.Search()

	// Part 1
	sum := 0
	for _, n := range numbers {
		adjacencies := n.AdjacentToType(grid, SYMBOL)
		// adjacentAst, ast := n.AdjacentToType(grid, ASTERISK)
		if len(adjacencies) > 0 {
			sum += n.Val
			fmt.Printf("Part %d at (%d, %d) is next to a symbol (%s)\n", n.Val, n.Start.Location.Row, n.Start.Location.Col, string(adjacencies[0].Contents))
		}
	}

	fmt.Printf("Part Sum: %d\n", sum)
	// End of Part 1

	// Part 2

	// Map cells to the Number they correspond to
	numberMap := make(map[Loc]Number)
	for _, n := range numbers {
		numberMap[n.Start.Location] = n
		numberMap[n.End.Location] = n
	}

	gearRatioSum := 0
	for _, sym := range symbols {
		if sym.Cell.Contents != '*' {
			continue
		} else {
			uniqueNumberMap := make(map[Number]bool)
			adjacentNumbers := sym.Cell.AdjacentToType(grid, NUMBER)
			for _, cell := range adjacentNumbers {
				if num, ok := numberMap[cell.Location]; ok {
					uniqueNumberMap[num] = true
				}
			}

			uniqueNumbers := maps.Keys(uniqueNumberMap)
			if len(uniqueNumbers) != 2 {
				continue
			} else {
				gearRatio := uniqueNumbers[0].Val * uniqueNumbers[1].Val
				gearRatioSum += gearRatio

				fmt.Printf("Gear using parts %d and %d gives a gear ratio of %d\n", uniqueNumbers[0].Val, uniqueNumbers[1].Val, gearRatio)
			}
		}
	}

	fmt.Printf("Sum of all gear ratios: %d\n", gearRatioSum)
	// End of Part 2

}

// Up one Row, same Col
func North(g Grid, c Cell) Cell {
	l := c.Location
	if l.Row == 0 {
		return g[l.Row][l.Col]
	} else {
		return g[l.Row-1][l.Col]
	}
}

// Down one Row, same Col
func South(g Grid, c Cell) Cell {
	l := c.Location
	if l.Row == len(g)-1 {
		return g[l.Row][l.Col]
	} else {
		return g[l.Row+1][l.Col]
	}
}

// Same Row, left one Col
func East(g Grid, c Cell) Cell {
	l := c.Location
	if l.Col == len(g[l.Row])-1 {
		return g[l.Row][l.Col]
	} else {
		return g[l.Row][l.Col+1]
	}
}

// Same Row, right one Col
func West(g Grid, c Cell) Cell {
	l := c.Location
	if l.Col == 0 {
		return g[l.Row][l.Col]
	} else {
		return g[l.Row][l.Col-1]
	}
}

func CharToCell(char rune, row, col int) (c Cell) {
	c.Contents = char
	c.Location = Loc{row, col}

	switch char {
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':
		c.Type = NUMBER
	case '.':
		c.Type = EMPTY
	// case '*':
	// 	c.Type = ASTERISK
	default:
		c.Type = SYMBOL
	}

	return
}

func (c Cell) IsNumber() bool {
	return c.Type == NUMBER
}

func (c Cell) IsSymbol() bool {
	return c.Type == SYMBOL
}

func (grid Grid) Search() (numbers []Number, symbols []Symbol) {
	for _, cells := range grid {
		for _, cell := range cells {
			switch cell.Type {
			case EMPTY:
				continue
			case SYMBOL:
				symbols = append(symbols, Symbol{cell})
			case NUMBER:
				if !grid.IsLeadingDigit(cell) {
					continue
				}

				secondDigit := East(grid, cell)
				if !secondDigit.IsNumber() || secondDigit == cell {
					// One unique digit
					num := RunesToInt(cell.Contents)
					numbers = append(numbers, Number{num, cell, cell})
					continue
				} else {
					thirdDigit := East(grid, secondDigit)
					if !thirdDigit.IsNumber() || thirdDigit == secondDigit {
						// Two unique digits
						num := RunesToInt(secondDigit.Contents, cell.Contents)
						numbers = append(numbers, Number{num, cell, secondDigit})
						continue
					} else {
						// Three unique digits
						num := RunesToInt(thirdDigit.Contents, secondDigit.Contents, cell.Contents)
						numbers = append(numbers, Number{num, cell, thirdDigit})
						continue
					}
				}
			}
		}
	}

	return
}

func (c Cell) AdjacentToType(grid Grid, t CellType) (adjacent []Cell) {
	adjacentCells := append(grid.AdjacencyList(c))

	for _, cell := range adjacentCells {
		if cell.Type == t {
			adjacent = append(adjacent, cell)
		}
	}

	return
}

func (n Number) AdjacentToType(grid Grid, t CellType) (adjacent []Cell) {
	return append(n.Start.AdjacentToType(grid, t), n.End.AdjacentToType(grid, t)...)
}

func (grid Grid) AdjacencyList(c Cell) []Cell {
	return []Cell{
		North(grid, c),
		North(grid, East(grid, c)),
		East(grid, c),
		South(grid, East(grid, c)),
		South(grid, c),
		South(grid, West(grid, c)),
		West(grid, c),
		North(grid, West(grid, c)),
	}
}

func (grid Grid) IsLeadingDigit(cell Cell) bool {
	if !cell.IsNumber() {
		return false
	}

	if West(grid, cell).IsNumber() && West(grid, cell) != cell {
		// This is already part of a number, ignore
		return false
	}

	return true
}

func RunesToInt(runes ...rune) int {
	n := 0

	for i, r := range runes {
		asInt, _ := strconv.Atoi(string(r))
		n += int(float64(asInt) * math.Pow(10, float64(i)))
	}

	return n
}
