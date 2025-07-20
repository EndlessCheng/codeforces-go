package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2065G(in io.Reader, out io.Writer) {
	const mx int = 2e5 + 1
	pf := [mx][]int{}
	for i := 2; i < mx; i++ {
		if pf[i] == nil {
			for j := i; j < mx; j += i {
				pf[j] = append(pf[j], i)
			}
		}
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		cnt := make([]int, n+1)
		cnt2 := make([]int, n+1)
		singleP := 0
		for range n {
			Fscan(in, &v)
			a := pf[v]
			if len(a) > 2 {
				continue
			}
			p := a[0]
			if len(a) == 2 {
				q := a[1]
				if p*q == v {
					ans += cnt[v] + cnt[p] + cnt[q] + 1
					cnt2[p]++
					cnt2[q]++
				}
			} else if p*p == v {
				ans += cnt[v] + cnt[p] + 1
				cnt2[p]++
			} else if p == v {
				ans += cnt2[p] + singleP - cnt[p]
				singleP++
			}
			cnt[v]++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2065G(bufio.NewReader(os.Stdin), os.Stdout) }
