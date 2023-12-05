package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func parseNumberWithTextDigits(text string) int {
	replace := map[string]string{
		"one":   "1ne",
		"two":   "2wo",
		"three": "3hree",
		"four":  "4our",
		"five":  "5ive",
		"six":   "6ix",
		"seven": "7even",
		"eight": "8ight",
		"nine":  "9ine",
	}
	for i := 0; i < len(text); i++ {
		for s, r := range replace {
			endIndex := min(i+len(s), len(text))
			if text[i:endIndex] == s {
				text = strings.Replace(text, s, r, 1)
			}
		}
	}
	return parseNumber(text)
}

func main() {
	file, err := os.Open("01/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum, sum2 := 0, 0
	for scanner.Scan() {
		text := scanner.Text()
		sum += parseNumber(text)
		sum2 += parseNumberWithTextDigits(text)
	}

	fmt.Println("The sum of parsed numbers:")
	fmt.Println(sum)
	fmt.Println("The sum of parsed numbers, including textual digits:")
	fmt.Println(sum2)
}
