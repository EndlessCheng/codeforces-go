package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1905C(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		st := []int{}
		for i, b := range s {
			for len(st) > 0 && b > s[st[len(st)-1]] {
				st = st[:len(st)-1]
			}
			st = append(st, i)
		}
		m := len(st)
		t := slices.Clone(s)
		for i, idx := range st {
			t[idx] = s[st[m-1-i]]
		}
		if !slices.IsSorted(t) {
			Fprintln(out, -1)
			continue
		}
		for _, i := range st {
			if s[i] == s[st[0]] {
				m--
			}
		}
		Fprintln(out, m)
	}
}

//func main() { cf1905C(bufio.NewReader(os.Stdin), os.Stdout) }
