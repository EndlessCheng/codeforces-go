package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func lcp92(s, t string) int {
	for i := range len(t) {
		if s[i] != t[i] {
			return i
		}
	}
	return len(t)
}

func cf1792D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]string, n)
		b := make([]string, n+2)
		s := make([]byte, m)
		pos := make([]byte, m)
		for i := range a {
			for j := range s {
				Fscan(in, &s[j])
				s[j]--
				pos[s[j]] = byte(j)
			}
			a[i] = string(s)
			b[i] = string(pos)
		}

		b[n+1] = "a"
		slices.Sort(b)

		for _, s := range a {
			i := sort.SearchStrings(b, s)
			Fprint(out, max(lcp92(s, b[i]), lcp92(s, b[i-1])), " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1792D(bufio.NewReader(os.Stdin), os.Stdout) }
