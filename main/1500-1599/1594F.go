package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1594F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s, n, k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &n, &k)
		if s == k || s < n*2-n%k {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1594F(os.Stdin, os.Stdout) }
