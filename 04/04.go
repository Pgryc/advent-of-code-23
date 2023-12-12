package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type scaratchcard struct {
	name    string
	numbers []int
	winning []int
}

// from github.com/juliangruber/go-intersect
func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

func simpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func countCardPoints(card scaratchcard) (points int) {
	matching := simpleGeneric(card.winning, card.numbers)
	if len(matching) > 0 {
		points = int(math.Pow(2, float64(len(matching)-1)))
	} else {
		points = 0
	}
	return
}

func parseNumbers(text string) (numbers []int) {
	numbersText := strings.Split(strings.Trim(text, " "), " ")
	for _, numberText := range numbersText {
		if numberText != "" {
			number, err := strconv.Atoi(strings.Trim(numberText, " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			numbers = append(numbers, number)
		}
	}
	return
}

func parseCardData(text string) (card scaratchcard) {
	cardName := strings.Split(text, ":")[0]
	text = strings.Split(text, ":")[1]
	winningText := strings.Split(text, "|")[0]
	numbersText := strings.Split(text, "|")[1]
	winning := parseNumbers(winningText)
	numbers := parseNumbers(numbersText)

	card = scaratchcard{
		cardName,
		numbers,
		winning,
	}
	return
}

func main() {
	file, err := os.Open("04/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var cards []scaratchcard

	for scanner.Scan() {
		cards = append(cards, parseCardData(scanner.Text()))
	}
	sum := 0
	for _, card := range cards {
		sum += countCardPoints(card)
	}

	fmt.Println(cards)
	fmt.Println(sum)
}
