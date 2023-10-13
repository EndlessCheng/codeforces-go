package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1765N(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &k)
		j := 0
		for i, b := range s[:k+1] {
			if b > '0' && b < s[j] {
				j = i
			}
		}
		k -= j
		st := s[j : j+1]
		for _, b := range s[j+1:] {
			for len(st) > 1 && k > 0 && b < st[len(st)-1] {
				st = st[:len(st)-1]
				k--
			}
			st = append(st, b)
		}
		st = st[:len(st)-k]
		Fprintf(out, "%s\n", st)
	}
}

//func main() { CF1765N(os.Stdin, os.Stdout) }
