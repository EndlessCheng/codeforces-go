package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1714E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		z, m := -1, 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			if v%5 == 0 {
				v += v % 10
				if z == -1 {
					z = v
				} else if v != z {
					z = -2
				}
			} else {
				m |= 1 << (v % 20)
			}
		}
		if m > 0 {
			if z == -1 && (m&729366 == 0 || m&729366 == m) {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		} else {
			if z != -2 {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		}
	}
}

//func main() { CF1714E(os.Stdin, os.Stdout) }
