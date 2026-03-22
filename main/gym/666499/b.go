package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cfB(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := make([]any, n)
		b := []int{}
		for i := len(a) - 1; i >= 0; i-- {
			v := a[i]
			j := sort.SearchInts(b, v)
			ans[i] = max(j, len(b)-sort.SearchInts(b, v+1))
			b = slices.Insert(b, j, v)
		}
		Fprintln(out, ans...)
	}
}

//func main() { cfB(bufio.NewReader(os.Stdin), os.Stdout) }
