package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2094G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, q, op, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &q)
		var n, s, ss, revS int
		var l, r []int
		for range q {
			Fscan(in, &op)
			if op == 1 {
				if len(r) > 0 {
					v = r[len(r)-1]
					r = r[:len(r)-1]
				} else {
					v = l[0]
					l = l[1:]
				}
				l = append(l, v)
				ss += s - v*n
				revS += v*n - s
			} else if op == 2 {
				l, r = r, l
				ss, revS = revS, ss
			} else {
				Fscan(in, &v)
				r = append(r, v)
				n++
				ss += v * n
				s += v
				revS += s
			}
			Fprintln(out, ss)
		}
	}
}

//func main() { cf2094G(bufio.NewReader(os.Stdin), os.Stdout) }
