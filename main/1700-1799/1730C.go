package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1730C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		st := s[:0]
		t := []byte{}
		for _, b := range s {
			for len(st) > 0 && b < st[len(st)-1] {
				t = append(t, min(st[len(st)-1]+1, '9'))
				st = st[:len(st)-1]
			}
			st = append(st, b)
		}
		st = append(st, t...)
		slices.Sort(st)
		Fprintf(out, "%s\n", st)
	}
}

//func main() { cf1730C(os.Stdin, os.Stdout) }
