package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func CF1626D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		for i := 0; i < n; i++ {
			cnt[i+1] += cnt[i]
		}

		ans := int(1e9)
		for i, c := range cnt {
			for k := 1; k < n*2; k <<= 1 {
				j := sort.SearchInts(cnt[:i+1], k+1) - 1
				c1 := cnt[j]
				c2 := c - c1
				c3 := cnt[n] - c
				x1, x2, x3 := 1, 1, 1
				if c1 > 0 {
					x1 = 1<<bits.Len(uint(c1-1)) - c1
				}
				if c2 > 0 {
					x2 = 1<<bits.Len(uint(c2-1)) - c2
				}
				if c3 > 0 {
					x3 = 1<<bits.Len(uint(c3-1)) - c3
				}
				ans = min(ans, x1+x2+x3)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1626D(os.Stdin, os.Stdout) }
