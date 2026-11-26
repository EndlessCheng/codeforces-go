package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2154C2(in io.Reader, out io.Writer) {
	const mx = 200_002
	pf := [mx][]int{}
	for i := 2; i < mx; i++ {
		if pf[i] == nil {
			for j := i; j < mx; j += i {
				pf[j] = append(pf[j], i)
			}
		}
	}

	cnt := [mx]int{}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		ok := false
		for i := range a {
			Fscan(in, &a[i])
			for _, p := range pf[a[i]] {
				if cnt[p] > 0 {
					ok = true
				}
				cnt[p]++
			}
		}

		b := make([]int, n)
		mn, mn2 := -1, -1
		for i := range b {
			Fscan(in, &b[i])
			if mn < 0 || b[i] < b[mn] {
				mn2 = mn
				mn = i
			} else if mn2 < 0 || b[i] < b[mn2] {
				mn2 = i
			}
		}

		ans := b[mn] + b[mn2]
		if ok {
			ans = 0
		}

		for i, v := range a {
			for _, p := range pf[v] {
				cnt[p]--
			}
			for _, p := range pf[v+1] {
				if cnt[p] > 0 {
					ans = min(ans, b[i])
					break
				}
			}
			for _, p := range pf[v] {
				cnt[p]++
			}
		}

		for i, v := range a {
			if i != mn {
				for _, p := range pf[v] {
					ans = min(ans, (p-a[mn]%p)*b[mn])
				}
			}
		}

		Fprintln(out, ans)

		for _, v := range a {
			for _, p := range pf[v] {
				cnt[p] = 0
			}
		}
	}
}

//func main() { cf2154C2(bufio.NewReader(os.Stdin), os.Stdout) }
