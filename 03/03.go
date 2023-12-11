package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type col struct {
	start int
	end   int
}

type number struct {
	value int
	row   int
	col   col
}

type symbol struct {
	row                int
	col                int
	neighboringNumbers []number
}

func calculateSums(symbols []symbol) (sum int, sum2 int) {
	sum = 0
	sum2 = 0
	for _, symbol := range symbols {
		for _, number := range symbol.neighboringNumbers {
			sum += number.value
		}
		if len(symbol.neighboringNumbers) == 2 {
			sum2 += symbol.neighboringNumbers[0].value * symbol.neighboringNumbers[1].value
		}

	}
	return
}

func findNeighboring(numbers []number, symbols []symbol) ([]number, []symbol) {
	for i := 0; i < len(symbols); i++ {
		for j := 0; j < len(numbers); j++ {
			if (numbers[j].row-1 <= symbols[i].row) &&
				(symbols[i].row <= numbers[j].row+1) &&
				(numbers[j].col.start-1 <= symbols[i].col) &&
				(symbols[i].col <= numbers[j].col.end+1) {
				symbols[i].neighboringNumbers = append(symbols[i].neighboringNumbers, numbers[j])
			}
		}
	}
	return numbers, symbols
}

func parseSchematic(schematic []string) (numbers []number, symbols []symbol) {
	for i, line := range schematic {
		for j := 0; j < len(line); j++ {
			isDigit := unicode.IsDigit(rune(line[j]))
			isDot := string(rune(line[j])) == "."

			if isDigit {
				lastNumberDigit := j
				for lastNumberDigit+1 < len(line) && unicode.IsDigit(rune(line[lastNumberDigit+1])) {
					lastNumberDigit++
				}
				nValue, err := strconv.Atoi(string(line[j : lastNumberDigit+1]))
				if err != nil {
					fmt.Println(err)
					return
				}
				n := number{
					value: nValue,
					row:   i,
					col:   col{start: j, end: lastNumberDigit},
				}
				numbers = append(numbers, n)
				j = lastNumberDigit
			} else if !isDigit && !isDot {
				s := symbol{
					row: i,
					col: j,
				}
				symbols = append(symbols, s)
			}
		}
	}

	return
}

func loadSchematic(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}
	return schematic
}

func main() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	schematic := loadSchematic(file)
	numbers, symbols := parseSchematic(schematic)
	fmt.Println(numbers)
	fmt.Println(symbols)

	numbers, symbols = findNeighboring(numbers, symbols)

	sum, sum2 := calculateSums(symbols)
	fmt.Println(sum)
	fmt.Println(sum2)
}
