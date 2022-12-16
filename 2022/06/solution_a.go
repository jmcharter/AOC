package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	var offset int = 4
	var idx int = 3
	for range text {
		idx++
		match := false
		for p1 := idx - offset; p1 < idx; p1++ {
			for p2 := p1 + 1; p2 <= idx; p2++ {
				if text[p1] == text[p2] {
					match = true
					break
				}
			}
			if match {
				break
			}
		}
		if !match {
			println(idx)
			return
		}
	}
}
