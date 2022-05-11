package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF557B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n, w int64
	Fscan(in, &n, &w)
	a := make([]int, n*2)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	x2 := min(int64(a[0]*2), int64(a[n]))
	ans2 := min(x2*n*3, w*2)
	Fprint(out, ans2/2)
	if ans2%2 > 0 {
		Fprint(out, ".5")
	}
}

//func main() { CF557B(os.Stdin, os.Stdout) }
