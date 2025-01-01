package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2035D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		type tuple struct{ base, e, pow2 int }
		st := []tuple{}
		s := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			e, pow2 := 0, 1
			for len(st) > 0 && st[len(st)-1].base>>e < v {
				t := st[len(st)-1]
				s = (s + t.base - t.base*t.pow2) % mod
				e += t.e
				pow2 = pow2 * t.pow2 % mod
				st = st[:len(st)-1]
			}
			s = (s + v*pow2 + mod) % mod
			z := bits.TrailingZeros(uint(v))
			st = append(st, tuple{v >> z, e + z, pow2 << z % mod})
			Fprint(out, s, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2035D(bufio.NewReader(os.Stdin), os.Stdout) }
