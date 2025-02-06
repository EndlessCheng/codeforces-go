package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1469A(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		if n%2 > 0 || s[0] == ')' || s[n-1] == '(' {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { cf1469A(bufio.NewReader(os.Stdin), os.Stdout) }
