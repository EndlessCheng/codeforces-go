package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2039B(in io.Reader, out io.Writer) {
	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		for i := 1; i < len(s); i++ {
			if s[i-1] == s[i] {
				Fprintln(out, s[i-1:i+1])
				continue o
			}
			if i > 1 && s[i-2] != s[i] {
				Fprintln(out, s[i-2:i+1])
				continue o
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { cf2039B(bufio.NewReader(os.Stdin), os.Stdout) }
