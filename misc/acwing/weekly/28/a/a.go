package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	ans := 0
	totQ := strings.Count(s, "Q")
	q := 0
	for _, b := range s {
		if b == 'Q' {
			q++
		} else if b == 'A' {
			ans += q * (totQ - q)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
