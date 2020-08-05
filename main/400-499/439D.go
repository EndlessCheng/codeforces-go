package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol439D(_r io.Reader, _w io.Writer) {
	ternarySearchInt := func(l, r int, f func(x int) int64) int64 {
		for l+3 <= r {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2
			} else {
				l = m1
			}
		}
		min := f(l)
		for i := l + 1; i <= r; i++ {
			if v := f(i); v < min {
				min = v
			}
		}
		return min
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	min, max := a[0], b[m-1]
	if min >= max {
		Fprintln(out, 0)
		return
	}
	Fprint(out, ternarySearchInt(min, max, func(tar int) (cnt int64) {
		for _, v := range a {
			if v >= tar {
				break
			}
			cnt += int64(tar - v)
		}
		for i := m - 1; i >= 0; i-- {
			if b[i] <= tar {
				break
			}
			cnt += int64(b[i] - tar)
		}
		return
	}))
}

//func main() {
//	Sol439D(os.Stdin, os.Stdout)
//}
