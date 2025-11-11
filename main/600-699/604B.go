package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf604B(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	m := a[n-1]
	ans := m + sort.Search(m, func(mx int) bool {
		mx += m
		l, r := 0, n-1
		k := k
		for l <= r {
			k--
			if a[l]+a[r] <= mx {
				l++
			}
			r--
		}
		return k >= 0
	})
	Fprint(out, ans)
}

//func main() { cf604B(bufio.NewReader(os.Stdin), os.Stdout) }
