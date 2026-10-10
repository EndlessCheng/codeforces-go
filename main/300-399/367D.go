package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf367D(in io.Reader, out io.Writer) {
	var n, m, d, L, x int
	Fscan(in, &n, &m, &d)
	a := make([]int, n)
	for i := range m {
		Fscan(in, &L)
		for range L {
			Fscan(in, &x)
			a[x-1] = i
		}
	}

	f := make([]int, 1<<m)
	c := make([]int, m)
	ok := make([]bool, 1<<m)
	x = 0
	for i := range n {
		if c[a[i]] == 0 {
			x ^= 1 << a[i]
		}
		c[a[i]]++
		if i >= d-1 {
			ok[x] = true
		}
		if i >= d-1 {
			c[a[i-d+1]]--
			if c[a[i-d+1]] == 0 {
				x ^= 1 << a[i-d+1]
			}
		}
	}

	ans := m
	for i := 1; i < 1<<m; i++ {
		f[i] = f[i&(i-1)] + 1
		for x = i; x != 0; x -= x & -x {
			ok[i] = ok[i] || ok[i^x&-x]
		}
		if !ok[i] {
			ans = min(ans, m-f[i])
		}
	}
	Fprint(out, ans)
}

//func main() { cf367D(bufio.NewReader(os.Stdin), os.Stdout) }
