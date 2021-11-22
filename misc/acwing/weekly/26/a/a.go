package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	s := ""
	Fscan(in, &s)
	ans := 1
	for _, b := range s[1:] {
		if b == '1' {
			ans += 10
		} else {
			ans += int(b & 15)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
