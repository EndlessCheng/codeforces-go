package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1012A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n*2)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	ans := (a[n-1] - a[0]) * (a[n*2-1] - a[n])
	for i := n; i < n*2-1; i++ {
		ans = min(ans, (a[i]-a[i-n+1])*(a[n*2-1]-a[0]))
	}
	Fprint(out, ans)
}

//func main() { cf1012A(bufio.NewReader(os.Stdin), os.Stdout) }
