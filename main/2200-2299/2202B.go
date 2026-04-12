package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2202B(in io.Reader, out io.Writer) {
	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if n%2 > 0 {
			s = "b" + s
		}
		for i := 1; i < len(s); i += 2 {
			if s[i] != '?' && s[i] == s[i-1] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2202B(bufio.NewReader(os.Stdin), os.Stdout) }
