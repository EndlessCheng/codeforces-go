package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

func cf1840D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		ans := sort.Search((a[n-1]-a[0]+3)/6, func(m int) bool {
			i := sort.SearchInts(a, a[0]+m*2+1)
			b := a[i:]
			i = sort.SearchInts(b, b[0]+m*2+1)
			return i == len(b) || b[len(b)-1]-b[i] <= m*2
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1840D(bufio.NewReader(os.Stdin), os.Stdout) }
