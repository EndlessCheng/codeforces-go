package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1118D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make(sort.IntSlice, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(a))
	ans := 1 + sort.Search(n, func(d int) bool {
		d++
		m := m
		for i, v := range a {
			v -= i / d
			if v <= 0 {
				break
			}
			m -= v
			if m <= 0 {
				return true
			}
		}
		return false
	})
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1118D2(os.Stdin, os.Stdout) }
