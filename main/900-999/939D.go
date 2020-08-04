package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF939D(_r io.Reader, _w io.Writer) {
	fa := make([]byte, 'z'+1)
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
	same := func(x, y byte) bool { return find(x) == find(y) }
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s1, s2 []byte
	Fscan(in, &n, &s1, &s2)
	ans := [][2]byte{}
	for i, c := range s1 {
		if !same(c, s2[i]) {
			merge(c, s2[i])
			ans = append(ans, [2]byte{c, s2[i]})
		}
	}
	Fprintln(out, len(ans))
	for _, a := range ans {
		Fprintf(out, "%c %c\n", a[0], a[1])
	}
}

//func main() {
//	CF939D(os.Stdin, os.Stdout)
//}
