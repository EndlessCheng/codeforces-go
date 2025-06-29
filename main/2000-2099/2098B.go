package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2098B(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		m := n - k
		Fprintln(out, a[k+m/2]-a[(m-1)/2]+1)
	}
}

//func main() { cf2098B(bufio.NewReader(os.Stdin), os.Stdout) }
