package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1006D(in io.Reader, out io.Writer) {
	var n, ans int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &s, &t)
	if n&1 > 0 && s[n/2] != t[n/2] {
		ans++
	}
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			ans += bits.OnesCount((1<<(t[i]&31) | 1<<(t[n-1-i]&31)) &^ (1<<(s[i]&31) | 1<<(s[n-1-i]&31)))
		} else if s[i] != s[n-1-i] {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1006D(os.Stdin, os.Stdout) }
