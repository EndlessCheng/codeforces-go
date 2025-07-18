package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2065C2(in io.Reader, out io.Writer) {
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		slices.Sort(b)

		pre := int(-1e9)
		for _, v := range a {
			j := sort.SearchInts(b, pre+v)
			if j < m {
				mn := b[j] - v
				if v >= pre {
					mn = min(mn, v)
				}
				v = mn
			} else if v < pre {
				Fprintln(out, "NO")
				continue o
			}
			pre = v
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2065C2(bufio.NewReader(os.Stdin), os.Stdout) }
