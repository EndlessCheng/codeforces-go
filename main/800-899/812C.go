package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf812C(in io.Reader, out io.Writer) {
	var n, s, mn int
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	ans := sort.Search(n, func(k int) bool {
		k++
		for i := range b {
			b[i] = a[i] + (i+1)*k
		}
		slices.Sort(b)
		tot := 0
		for _, v := range b[:k] {
			tot += v
		}
		if tot > s {
			return true
		}
		mn = tot
		return false
	})
	Fprintln(out, ans, mn)
}

//func main() { cf812C(bufio.NewReader(os.Stdin), os.Stdout) }
