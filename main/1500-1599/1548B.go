package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1548B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n int
	var pre, v int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		n--
		st := make([][18]int64, n)
		for i := range st {
			Fscan(in, &v)
			st[i][0] = abs(v - pre)
			pre = v
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				st[i][j] = gcd(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}

		ans := 1
		for r := 1; r <= n; r++ {
			// 或者双指针 https://codeforces.com/contest/1548/submission/124638948
			l := sort.Search(r, func(l int) bool { k := bits.Len(uint(r-l)) - 1; return gcd(st[l][k], st[r-1<<k][k]) > 1 })
			ans = max(ans, r-l+1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1548B(os.Stdin, os.Stdout) }
