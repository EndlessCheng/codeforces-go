package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func floorSqrt(x int64) int64 {
	res := int64(math.Sqrt(float64(x)))
	if res*res > x {
		res--
	}
	return res
}

func CF1857F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q int
	var v, x, y int64
	for Fscan(in, &T); T > 0; T-- {
		cnt := map[int64]int64{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			cnt[v]++
		}
		for Fscan(in, &q); q > 0; q-- {
			Fscan(in, &x, &y)
			d := x*x - y*4
			if d < 0 {
				Fprint(out, "0 ")
				continue
			}
			rt := floorSqrt(d)
			if rt*rt != d || (x+rt)%2 > 0 {
				Fprint(out, "0 ")
				continue
			}
			ai := (x + rt) / 2
			aj := x - ai
			if ai == aj {
				Fprint(out, cnt[ai]*(cnt[ai]-1)/2, " ")
			} else {
				Fprint(out, cnt[ai]*cnt[aj], " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1857F(os.Stdin, os.Stdout) }
