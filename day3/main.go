package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	num   int
	start Cell
	end   Cell
}

type Symbol struct {
	c Cell
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

	sum := 0
	numbers, _ /*symbols*/ := grid.Search()
	for _, n := range numbers {
		adjacent, sym := n.AdjacentToType(grid, SYMBOL)
		// adjacentAst, ast := n.AdjacentToType(grid, ASTERISK)
		if adjacent {
			sum += n.num
			fmt.Printf("Part %d at (%d, %d) is next to a symbol (%s)\n", n.num, n.start.Location.Row, n.start.Location.Col, sym)
		}
	}

	fmt.Printf("Part Sum: %d\n", sum)

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
	return c.Type == SYMBOL /*|| c.Type == ASTERISK*/
}

func (grid Grid) CellAt(l Loc) Cell {
	return grid[l.Row][l.Col]
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
				if !grid.IsStartOfNumber(cell) {
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

func (n Number) AdjacentToType(grid Grid, t CellType) (bool, string) {
	adjacencies := append(grid.AdjacencyList(n.start), grid.AdjacencyList(n.end)...)

	for _, cell := range adjacencies {
		if cell.Type == t {
			return true, string(cell.Contents)
		}
	}

	return false, ""
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

func (grid Grid) IsStartOfNumber(cell Cell) bool {
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
