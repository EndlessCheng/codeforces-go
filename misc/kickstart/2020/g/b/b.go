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
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	solve := func(Case int) {
		var n, v, ans int
		Fscan(in, &n)
		sum := map[int]int{}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				Fscan(in, &v)
				sum[i-j] += v
				ans = max(ans, sum[i-j])
			}
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
