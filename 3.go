package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValidString(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}

	for i := 0; i < len(s)-1; i++ {
		if i < len(s)-2 {
			if s[i] == s[i+1] {
				continue
			}
			if s[i] != s[i+2] {
				return false
			} else {
				i++
			}
		} else {
			if s[i] != s[i+1] {
				return false
			}
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
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		if isValidString(s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
