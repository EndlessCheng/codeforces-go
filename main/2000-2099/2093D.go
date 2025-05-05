package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2093D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, x, y int
	var op string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		for range q {
			Fscan(in, &op, &x)
			x--
			if op[0] == '-' {
				Fscan(in, &y)
				y--
				ans := 0
				for i := range n {
					ans |= (y>>i&1*3 ^ x>>i&1<<1) << (i * 2)
				}
				Fprintln(out, ans+1)
			} else {
				r, c := 0, 0
				for i := range n {
					v := x >> (i * 2) & 3
					r |= (v>>1 ^ v&1) << i
					c |= v & 1 << i
				}
				Fprintln(out, r+1, c+1)
			}
		}
	}
}

//func main() { cf2093D(bufio.NewReader(os.Stdin), os.Stdout) }
