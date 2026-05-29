package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	T, s := 0, []byte{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		type pair struct {
			c int
			b byte
		}
		a := [26]pair{}
		for i := range a {
			a[i].b = 'a' + byte(i)
		}
		for _, b := range s {
			a[b-'a'].c++
		}
		slices.SortFunc(a[:], func(a, b pair) int { return b.c - a.c })

		if a[0].c*2-1 > n {
			Fprintln(out, "No")
			continue
		}

		i := 0
		for _, p := range a {
			for range p.c {
				s[i] = p.b
				i += 2
				if i >= n {
					i = 1
				}
			}
		}
		Fprintf(out, "Yes\n%s\n", s)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
