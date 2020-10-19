package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(Case int) {
		var n, mx int
		Fscan(in, &n, &mx)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}

		ans := int(1e18)
		for i, v := range a {
			j := sort.Search(n, func(j int) bool { return j > i && 2*(a[j]-v) > mx })
			s1 := sum[j] - sum[i] - (j-i)*v
			s2 := (n-j)*(mx+v) - (sum[n] - sum[j])
			s := s1 + s2

			j = sort.Search(i, func(j int) bool { return 2*(v-a[j]) < mx })
			s1 = sum[j] + j*(mx-v)
			s2 = (i-j)*v - (sum[i] - sum[j])
			s += s1 + s2
			if s < ans {
				ans = s
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
