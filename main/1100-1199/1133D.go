package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1133D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, y, ans, c0 int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type vec struct{ x, y int }
	cnt := map[vec]int{}
	for _, x := range a {
		Fscan(in, &y)
		if x == 0 {
			if y == 0 {
				c0++
			}
			continue
		}
		g := gcd(x, y)
		v := vec{x / g, y / g}
		cnt[v]++
		ans = max(ans, cnt[v])
	}
	Fprint(out, ans+c0)
}

//func main() { cf1133D(bufio.NewReader(os.Stdin), os.Stdout) }
