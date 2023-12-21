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

func countCardPoints(card scaratchcard) (matches int, points int) {
	matching := simpleGeneric(card.winning, card.numbers)
	if len(matching) > 0 {
		points = int(math.Pow(2, float64(len(matching)-1)))
	} else {
		points = 0
	}
	matches = len(matching)
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
	// go comment
	// start with array of 1 card each
	// for each card, you get 1 of <points> next cards
	//
	// how much are there at the end

	scanner := bufio.NewScanner(file)

	var cards []scaratchcard

	for scanner.Scan() {
		cards = append(cards, parseCardData(scanner.Text()))
	}
	size := len(cards)
	numberOfCards := make([]int, size)

	for i := 0; i < len(numberOfCards); i++ {
		numberOfCards[i] = 1
	}

	sum, sum2 := 0, 0
	for i, card := range cards {
		matches, currentCardPoints := countCardPoints(card)
		sum += currentCardPoints

		for j := 1; j <= matches && j+i < len(cards); j++ {
			numberOfCards[i+j] += numberOfCards[i]
		}
	}

	for i := 0; i < len(numberOfCards); i++ {
		sum2 += numberOfCards[i]
	}

	fmt.Println(cards)
	fmt.Println(sum)
	fmt.Println(sum2)
}
