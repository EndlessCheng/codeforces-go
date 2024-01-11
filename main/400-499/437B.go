package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf437B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s, lim int
	Fscan(in, &s, &lim)
	ans := []any{}
	for s > 0 && lim > 0 {
		if s >= lim&-lim {
			ans = append(ans, lim)
			s -= lim & -lim
		}
		lim--
	}
	if s > 0 {
		Fprint(out, -1)
	} else {
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf437B(os.Stdin, os.Stdout) }
