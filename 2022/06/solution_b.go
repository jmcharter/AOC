package main

import (
	"bufio"
	"os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	var offset int = 14
	// var idx int = 13
	match := false
	for idx, _ := range text {
		match = false
		if idx+offset > len(text) {
			diff := (idx + offset) - len(text)
			offset -= diff
		}
		for p1 := idx; p1 < idx+offset; p1++ {
			for p2 := p1 + 1; p2 < idx+offset; p2++ {
				if text[p1] == text[p2] {
					match = true
					break
				}

			}
		}
		if !match {
			println(idx + offset)
			return
		}

	}
}
