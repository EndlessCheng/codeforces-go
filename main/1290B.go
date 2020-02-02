package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1290B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	var q, l, r int
	Fscan(in, &s, &q)
	pos := [26][]int{}
	for i, b := range s {
		pos[b-'a'] = append(pos[b-'a'], i)
	}

	solve := func() bool {
		Fscan(in, &l, &r)
		l--
		if r-l == 1 {
			return true
		}
		if s[l] != s[r-1] {
			return true
		}
		cnt := 0
		for _, pi := range pos {
			if sort.SearchInts(pi, l) != sort.SearchInts(pi, r) {
				cnt++
			}
		}
		return cnt >= 3
	}
	for ; q > 0; q-- {
		Fprintln(out, map[bool]string{true: "Yes", false: "No"} [solve()])
	}
}

//func main() { CF1290B(os.Stdin, os.Stdout) }
