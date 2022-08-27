package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF79C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s string
	var n, l, ll, rr int
	Fscan(in, &s, &n)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return len(a[i]) < len(a[j]) })
	for r := 1; r <= len(s); r++ {
		for _, t := range a {
			if strings.HasSuffix(s[l:r], t) {
				l = r - len(t) + 1
				break
			}
		}
		if r-l > rr-ll {
			ll, rr = l, r
		}
	}
	Fprint(out, rr-ll, ll)
}

//func main() { CF79C(os.Stdin, os.Stdout) }
