package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
	"strings"
)

func cf1870D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, c, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ i, c int }
		st := []pair{}
		for i := 1; i <= n; i++ {
			Fscan(in, &c)
			for len(st) > 0 && c <= st[len(st)-1].c {
				st = st[:len(st)-1]
			}
			st = append(st, pair{i, c})
		}

		Fscan(in, &k)
		preI, preC, preH := 0, 0, int(1e9)
		for _, p := range st {
			d := p.c - preC
			h := min(k/d, preH)
			k -= d * h
			Fprint(out, strings.Repeat(strconv.Itoa(h)+" ", p.i-preI))
			preI, preC, preH = p.i, p.c, h
		}
		Fprintln(out)
	}
}

//func main() { cf1870D(bufio.NewReader(os.Stdin), os.Stdout) }
