package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func runL(in io.Reader, _w io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e7
	cnt := [mx + 1]int32{}
	for u := 3; u*u < mx*2; u += 2 {
		for v := 1; v < u && (u*u+v*v)/2 <= mx; v += 2 {
			if u*u > v*(v+u*2) && gcd(u, v) == 1 {
				cnt[(u*u+v*v)/2]++
			}
		}
	}
	for i := 2; i <= mx; i++ {
		cnt[i] += cnt[i-1]
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, cnt[n])
	}
}

//func main() { runL(bufio.NewReader(os.Stdin), os.Stdout) }
