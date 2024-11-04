package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1793D(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := range n {
		Fscan(in, &v)
		a[v] = i
	}
	b := make([]int, n+1)
	for i := range n {
		Fscan(in, &v)
		b[v] = i
	}

	i, j := a[1], b[1]
	l1, r1 := i, i
	l2, r2 := j, j
	if i > j {
		i, j = j, i
	}
	// 不含 1 的方案数
	// 最后 +1 是指整个数组的方案数
	ans := i*(i+1)/2 + (j-i-1)*(j-i)/2 + (n-j-1)*(n-j)/2 + 1
	for v = 2; v <= n; v++ {
		// 计算含 1,2,...,v-1 不含 v 的方案数
		i, j = a[v], b[v]
		if !(l1 < i && i < r1 || l2 < j && j < r2) {
			l := -1
			if i < l1 {
				l = i
			}
			if j < l2 {
				l = max(l, j)
			}
			r := n
			if i > r1 {
				r = i
			}
			if j > r2 {
				r = min(r, j)
			}
			ans += max(min(l1, l2)-l, 0) * max(r-max(r1, r2), 0)
		}
		l1 = min(l1, i)
		r1 = max(r1, i)
		l2 = min(l2, j)
		r2 = max(r2, j)
	}
	Fprint(out, ans)
}

//func main() { cf1793D(bufio.NewReader(os.Stdin), os.Stdout) }
