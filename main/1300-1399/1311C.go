package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1311C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		d := make([]int, n)
		for d[0] = m; m > 0; m-- {
			Fscan(in, &v)
			d[v]--
		}
		cnt, c := [26]int{}, 0
		for i, v := range d {
			c += v
			cnt[s[i]-'a'] += c + 1
		}
		for _, v := range cnt {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1311C(os.Stdin, os.Stdout) }
