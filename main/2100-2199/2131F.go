package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2131F(in io.Reader, out io.Writer) {
	var T, n int
	var a, b string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b)
		type tuple struct{ dy, c0, c1 int }
		ts := make([]tuple, n)
		c1 := 0
		for j, c := range b {
			c1 += int(c & 1)
			ts[j] = tuple{j + 1 - c1*2, j + 1 - c1, c1}
		}
		slices.SortFunc(ts, func(a, b tuple) int { return a.dy - b.dy })

		s0 := make([]int, n+1)
		s1 := make([]int, n+1)
		for i, t := range ts {
			s0[i+1] = s0[i] + t.c0
			s1[i+1] = s1[i] + t.c1
		}

		ans := 0
		c1 = 0
		for i, c := range a {
			c1 += int(c & 1)
			dx := c1*2 - i - 1
			j := sort.Search(n, func(j int) bool { return ts[j].dy >= dx })
			ans += (i+1-c1)*j + s0[j]
			ans += c1*(n-j) + s1[n] - s1[j]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2131F(bufio.NewReader(os.Stdin), os.Stdout) }
