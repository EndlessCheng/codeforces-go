package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2075B(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if k > 1 {
			slices.Sort(a)
			s := 0
			for _, v := range a[n-k-1:] {
				s += v
			}
			Fprintln(out, s)
		} else if n == 2 {
			Fprintln(out, a[0]+a[n-1])
		} else {
			Fprintln(out, max(a[0]+a[n-1], slices.Max(a[1:n-1])+max(a[0], a[n-1])))
		}
	}
}

//func main() { cf2075B(bufio.NewReader(os.Stdin), os.Stdout) }
