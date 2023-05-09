package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1824A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		has := map[int]bool{}
		l, r := 0, 0
		for ; n > 0; n-- {
			Fscan(in, &v)
			if v == -1 {
				l++
			} else if v == -2 {
				r++
			} else {
				has[v-1] = true
			}
		}

		a := make([]int, 0, len(has))
		for k := range has {
			a = append(a, k)
		}
		sort.Ints(a)

		fixed := len(a)
		maxFree := min(max(l, r), m-fixed)
		for i, p := range a {
			maxFree = max(maxFree, min(l, p-i)+min(r, (m-1-p)-(fixed-i-1)))
		}
		Fprintln(out, maxFree+fixed)
	}
}

//func main() { CF1824A(os.Stdin, os.Stdout) }
