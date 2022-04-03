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
		var n,m int
		Fscan(in, &n,&m)
		for i := 0; i < n; i++ {
			s := strings.Repeat("+-", m) + "+"
			if i == 0 {
				t := []byte(s)
				t[0] = '.'
				t[1] = '.'
				s = string(t)
			}
			Fprintln(out, s)

			s = strings.Repeat("|.", m) + "|"
			if i == 0 {
				t := []byte(s)
				t[0] = '.'
				s = string(t)
			}
			Fprintln(out, s)
		}
		s := strings.Repeat("+-", m) + "+"
		Fprintln(out, s)
	}

	var T int
	Fscan(in, &T)
	for Case := 1; Case <= T; Case++ {
		Fprintf(out, "Case #%d:\n", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
