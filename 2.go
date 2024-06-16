package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(pin string, allowedLetters map[rune]int) bool {
	pinLettersMap := make(map[rune]int)
	for _, r := range pin {
		pinLettersMap[r]++
	}

	if len(pinLettersMap) != len(allowedLetters) {
		return false
	}

	for letter := range allowedLetters {
		if allowedLetters[letter] != pinLettersMap[letter] {
			return false
		}
	}
	return true
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t int
	fmt.Fscan(in, &n, &t)

	in.ReadString('\n')
	allowedLettersInput, _ := in.ReadString('\n')
	allowedLetters := strings.Fields(allowedLettersInput)
	allowedLettersMap := make(map[rune]int)
	for _, letter := range allowedLetters {
		allowedLettersMap[rune(letter[0])]++
	}

	for i := 0; i < t; i++ {
		pin, _ := in.ReadString('\n')
		pin = strings.TrimSpace(pin)
		if len(pin) != n {
			fmt.Fprintln(out, "NO")
			continue
		}
		if isValid(pin, allowedLettersMap) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
