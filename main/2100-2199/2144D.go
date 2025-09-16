package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2144D(in io.Reader, out io.Writer) {
	const mx = 200001
	var T, n, y, v int
	sum := [mx + 1]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &y)
		cnt := [mx]int{}
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}
		for i := 1; i < mx; i++ {
			sum[i+1] = sum[i] + cnt[i]
		}
		ans := int(-1e18)
		for x := 2; x < mx; x++ {
			s := 0
			for k := 1; (k-1)*x+1 < mx; k++ {
				c := sum[min(k*x+1, mx)] - sum[(k-1)*x+1]
				s += k*c - y*max(c-cnt[k], 0)
			}
			ans = max(ans, s)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2144D(bufio.NewReader(os.Stdin), os.Stdout) }
