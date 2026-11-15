package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2276D(in io.Reader, out io.Writer) {
	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		deg := make([]int, n+1)
		for range m {
			Fscan(in, &x, &y)
			deg[x]++
			deg[y]++
		}
		if slices.Contains(deg[1:], 0) {
			Fprintln(out, "Brain Melt")
		} else if m == n-1 && slices.Max(deg) == n-1 {
			Fprintln(out, 2)
		} else {
			Fprintln(out, "Many")
		}
	}
}

//func main() { cf2276D(bufio.NewReader(os.Stdin), os.Stdout) }
