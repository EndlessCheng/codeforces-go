package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1334A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, a, b int
	for Fscan(in, &t); t > 0; t-- {
		la, lb, ans := 0, 0, true
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &a, &b)
			if b < lb || a-la < b-lb {
				ans = false
			}
			la, lb = a, b
		}
		if ans {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1334A(os.Stdin, os.Stdout) }
