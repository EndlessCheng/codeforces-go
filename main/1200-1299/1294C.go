package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1294C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for d := 2; d*d < n; d++ {
			if n%d > 0 {
				continue
			}
			a := d
			n /= d
			for d++; d*d < n; d++ {
				if n%d == 0 {
					Fprintln(out, "YES")
					Fprintln(out, a, d, n/d)
					continue o
				}
			}
			break
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1294C(os.Stdin, os.Stdout) }
