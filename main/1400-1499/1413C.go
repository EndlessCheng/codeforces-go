package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1413C(in io.Reader, out io.Writer) {
	var n int
	a := [6]int{}
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a[:])
	Fscan(in, &n)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	slices.Sort(b)

	ans := int(1e9)
	for _, v := range a {
		it := [5]int{}
		for _, mn := range b {
			mn -= v
			if b[0]-a[0] < mn {
				continue
			}
			mx := b[n-1] - a[5]
			for i := range 5 {
				for it[i] < n && b[it[i]]-a[i+1] < mn {
					it[i]++
				}
				if it[i] > 0 {
					mx = max(mx, b[it[i]-1]-a[i])
				}
			}
			ans = min(ans, mx-mn)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1413C(bufio.NewReader(os.Stdin), os.Stdout) }
