package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		ps := suffixarray.New(s).Lookup(t, -1)
		sort.Ints(ps)
		c := 0
		for i := range s {
			if j := sort.SearchInts(ps, i); j < len(ps) {
				c += ps[j] + len(t) - 1 - i
			} else {
				c += len(s) - i
			}
		}
		Fprintln(out, c)
	}
}

func main() { run(os.Stdin, os.Stdout) }
