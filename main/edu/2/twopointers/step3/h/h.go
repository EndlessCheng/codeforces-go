package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, mxW, wa, wb int
	Fscan(in, &n, &m, &mxW, &wa, &wb)
	a := make(sort.IntSlice, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(a))
	b := make(sort.IntSlice, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Sort(sort.Reverse(b))
	sumB := make([]int64, m+1)
	for i, v := range b {
		sumB[i+1] = sumB[i] + int64(v)
	}

	ans := sumB[min(mxW/wb, m)]
	sum := int64(0)
	for _, v := range a {
		mxW -= wa
		if mxW < 0 {
			break
		}
		sum += int64(v)
		if s := sum + sumB[min(mxW/wb, m)]; s > ans {
			ans = s
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
