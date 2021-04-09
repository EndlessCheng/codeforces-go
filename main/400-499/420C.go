package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF420C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, low, v, w int
	Fscan(in, &n, &low)
	deg := make([]int, n)
	type edge struct{ v, w int }
	cntE := map[edge]int{}
	for ; n > 0; n-- {
		Fscan(in, &v, &w)
		v--
		w--
		if v > w {
			v, w = w, v
		}
		deg[v]++
		deg[w]++
		cntE[edge{v, w}]++
	}

	ans := int64(0)
	for e, c := range cntE {
		if s := deg[e.v] + deg[e.w]; s >= low && s-c < low {
			ans--
		}
	}
	sort.Ints(deg)
	for i, d := range deg {
		ans += int64(i - sort.SearchInts(deg[:i], low-d))
	}
	Fprint(out, ans)
}

//func main() { CF420C(os.Stdin, os.Stdout) }
