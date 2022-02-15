package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1188C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	ans := int64(0)
	s := make([][]int, k)
	for i := range s {
		s[i] = make([]int, n+1)
	}
	for j := range s[0] {
		s[0][j] = j
	}
	for d := 1; d*(k-1) <= a[n-1]-a[0]; d++ {
		for i := 1; i < k; i++ {
			p := 0
			for j, v := range a {
				for a[p] <= v-d {
					p++
				}
				s[i][j+1] = (s[i][j] + s[i-1][p]) % mod
			}
		}
		ans += int64(s[k-1][n])
	}
	Fprint(out, ans%mod)
}

//func main() { CF1188C(os.Stdin, os.Stdout) }
