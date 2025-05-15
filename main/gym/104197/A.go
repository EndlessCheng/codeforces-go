package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func runA(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		ans := a[0]*a[1] + a[n-2]*a[n-1]
		for i := range n - 2 {
			ans += a[i] * a[i+2]
		}
		Fprintln(out, ans)
	}
}

//func main() { runA(bufio.NewReader(os.Stdin), os.Stdout) }
