package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1728D(in io.Reader, out io.Writer) {
	var T int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		for i := range n / 2 {
			if s[i] != s[n-1-i] {
				for j := i; j < n-i; j += 2 {
					if s[j] != s[j+1] {
						Fprintln(out, "Alice")
						continue o
					}
				}
				break
			}
		}
		Fprintln(out, "Draw")
	}
}

//func main() { cf1728D(bufio.NewReader(os.Stdin), os.Stdout) }
