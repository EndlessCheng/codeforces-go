package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf724E(in io.Reader, out io.Writer) {
	var n, c, tot, s int
	Fscan(in, &n, &c)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	for i, v := range a {
		Fscan(in, &s)
		a[i] = c*(n-i-1) + s - v
	}

	slices.Sort(a)
	ans := tot
	for i, v := range a {
		tot += v - c*i
		ans = min(ans, tot)
	}
	Fprint(out, ans)
}

//func main() { cf724E(bufio.NewReader(os.Stdin), os.Stdout) }
