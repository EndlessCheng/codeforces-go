package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func isPal(s string) bool {
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func cf1951E(in io.Reader, out io.Writer) {
	pr := func(a ...any) {
		Fprintln(out, "YES")
		Fprintln(out, len(a))
		Fprintln(out, a...)
	}
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		if !isPal(s) {
			pr(s)
			continue
		}
		n := len(s)
		p := -1
		for i := 1; i < n; i++ {
			if s[i] != s[0] {
				p = i
				break
			}
		}
		if p < 0 {
			Fprintln(out, "NO")
		} else if !isPal(s[p+1:]) {
			pr(s[:p+1], s[p+1:])
		} else if p == 1 || p == n/2 {
			Fprintln(out, "NO")
		} else {
			pr(s[:p+2], s[p+2:])
		}
	}
}

//func main() { cf1951E(bufio.NewReader(os.Stdin), os.Stdout) }
