package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1082E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, tar, v, ans int
	Fscan(in, &n, &tar)
	sum := make([]int, n+1)
	ps := [5e5 + 1]struct{ minS, sum, preI int }{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i]
		if v == tar {
			sum[i+1]++
			continue
		}
		p := &ps[v]
		cntT := sum[i] - sum[p.preI]
		p.minS = min(p.minS, p.sum-cntT)
		p.sum += 1 - cntT
		ans = max(ans, p.sum-p.minS)
		p.preI = i
	}
	Fprint(out, sum[n]+ans)
}

//func main() { CF1082E(os.Stdin, os.Stdout) }
