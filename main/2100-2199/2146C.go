package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2146C(in io.Reader, out io.Writer) {
	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := make([]any, n)
		for i := 0; i < n; {
			if s[i] == '1' {
				ans[i] = i + 1
				i++
				continue
			}
			st := i
			for i++; i < n && s[i] == s[i-1]; i++ {
			}
			if i-st == 1 {
				Fprintln(out, "NO")
				continue o
			}
			for j := st; j < i; j++ {
				ans[j] = i - j + st
			}
		}
		Fprintln(out, "YES")
		Fprintln(out, ans...)
	}
}

//func main() { cf2146C(bufio.NewReader(os.Stdin), os.Stdout) }
