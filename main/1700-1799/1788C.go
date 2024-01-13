package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1788C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 == 0 {
			Fprintln(out, "No")
			continue
		}
		Fprintln(out, "Yes")
		for i := 1; i <= n/2+1; i++ {
			Fprintln(out, i, i+n*3/2)
		}
		for i := n/2 + 2; i <= n; i++ {
			Fprintln(out, i, i+n/2)
		}
	}
}

//func main() { cf1788C(os.Stdin, os.Stdout) }
