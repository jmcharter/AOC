package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	crates := make(map[int][]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var stack int
		var whitespace int
		line := scanner.Text()
		if len(line) < 1 {
			continue
		}
		for i, char := range line {
			if char == 'm' { // Process a move command
				qty, from, to := parseMove(line)
				moveStack := make([]string, len(crates[from][:qty]))
				copy(moveStack, crates[from][:qty])
				// Reverse the moveStack
				for i2, j := 0, len(moveStack)-1; i2 < j; i2, j = i2+1, j-1 {
					moveStack[i2], moveStack[j] = moveStack[j], moveStack[i2]
				}
				crates[to] = append(moveStack, crates[to]...)
				crates[from] = crates[from][qty:]

				break
			}
			if char == '[' {
				crates[stack] = append(crates[stack], string(line[i+1]))
				stack++
				whitespace = 0
			}
			if char == ' ' {
				whitespace++
			}
			if whitespace > 3 {
				stack++
				whitespace = 0
			}
		}

	}
	var result string
	for i := 0; i < len(crates); i++ {
		result += crates[i][0]
	}
	fmt.Println(result)

}

func parseMove(line string) (int, int, int) {
	splitline := strings.Split(line, " ")
	quantity, _ := strconv.Atoi(splitline[1])
	from, _ := strconv.Atoi(splitline[3])
	to, _ := strconv.Atoi(splitline[5])

	return quantity, from - 1, to - 1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
