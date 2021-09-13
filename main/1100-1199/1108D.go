package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1108D(in io.Reader, out io.Writer) {
	t := []byte("RGB")
	var n, ans int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s)
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			ans++
			for _, c := range t {
				if c != s[i-1] && (i == n-1 || c != s[i+1]) {
					s[i] = c
					break
				}
			}
		}
	}
	Fprintf(out, "%d\n%s", ans, s)
}

//func main() { CF1108D(os.Stdin, os.Stdout) }
