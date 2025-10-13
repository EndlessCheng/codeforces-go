package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func cf1470B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1_000_001
	core := [mx]int{}
	for i := 1; i < mx; i++ {
		if core[i] == 0 {
			for j := 1; i*j*j < mx; j++ {
				core[i*j*j] = i
			}
		}
	}

	var T, n, v, q, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			cnt[core[v]]++
		}
		maxC, c1 := 0, 0
		for v, c := range cnt {
			maxC = max(maxC, c)
			if v == 1 || c%2 == 0 {
				c1 += c
			}
		}

		Fscan(in, &q)
		for range q {
			Fscan(in, &w)
			if w > 0 {
				Fprintln(out, max(maxC, c1))
			} else {
				Fprintln(out, maxC)
			}
		}
	}
}

//func main() { cf1470B(bufio.NewReader(os.Stdin), os.Stdout) }
