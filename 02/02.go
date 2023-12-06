package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parse(text string) int {
	gameHeader := strings.Split(text, ":")[0]
	gamesData := strings.Split(text, ":")[1]

	gameId, err := strconv.Atoi(strings.Split(gameHeader, " ")[1])
	check(err)
	games := strings.Split(gamesData, ";")

	gamePossible := true
	for _, g := range games {
		cubes := strings.Split(g, ",")
		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			cubeCount, err := strconv.Atoi(strings.Split(cube, " ")[0])
			check(err)
			cubeColor := strings.Split(cube, " ")[1]

			if cubeCount > maxCubes[cubeColor] {
				gamePossible = false
			}
		}
	}
	if gamePossible {
		return gameId
	}
	return 0
}

func power(text string) int {
	min := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	gamesData := strings.Split(text, ":")[1]
	games := strings.Split(gamesData, ";")

	for _, g := range games {
		cubes := strings.Split(g, ",")
		for _, cube := range cubes {
			cube = strings.Trim(cube, " ")
			cubeCount, err := strconv.Atoi(strings.Split(cube, " ")[0])
			check(err)
			cubeColor := strings.Split(cube, " ")[1]

			if cubeCount > min[cubeColor] {
				min[cubeColor] = cubeCount
			}
		}
	}
	return min["red"] * min["green"] * min["blue"]
}

func main() {
	file, err := os.Open("02/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum, sum2 := 0, 0
	for scanner.Scan() {
		sum += parse(scanner.Text())
		sum2 += power(scanner.Text())
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
