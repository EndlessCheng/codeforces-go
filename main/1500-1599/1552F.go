package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1552F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353

	ans := int64(0)
	var n, s, y int
	Fscan(in, &n)
	x := make([]int, n)
	sum := make([]int64, n+1)
	for i := range x {
		Fscan(in, &x[i], &y, &s)
		l := sort.SearchInts(x[:i], y)
		v := (sum[i] - sum[l] + int64(x[i]-y)) % mod
		sum[i+1] = sum[i] + v
		if s > 0 {
			ans += v
		}
	}
	Fprint(out, ((ans+int64(x[n-1]+1))%mod+mod)%mod)
}

//func main() { CF1552F(os.Stdin, os.Stdout) }
