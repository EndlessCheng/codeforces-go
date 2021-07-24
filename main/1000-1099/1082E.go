package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1082E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, tar, v, ans int
	Fscan(in, &n, &tar)
	sum := make([]int, n+1)
	ps := [5e5 + 1]struct{ minS, minI, sum, preI int }{}
	for i, ti := 0, 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i]
		if v == tar {
			sum[i+1]++
			ti = i
			continue
		}
		p := &ps[v]
		cntT := sum[i] - sum[p.preI]
		if p.sum-cntT < p.minS {
			p.minS = p.sum - cntT
			p.minI = ti + 1
		}
		p.sum += 1 - cntT
		ans = max(ans, p.sum-p.minS)
		p.preI = i
	}
	Fprint(out, sum[n]+ans)
}

//func main() { CF1082E(os.Stdin, os.Stdout) }
