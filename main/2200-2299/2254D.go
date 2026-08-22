package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2254D(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ v, i int }
		b := make([]pair, n)
		for i := range b {
			Fscan(in, &b[i].v)
			b[i].i = i
		}
		slices.SortFunc(b, func(a, b pair) int { return a.v - b.v })
		if b[0].v > 0 {
			Fprintln(out, -1)
			continue
		}

		ans := make([]any, n)
		st, pre := 0, 0
		for i, p := range b[:n-1] {
			if p.v == b[i+1].v {
				continue
			}
			x, sz := b[i+1].v-p.v, i-st+1
			if x%sz > 0 || x/sz <= pre {
				Fprintln(out, -1)
				continue o
			}
			x /= sz
			for _, q := range b[st : i+1] {
				ans[q.i] = x
			}
			st = i + 1
			pre = x
		}
		for _, q := range b[st:] {
			ans[q.i] = pre + 1
		}
		Fprintln(out, ans...)
	}
}

//func main() { cf2254D(bufio.NewReader(os.Stdin), os.Stdout) }
