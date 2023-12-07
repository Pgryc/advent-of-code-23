package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func loadSchematic(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}
	return schematic
}

func isPartNo(surroundings []string) bool {
	for _, text := range surroundings {
		for _, character := range text {
			if string(character) != "." {
				return true
			}
		}
	}
	return false
}

func countNumbers(schematic []string) int {
	sum := 0
	for i, line := range schematic {
		for j := 0; j < len(line); j++ {
			if unicode.IsDigit(rune(line[j])) {
				var surroundings []string
				lastNumberDigit := j
				for lastNumberDigit+1 < len(line) && unicode.IsDigit(rune(line[lastNumberDigit+1])) {
					lastNumberDigit++
				}

				if j > 0 {
					surroundings = append(surroundings, string(line[j-1]))
				}
				if lastNumberDigit < len(line)-1 {
					surroundings = append(surroundings, string(line[lastNumberDigit+1]))
				}
				if i > 0 {
					surroundings = append(surroundings, schematic[i-1][max(0, j-1):min(lastNumberDigit+2, len(line))])
				}
				if i < len(schematic)-1 {
					surroundings = append(surroundings, schematic[i+1][max(0, j-1):min(lastNumberDigit+2, len(line))])
				}

				if isPartNo(surroundings) {
					partNo, err := strconv.Atoi(string(line[j : lastNumberDigit+1]))
					if err != nil {
						fmt.Println(err)
					}
					sum += partNo
				}
				j = lastNumberDigit
			}
		}
	}
	return sum
}

func main() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	schematic := loadSchematic(file)

	sum := countNumbers(schematic)
	fmt.Println(sum)
}
