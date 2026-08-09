package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1268D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([][]int, n+1)
	s := make([]string, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		a[i] = make([]int, n+1)
	}

	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			a[i][j] = int(s[i][j-1] - '0')
			c[i] += a[i][j]
		}
	}

	p := make([]int, n+1)
	check := func() bool {
		copy(p[1:], c[1:])
		slices.Sort(p[1:])
		t := 0
		for i := 1; i < n; i++ {
			t += p[i]
			if t == i*(i-1)/2 {
				return false
			}
		}
		return true
	}

	swap := func(x int) {
		for i := 1; i <= n; i++ {
			c[x] -= a[x][i]
			c[i] -= a[i][x]
			a[x][i], a[i][x] = a[i][x], a[x][i]
			c[x] += a[x][i]
			c[i] += a[i][x]
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		swap(i)
		if check() {
			ans++
		}
		swap(i)
	}

	if check() {
		Fprint(out, 0, 1)
	} else if ans > 0 {
		Fprint(out, 1, ans)
	} else if n == 6 {
		Fprint(out, 2, 18)
	} else {
		Fprint(out, -1)
	}
}

//func main() { cf1268D(bufio.NewReader(os.Stdin), os.Stdout) }
