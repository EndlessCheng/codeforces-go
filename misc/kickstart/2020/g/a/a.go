package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(Case int) {
		var s string
		Fscan(in, &s)
		ans, c := 0, 0
		for {
			i := strings.Index(s, "START")
			if i == -1 {
				break
			}
			for t := s[:i]; ; {
				j := strings.Index(t, "KICK")
				if j == -1 {
					break
				}
				t = t[j+1:]
				c++
			}
			s = s[i+1:]
			ans += c
		}
		Fprintln(out, ans)
	}

	var t int
	Fscan(in, &t)
	for Case := 1; Case <= t; Case++ {
		Fprintf(out, "Case #%d: ", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
