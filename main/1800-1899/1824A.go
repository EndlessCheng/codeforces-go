package main

import (
	"bufio"
	. "fmt"
	"io"
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
		pos := make([]int, m)
		a := []int{}
		l, r := 0, 0
		for ; n > 0; n-- {
			Fscan(in, &v)
			if v == -1 {
				l++
			} else if v == -2 {
				r++
			} else {
				a = append(a, v-1)
				pos[v-1] = 1
			}
		}
		sum := make([]int, m+1)
		for i, v := range pos {
			sum[i+1] = sum[i] + v
		}
		ans := 0
		ans = max(ans, min(r, m-sum[m]))
		ans = max(ans, min(l, m-sum[m]))
		for _, p := range a {
			ll := min(l, p-sum[p])
			rr := min(r, (m-p)-sum[m]+sum[p])
			ans = max(ans, ll+rr)
		}
		Fprintln(out, ans+sum[m])
	}
}

//func main() { CF1824A(os.Stdin, os.Stdout) }
