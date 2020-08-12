package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1091C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	doDivisors := func(n int, do func(d int)) {
		for d := 1; d*d <= n; d++ {
			if n%d == 0 {
				do(d)
				if d*d < n {
					do(n / d)
				}
			}
		}
		return
	}

	var n int
	Fscan(in, &n)
	ans := []int64{}
	doDivisors(n, func(d int) { ans = append(ans, int64(n-d+2)*int64(n/d)/2) })
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() {
//	CF1091C(os.Stdin, os.Stdout)
//}
