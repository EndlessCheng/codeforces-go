package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1352A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		ans := []interface{}{}
		for ten := 1; n > 0; n /= 10 {
			if n%10 > 0 {
				ans = append(ans, n%10*ten)
			}
			ten *= 10
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1352A(os.Stdin, os.Stdout) }
