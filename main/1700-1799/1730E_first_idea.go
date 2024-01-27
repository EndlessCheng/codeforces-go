package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1730E_SLOW(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx = 1000001
	divisors := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	type pair struct{ l, r int }
	var leftRange, rightRange [mx]pair
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		left := make([]int, n)  // a[left[i]] >= a[i]
		right := make([]int, n) // a[right[i]] > a[i]
		st := []int{-1}
		for i, v := range a {
			for len(st) > 1 && v > a[st[len(st)-1]] {
				right[st[len(st)-1]] = i
				st = st[:len(st)-1]
			}
			left[i] = st[len(st)-1]
			st = append(st, i)
		}
		for _, i := range st[1:] {
			right[i] = n
		}

		for i := range leftRange {
			leftRange[i].l = -2
			leftRange[i].r = -2
			rightRange[i].l = -2
			rightRange[i].r = -2
		}

		type data struct {
			v int
			p pair
		}
		popV := make([][]data, n)
		oldL := make([]pair, n)
		pushed := make([]bool, n)
		suf := st[:0] // 保存的是值
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(suf) > 0 && v < suf[len(suf)-1] {
				top := suf[len(suf)-1]
				popV[i] = append(popV[i], data{top, rightRange[top]})
				rightRange[top].l = -2
				rightRange[top].r = -2
				suf = suf[:len(suf)-1]
			}

			oldL[i] = rightRange[v]

			if rightRange[v].r == -2 {
				rightRange[v].r = i
			}
			rightRange[v].l = i

			if len(suf) == 0 || suf[len(suf)-1] != v {
				pushed[i] = true
				suf = append(suf, v)
			}
		}

		ans := 0
		pre := []int{} // 保存的是值
		for i, v := range a {
			for len(pre) > 0 && v < pre[len(pre)-1] {
				top := pre[len(pre)-1]
				leftRange[top].l = -2
				leftRange[top].r = -2
				pre = pre[:len(pre)-1]
			}

			if leftRange[v].l == -2 {
				leftRange[v].l = i
			}
			leftRange[v].r = i

			if len(pre) == 0 || v != pre[len(pre)-1] {
				pre = append(pre, v)
			}

			// a[l] >= v
			// a[r] > v
			l, r := left[i], right[i]
			_, _ = l, r
			ds := divisors[v]
			p := rightRange[v]
			ans += p.r - p.l + 1
			for _, d := range ds {
				lp := leftRange[d]
				rp := rightRange[d]
				if lp.l < 0 && rp.l < 0 {
					continue
				}
				if rp.l < 0 {
					
				} else if lp.l < 0 {
					
				} else {
					ans += (i-lp.l+1)*(rp.r-i+1) - (i-lp.r)*(rp.l-i)
				}
			}

			//

			if pushed[i] {
				suf = suf[:len(suf)-1]
			}
			rightRange[v] = oldL[i]
			ps := popV[i]
			for i := len(ps) - 1; i >= 0; i-- {
				d := ps[i]
				rightRange[d.v] = d.p
				suf = append(suf, d.v)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1730E(os.Stdin, os.Stdout) }
