package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1430F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	l := make([]int, n)
	r := make([]int, n)
	a := make([]int, n)
	for i := range n {
		Fscan(in, &l[i], &r[i], &a[i])
	}

	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if i+1 < n && r[i] == l[i+1] {
			v += f[i+1]
		}
		if (r[i]-l[i]+1)*k < v {
			Fprint(out, -1)
			return
		}
		f[i] = max(v-(r[i]-l[i])*k, 0)
	}

	ans := 0
	left := k
	for i, v := range a {
		if left < f[i] {
			ans += left
			left = k
		}
		ans += v
		left = (left - v%k + k) % k
	}
	Fprint(out, ans)
}

//func main() { cf1430F(bufio.NewReader(os.Stdin), os.Stdout) }
