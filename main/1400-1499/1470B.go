package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1470B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 1e6
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	core := func(x int) int {
		c := 1
		for x > 1 {
			p := lpf[x]
			e := 1
			for x /= p; lpf[x] == p; x /= p {
				e ^= 1
			}
			if e > 0 {
				c *= p
			}
		}
		return c
	}

	var T, n, v, q int
	var w int64
	for Fscan(in, &T); T > 0; T-- {
		cnt := map[int]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			cnt[core(v)]++
		}
		maxC, c1 := 0, cnt[1]
		for v, c := range cnt {
			if c > maxC {
				maxC = c
			}
			if c&1 == 0 && v > 1 {
				c1 += c
			}
		}
		if maxC > c1 {
			c1 = maxC
		}
		for Fscan(in, &q); q > 0; q-- {
			if Fscan(in, &w); w > 0 {
				Fprintln(out, c1)
			} else {
				Fprintln(out, maxC)
			}
		}
	}
}

//func main() { CF1470B(os.Stdin, os.Stdout) }
