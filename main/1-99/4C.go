package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF4C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	cnt := map[string]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		if cnt[s] > 0 {
			Fprintf(out, "%s%d\n", s, cnt[s])
		} else {
			Fprintln(out, "OK")
		}
		cnt[s]++
	}
}

//func main() { CF4C(os.Stdin, os.Stdout) }
