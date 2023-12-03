package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(sumOfViableGames(lines))
	fmt.Println(sumOfMinColorProducts(lines))
}

func sumOfViableGames(lines []string) int {
	sum := 0
	for ln, line := range lines {
		sets := getSets(line)
		viable := true
		for _, color := range sets {
			colors, err := getColors(color)
			if err != nil {
				log.Fatal(err)
			}
			if colors["red"] > 12 || colors["blue"] > 14 || colors["green"] > 13 {
				viable = false
				break
			}
		}
		if viable {
			sum += ln + 1
		}
	}
	return sum
}

func sumOfMinColorProducts(lines []string) int {
	sum := 0
	for _, line := range lines {
		sets := getSets(line)
		minColors := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, color := range sets {
			colors, err := getColors(color)
			if err != nil {
				log.Fatal(err)
			}
			if colors["red"] > minColors["red"] {
				minColors["red"] = colors["red"]
			}
			if colors["blue"] > minColors["blue"] {
				minColors["blue"] = colors["blue"]
			}
			if colors["green"] > minColors["green"] {
				minColors["green"] = colors["green"]
			}
		}
		power := minColors["red"] * minColors["blue"] * minColors["green"]
		sum += power
	}
	return sum
}

func getSets(s string) []string {
	s = strings.Split(s, ":")[1]
	return strings.Split(s, ";")
}

func getColors(s string) (map[string]int, error) {
	colors := map[string]int{}
	colorSplit := strings.Split(s, ",")
	for _, c := range colorSplit {
		c = strings.TrimSpace(c)
		s := strings.Split(c, " ")
		amnt, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, err
		}
		colors[s[1]] += amnt
	}
	return colors, nil
}
