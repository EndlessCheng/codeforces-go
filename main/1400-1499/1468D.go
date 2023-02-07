package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1468D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, m, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &a, &b)
		if a > b {
			a, b = n+1-a, n+1-b
		}
		s := make([]int, m)
		for i := range s {
			Fscan(in, &s[i])
		}
		sort.Ints(s)
		ans := sort.Search(min(m, b-a-1), func(m int) bool {
			m++
			for i, s := range s[:m] {
				if m-i+s >= b {
					return true
				}
			}
			return false
		})
		Fprintln(out, ans)
	}
}

//func main() { CF1468D(os.Stdin, os.Stdout) }
