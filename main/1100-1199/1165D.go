package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go

// 一种更好的写法是：ans=最小因子*最大因子！ 这样枚举因子就行了！

func CF1165D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	divisors := func(n int64, a []int) bool {
		ds := []int{}
		for d := int64(2); d*d <= n; d++ {
			if n%d == 0 {
				if d > 1e6 {
					return false
				}
				ds = append(ds, int(d))
				if d*d < n {
					if n/d > 1e6 {
						return false
					}
					ds = append(ds, int(n/d))
				}
				if len(ds) > len(a) {
					return false
				}
			}
		}
		if len(ds) < len(a) {
			return false
		}
		sort.Ints(ds)
		sort.Ints(a)
		for i, v := range ds {
			if v != a[i] {
				return false
			}
		}
		return true
	}

	var t, n int
o:
	for Fscan(in, &t); t > 0; t-- {
		ps := map[int]int8{}
		minP := int(1e9)
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			for v := a[i]; v > 1; {
				p := lpf[v]
				e := int8(1)
				for v /= p; lpf[v] == p; v /= p {
					e++
				}
				if e > ps[p] {
					ps[p] = e
				}
				if p < minP {
					minP = p
				}
			}
		}
		ans := int64(1)
		for p, e := range ps {
			for ; e > 0; e-- {
				ans *= int64(p)
				if ans > 1e12 {
					Fprintln(out, -1)
					continue o
				}
			}
		}
		for _, v := range a {
			if int64(v) == ans {
				ans *= int64(minP)
				if ans > 1e12 {
					Fprintln(out, -1)
					continue o
				}
				break
			}
		}
		if !divisors(ans, a) {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1165D(os.Stdin, os.Stdout) }
