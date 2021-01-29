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
	var n, mxW, j, sumW int
	Fscan(in, &n, &mxW)
	w := make([]int, n)
	for i := range w {
		Fscan(in, &w[i])
	}
	c := make([]int64, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	var sumC, ans int64
	for i, v := range w {
		sumW += v
		sumC += c[i]
		for sumW > mxW {
			sumW -= w[j]
			sumC -= c[j]
			j++
		}
		if sumC > ans {
			ans = sumC
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
