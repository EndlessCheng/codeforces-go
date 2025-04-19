package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1418E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow := func(x, n int) (res int) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	div := func(a, b int) int { return a * pow(b, mod-2) % mod }

	var n, q, def, dur int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = (sum[i] + v) % mod
	}
	for ; q > 0; q-- {
		Fscan(in, &dur, &def)
		p := sort.SearchInts(a, def)
		cnt := n - p
		if cnt < dur {
			Fprintln(out, 0)
			continue
		}
		ans := div(cnt+1-dur, cnt+1) * sum[p]
		if cnt > dur {
			ans += div(cnt-dur, cnt) * (sum[n] - sum[p] + mod)
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { cf1418E(bufio.NewReader(os.Stdin), os.Stdout) }
