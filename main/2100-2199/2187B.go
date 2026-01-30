package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2187B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x&y == 0 {
			Fprintln(out, x, y)
			continue
		}

		w := bits.Len(uint(x & y))
		v := 1 << (w - 1)
		p, q := v-1|x>>w<<w, v|y>>w<<w
		minD := x - p + y - q

		w--
		k := bits.TrailingZeros(^uint(y>>w)) + w
		p2 := x &^ (1 << k)
		q2 := (y>>w + 1) << w
		if d := x - p2 + q2 - y; d < minD {
			minD = d
			p, q = p2, q2
		}

		k = bits.TrailingZeros(^uint(x>>w)) + w
		q2 = y &^ (1 << k)
		p2 = (x>>w + 1) << w
		if p2-x+y-q2 < minD {
			p, q = p2, q2
		}

		Fprintln(out, p, q)
	}
}

//func main() { cf2187B(bufio.NewReader(os.Stdin), os.Stdout) }
