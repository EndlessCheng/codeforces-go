package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1582D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 > 0 {
			n -= 3
			Fscan(in, &a, &b, &c)
			if a+b != 0 {
				Fprint(out, -c, -c, a+b)
			} else if a+c != 0 {
				Fprint(out, -b, a+c, -b)
			} else {
				Fprint(out, b+c, -a, -a)
			}
		}
		for ; n > 0; n -= 2 {
			Fscan(in, &a, &b)
			Fprint(out, " ", b, -a)
		}
		Fprintln(out)
	}
}

//func main() { cf1582D(os.Stdin, os.Stdout) }
