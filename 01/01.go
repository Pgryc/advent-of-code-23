package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func parseNumber(text string) int {
	expr := regexp.MustCompile(`[0-9]`)
	indexes := expr.FindAllStringIndex(text, -1)
	parsed := string(
		[]byte{
			text[indexes[0][0]],
			text[indexes[len(indexes)-1][0]],
		})
	result, err := strconv.Atoi(parsed)
	check(err)
	return result
}

func main() {
	file, err := os.Open("01/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		sum += parseNumber(scanner.Text())
	}

	fmt.Println("The sum of parsed numbers:")
	fmt.Println(sum)
}
