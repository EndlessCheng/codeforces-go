package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1766D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx uint32 = 1e7
	lpf := [mx]uint32{}
	for i := uint32(2); i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	var n, x, y uint32
	Fscan(in, &n)
	for range n {
		Fscan(in, &x, &y)
		y -= x
		if y == 1 {
			Fprintln(out, -1)
			continue
		}
		mn := y
		for y > 1 {
			p := lpf[y]
			for y /= p; y%p == 0; y /= p {
			}
			mn = min(mn, (p-x%p)%p)
		}
		Fprintln(out, mn)
	}
}

//func main() { cf1766D(bufio.NewReader(os.Stdin), os.Stdout) }
