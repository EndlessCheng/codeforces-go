package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF1822G2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 999
	primes := []int{}
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}

	_ds := [128]int{1}
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int64{}
		for ; n > 0; n-- {
			Fscan(in, &v)
			cnt[v]++
		}
		ans := int64(0)
		for v, c := range cnt {
			ans += c * (c - 1) * (c - 2)
			ds := _ds[:1]
			x := v
			// 计算所有平方因子
			// 至多是 p^2 * q，因为 pq 不讨论，p^2 下面的 if x > 1 又包括了
			// 所以 p <= U^(1/3)
			for _, p := range primes {
				p2 := p * p
				if p2 > x {
					break
				}
				if x%p2 == 0 {
					e := 1
					for x /= p2; x%p2 == 0; x /= p2 {
						e++
					}
					d := ds
					for pp := p; e > 0; e-- {
						for _, d := range d {
							ds = append(ds, d*pp)
						}
						pp *= p
					}
				}
				if x%p == 0 {
					x /= p
				}
			}
			if x > 1 {
				rt := int(math.Sqrt(float64(x)))
				if rt*rt == x {
					for _, d := range ds {
						ds = append(ds, d*rt)
					}
				}
			}
			for _, d := range ds[1:] {
				ans += c * cnt[v/d] * cnt[v/(d*d)]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1822G2(os.Stdin, os.Stdout) }
