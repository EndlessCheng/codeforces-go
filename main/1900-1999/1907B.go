package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, []byte{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		a := [2][]int{}
		del := make([]bool, len(s))
		for i, b := range s {
			k := b >> 5 & 1
			if b&31 != 2 {
				a[k] = append(a[k], i)
			} else if len(a[k]) > 0 {
				del[a[k][len(a[k])-1]] = true
				a[k] = a[k][:len(a[k])-1]
			}
		}
		t := s[:0]
		for i, b := range del {
			if !b && s[i]&31 != 2 {
				t = append(t, s[i])
			}
		}
		Fprintf(out, "%s\n", t)
	}
}

//func main() { cf1907B(os.Stdin, os.Stdout) }
