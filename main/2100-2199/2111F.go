package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2111F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, p, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &p, &s)
		r := p / s
		if r > 4 {
			Fprintln(out, -1)
		} else if r == 4 || r == 3 {
			if p%s > 0 {
				Fprintln(out, -1)
			} else {
				k := 5 - r
				Fprintln(out, k)
				for x := range k {
					Fprintln(out, x, 0)
				}
			}
		} else if r == 2 && p%s > 0 {
			if s*2%(p-s*2) > 0 {
				Fprintln(out, -1)
			} else {
				k := s * 2 / (p - s*2)
				Fprintln(out, k)
				for x := range k {
					Fprintln(out, x, 0)
				}
			}
		} else {
			const sz = 50 * 4
			k := (sz*2-1)/p + 1
			base := p*k - 1
			s *= k * 2
			Fprintln(out, s)
			for x := range sz {
				Fprintln(out, x, 0)
			}
			for y := 1; y <= base-sz; y++ {
				Fprintln(out, 0, y)
			}
			need := s - base
			for y := 1; y <= need/(sz-1); y++ {
				for x := 1; x < sz; x++ {
					Fprintln(out, x, y)
				}
			}
			for x := 1; x <= need%(sz-1); x++ {
				Fprintln(out, x, need/(sz-1)+1)
			}
		}
	}
}

//func main() { cf2111F(os.Stdin, os.Stdout) }
