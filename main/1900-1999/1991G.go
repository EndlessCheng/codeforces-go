package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1991G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k, Q int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k, &Q, &s)
		if n == k && m == k {
			for range Q {
				Fprintln(out, "1 1")
			}
		} else if n == k {
			p := 0
			for _, b := range s {
				if b == 'H' {
					p++
					Fprintln(out, p, 1)
					if p == n {
						p = 0
					}
				} else {
					Fprintln(out, 1, m)
				}
			}
		} else if m == k {
			p := 0
			for _, b := range s {
				if b == 'V' {
					p++
					Fprintln(out, 1, p)
					if p == m {
						p = 0
					}
				} else {
					Fprintln(out, n, 1)
				}
			}
		} else {
			x, y := n, m
			p, q := k, k
			for _, b := range s {
				if b == 'H' {
					Fprintln(out, x, 1)
					x--
					if x == 0 {
						if y <= k {
							q = y
							y = m
							x = k
						} else {
							x = n
						}
					} else if x == k {
						x, p = p, k
					}
				} else {
					Fprintln(out, 1, y)
					y--
					if y == 0 {
						if x <= k {
							p = x
							x = n
							y = k
						} else {
							y = m
						}
					} else if y == k {
						y, q = q, k
					}
				}
			}
		}
	}
}

//func main() { cf1991G(bufio.NewReader(os.Stdin), os.Stdout) }
