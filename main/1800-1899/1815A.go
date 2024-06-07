package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1815A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			s += (i%2*2 - 1) * v
		}
		if n%2 > 0 || s >= 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1815A(bufio.NewReader(os.Stdin), os.Stdout) }
