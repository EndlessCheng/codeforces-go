package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2013D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ s, c int }
		st := []pair{}
		for ; n > 0; n-- {
			Fscan(in, &s)
			c := 1
			for len(st) > 0 {
				p := st[len(st)-1]
				if (p.s-1)/p.c+1 <= s/c {
					break
				}
				s += p.s
				c += p.c
				st = st[:len(st)-1]
			}
			st = append(st, pair{s, c})
		}
		p := st[len(st)-1]
		Fprintln(out, (p.s-1)/p.c+1-st[0].s/st[0].c)
	}
}

//func main() { cf2013D(bufio.NewReader(os.Stdin), os.Stdout) }
