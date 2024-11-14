package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1736C2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, l2, r2, x int
	Fscan(in, &n)
	ans := n * (n + 1) / 2
	right := make([]int, n)
	right2 := make([]int, n)
	for i := range right {
		right[i] = n
		right2[i] = n
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i] - i - 1
		for l2 < r2 && -l2 > v {
			right2[l2] = i
			l2++
		}
		for l < i && -l > v {
			ans -= n - i
			right[l] = i
			l++
			r2++
		}
	}

	sumR := make([]int, n+1)
	sumInc := make([]int, n+1)
	for i, v := range right {
		sumR[i+1] = sumR[i] + v
		sumInc[i+1] = sumInc[i] + right2[i] - v
	}

	ge := make([]int, n+1)
	p := 0
	for i := range ge {
		for right[p] < i {
			p++
		}
		ge[i] = p
	}

	Fscan(in, &m)
	for range m {
		Fscan(in, &p, &x)
		p--
		tar := x - p - 1 // 变为 tar
		l := ge[p+1] // 第一个包含 p 的区间
		if x <= a[p] {
			if tar >= -l {
				Fprintln(out, ans) // 无影响
			} else {
				r := -tar
				// 左端点为 l,l+1,...,r-1 的区间，右开端点从 right[i] 缩小为 p
				Fprintln(out, ans-(sumR[r]-sumR[l]-(r-l)*p))
			}
		} else {
			if p == 0 || right[p-1] < p-1 {
				Fprintln(out, ans) // 无影响
			} else {
				ll := max(ge[p], -tar)
				// 左端点为 ll,ll+1,...,l-1 的区间，都受到 a[p] 影响，右端点从 right[i] 扩大为 right2[i]
				Fprintln(out, ans+sumInc[l]-sumInc[ll])
			}
		}
	}
}

//func main() { cf1736C2(bufio.NewReader(os.Stdin), os.Stdout) }
