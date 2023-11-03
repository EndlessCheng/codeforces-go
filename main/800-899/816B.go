package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF816B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, q, l, r int
	Fscan(in, &n, &k, &q)
	var s, d [2e5 + 2]int
	for ; n > 0; n-- {
		Fscan(in, &l, &r)
		d[l]++
		d[r+1]--
	}
	sd := 0
	for i, d := range d[:2e5+1] {
		sd += d
		s[i+1] = s[i]
		if sd >= k {
			s[i+1]++
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, s[r+1]-s[l])
	}
}

//func main() { CF816B(os.Stdin, os.Stdout) }
