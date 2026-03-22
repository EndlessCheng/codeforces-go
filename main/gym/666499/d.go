package main

import (
	. "fmt"
	"io"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
func cfD(in io.Reader, out io.Writer) {
	var T int
	for Fscan(in, &T); T > 0; T-- {
		type pair struct {
			v  int
			ch byte
		}
		a := [3]pair{{0, 'R'}, {0, 'G'}, {0, 'B'}}
		Fscan(in, &a[0].v, &a[1].v, &a[2].v)
		slices.SortFunc(a[:], func(a, b pair) int { return a.v - b.v })

		if a[0].v < 3 && a[1].v == a[2].v {
			if a[0].v > 0 {
				Fprintf(out, "%c", a[0].ch)
			}
			Fprint(out, strings.Repeat(string(a[1].ch)+string(a[2].ch), a[1].v))
			if a[0].v > 1 {
				Fprintf(out, "%c", a[0].ch)
			}
			Fprintln(out)
			continue
		}

		a[2].v = min(a[2].v, a[0].v+a[1].v+1)
		ans := make([]byte, a[0].v+a[1].v+a[2].v)
		j := 2
		get := func() byte {
			for a[j].v == 0 {
				j--
			}
			a[j].v--
			return a[j].ch
		}
		for i := 0; i < len(ans); i += 2 {
			ans[i] = get()
		}
		for i := 1; i < len(ans); i += 2 {
			ans[i] = get()
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { cfD(bufio.NewReader(os.Stdin), os.Stdout) }
