package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF607A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ p, d int }, n)
	for i := range a {
		Fscan(in, &a[i].p, &a[i].d)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].p < a[j].p })
	dp := make([]int, n)
	for i, p := range a {
		j := sort.Search(len(a), func(i int) bool { return a[i].p >= p.p-p.d })
		if j > 0 {
			dp[i] = dp[j-1] + 1
		} else {
			dp[i] = 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	Fprint(out, n-ans)
}

//func main() { CF607A(os.Stdin, os.Stdout) }
