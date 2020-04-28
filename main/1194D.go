package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1194D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		if k%3 == 0 {
			n %= k + 1
			if n == k {
				Fprintln(out, "Alice")
				continue
			}
		}
		if n%3 != 0 {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { CF1194D(os.Stdin, os.Stdout) }
