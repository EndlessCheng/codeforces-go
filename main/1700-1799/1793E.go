package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1793E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	f := make([]int, n+1)
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		v := a[i-1]
		if i >= v {
			f[i] = f[i-v] + 1
			ans[f[i]+n-i] = i
		} else {
			ans[1+n-v] = i
		}
		f[i] = max(f[i], f[i-1])
	}
	for i := n - 1; i > 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &x)
		Fprintln(out, ans[x])
	}
}

//func main() { cf1793E(bufio.NewReader(os.Stdin), os.Stdout) }
