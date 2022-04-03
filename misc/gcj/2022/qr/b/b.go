package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(Case int) {
		var v int
		mi := [4]int{}
		for i := 0; i < 4; i++ {
			mi[i] = 1e9
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				Fscan(in, &v)
				mi[j] = min(mi[j], v)
			}
		}
		sum := 0
		for _, v := range mi {
			sum += v
		}
		if sum < 1e6 {
			Fprintln(out, "IMPOSSIBLE")
			return
		}
		s := 0
		for i, v := range mi {
			if s+v < 1e6 {
				s += v
			} else {
				mi[i] = 1e6 - s
				s = 1e6
			}
			Fprint(out, mi[i], " ")
		}
		Fprintln(out)
	}

	var T int
	Fscan(in, &T)
	for Case := 1; Case <= T; Case++ {
		Fprintf(out, "Case #%d: ", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
