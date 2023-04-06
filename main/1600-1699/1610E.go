package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1610E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		mx, cnt := 0, 0
		for i, x := range a {
			cnt++
			if i == n-1 || x != a[i+1] {
				for j := i + 1; j < n; cnt++ {
					j += 1 + sort.SearchInts(a[j+1:], a[j]*2-x)
				}
				mx = max(mx, cnt)
				cnt = 0
			}
		}
		Fprintln(out, n-mx)
	}
}

//func main() { CF1610E(os.Stdin, os.Stdout) }
