package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF922D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, cntS int
	Fscan(in, &n)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		s, t := a[i], a[j]
		return int64(strings.Count(s, "s"))*int64(len(t)) > int64(strings.Count(t, "s"))*int64(len(s))
	})
	ans := int64(0)
	for _, s := range a {
		for _, c := range s {
			if c == 's' {
				cntS++
			} else {
				ans += int64(cntS)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF922D(os.Stdin, os.Stdout) }
