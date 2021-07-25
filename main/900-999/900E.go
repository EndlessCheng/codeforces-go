package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF900E(in io.Reader, out io.Writer) {
	var n, m int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s, &m)
	pos := [2][2][]int{}
	sum := make([]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		if b == '?' {
			sum[i+1]++
		} else {
			pos[b-'a'][i&1] = append(pos[b-'a'][i&1], i)
		}
	}
	dp := make([]struct{ v, op int }, n+1)
	for i := m; i <= n; i++ {
		dp[i] = dp[i-1]
		x := (i ^ m) & 1
		ps := pos[0][x^1]
		if p := sort.SearchInts(ps, i-m); p < len(ps) && ps[p] < i {
			continue
		}
		ps = pos[1][x]
		if p := sort.SearchInts(ps, i-m); p < len(ps) && ps[p] < i {
			continue
		}
		p := dp[i-m]
		p.v++
		p.op += sum[i] - sum[i-m]
		if p.v > dp[i].v || p.v == dp[i].v && p.op < dp[i].op {
			dp[i] = p
		}
	}
	Fprint(out, dp[n].op)
}

//func main() { CF900E(os.Stdin, os.Stdout) }
