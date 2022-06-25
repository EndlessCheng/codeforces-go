package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1283B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if n%k <= k/2 {
			Fprintln(out, n)
		} else {
			Fprintln(out, n-n%k+k/2)
		}
	}
}

//func main() { CF1283B(os.Stdin, os.Stdout) }
