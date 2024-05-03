package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1968E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for i := 1; i < n-1; i++ {
			Fprintln(out, i, i)
		}
		Fprintln(out, n-1, n)
		Fprintln(out, n, n)
	}
}

//func main() { cf1968E(os.Stdin, os.Stdout) }
