package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1263D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	fa := [26]byte{}
	for i := range fa {
		fa[i] = byte(i)
	}
	var find func(byte) byte
	find = func(x byte) byte {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to byte) { fa[find(from)] = find(to) }

	var n, cnt int
	var s []byte
	shown := [26]bool{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		h := s[0] - 'a'
		shown[h] = true
		for _, b := range s[1:] {
			merge(b-'a', h)
		}
	}
	for i, b := range shown {
		if b && find(byte(i)) == byte(i) {
			cnt++
		}
	}
	Fprint(_w, cnt)
}

//func main() { CF1263D(os.Stdin, os.Stdout) }
