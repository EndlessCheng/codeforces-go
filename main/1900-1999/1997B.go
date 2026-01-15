package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1997B(in io.Reader, out io.Writer) {
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		ans := 0
		for i := 3; i <= n; i++ {
			if s[i-3:i] == "..." && t[i-3:i] == "x.x" ||
				t[i-3:i] == "..." && s[i-3:i] == "x.x" {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1997B(bufio.NewReader(os.Stdin), os.Stdout) }
