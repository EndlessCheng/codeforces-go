package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf385B(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	ans := 0
	pre := -1
	for i := range n - 3 {
		if s[i:i+4] == "bear" {
			ans += (i - pre) * (n - 3 - i)
			pre = i
		}
	}
	Fprint(out, ans)
}

//func main() { cf385B(os.Stdin, os.Stdout) }
