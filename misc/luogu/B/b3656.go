package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func b3656(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, i int32
	var s string
	qs := [1e6 + 1][2][]int32{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s, &i)
		q := qs[i]
		switch s[1] {
		case 'u':
			Fscan(in, &v)
			j := s[6] & 1
			qs[i][j] = append(q[j], v)
		case 'o':
			j := s[5] & 1
			if len(q[j]) > 0 {
				qs[i][j] = q[j][:len(q[j])-1]
			} else if len(q[j^1]) > 0 {
				qs[i][j^1] = q[j^1][1:]
			}
		case 'i':
			Fprintln(out, len(q[0])+len(q[1]))
		default:
			j := s[1] & 1
			if len(q[j]) > 0 {
				Fprintln(out, q[j][len(q[j])-1])
			} else if len(q[j^1]) > 0 {
				Fprintln(out, q[j^1][0])
			}
		}
	}
}

//func main() { b3656(os.Stdin, os.Stdout) }
