package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1922E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		x--
		a := []int{}
		cur := 99
		for x > 0 {
			if x%2 > 0 {
				a = append(a, cur)
				cur--
				x /= 2
			} else {
				a = append(a, 0)
				x--
			}
		}
		Fprintln(out, len(a))
		for i := len(a) - 1; i >= 0; i-- {
			Fprint(out, a[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1922E(os.Stdin, os.Stdout) }
