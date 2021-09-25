package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1473C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		for i := 1; i < k*2-n; i++ {
			Fprint(out, i, " ")
		}
		for i := k; i >= k*2-n; i-- {
			Fprint(out, i, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1473C(os.Stdin, os.Stdout) }
