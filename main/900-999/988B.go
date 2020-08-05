package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF988B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	n, s, p := 0, "", ""
	g := [101][]string{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		g[len(s)] = append(g[len(s)], s)
	}
	for _, ss := range g {
		if len(ss) == 0 {
			continue
		}
		for _, s := range ss {
			if s != ss[0] || !strings.Contains(s, p) {
				Fprint(out, "NO")
				return
			}
		}
		p = ss[0]
	}
	Fprintln(out, "YES")
	for _, ss := range g {
		for _, s := range ss {
			Fprintln(out, s)
		}
	}
}

//func main() { CF988B(os.Stdin, os.Stdout) }
