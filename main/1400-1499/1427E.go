package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1427E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		x, y uint64
		op   byte
	}

	var x uint64
	Fscan(in, &x)
	ans := []pair{{x, x, '^'}} // write 0
	for i := 0; i < 20; i++ {
		ans = append(ans, pair{x << i, x << i, '+'}) // write x * 2^k, so we can write any multi of x
	}
	const mx = 39
	basis := [mx + 1]uint64{} // xor basis
o:
	for i := uint64(1); basis[0] == 0; i++ {
		bs := []uint64{}
		for j, y := mx, x*i; j >= 0; j-- {
			if y>>j&1 > 0 {
				if basis[j] == 0 {
					basis[j] = y
					y = uint64(0)
					for k := i; k > 0; k &= k - 1 {
						v := x << bits.TrailingZeros64(k)
						ans = append(ans, pair{y, v, '+'}) // write from 0 to x*i by splitting i to some 2^k
						y += v
					}
					for _, b := range bs {
						ans = append(ans, pair{y, b, '^'}) // write from x*i to basis[j]
						y ^= b
					}
					continue o
				}
				if basis[j] > 0 {
					y ^= basis[j]
					bs = append(bs, basis[j])
				}
			}
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintf(out, "%d %c %d\n", p.x, p.op, p.y)
	}
}

//func main() { CF1427E(os.Stdin, os.Stdout) }
