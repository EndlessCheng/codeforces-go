package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var a, b, c int
	Fscan(in, &a, &b, &c)
	dp := [100][100][100]float64{}
	var f func(a, b, c int) float64
	f = func(a, b, c int) (res float64) {
		if a == 100 || b == 100 || c == 100 {
			return
		}
		dv := &dp[a][b][c]
		if *dv > 0 {
			return *dv
		}
		defer func() { *dv = res }()
		e1 := float64(a) * (1 + f(a+1, b, c))
		e2 := float64(b) * (1 + f(a, b+1, c))
		e3 := float64(c) * (1 + f(a, b, c+1))
		return (e1 + e2 + e3) / float64(a+b+c)
	}
	ans := f(a, b, c)
	Fprintf(out, "%.9f", ans)
}

func main() { run(os.Stdin, os.Stdout) }
