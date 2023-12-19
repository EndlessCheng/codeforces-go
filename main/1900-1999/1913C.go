package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1913C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m, op, v int
	cnt := make([]int, 30)
o:
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &op, &v)
		if op == 1 {
			cnt[v]++
			continue
		}
		s := 0
		for i, c := range cnt {
			s += c
			if v>>i&1 > 0 {
				if s == 0 {
					Fprintln(out, "NO")
					continue o
				}
				s--
			}
			s >>= 1
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1913C(os.Stdin, os.Stdout) }
