package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1694B(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := n
		for i := 1; i < n; i++ {
			if s[i] != s[i-1] {
				ans += i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1694B(bufio.NewReader(os.Stdin), os.Stdout) }
