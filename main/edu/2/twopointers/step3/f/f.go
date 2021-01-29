package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, j int
	var s, t []byte
	Fscan(bufio.NewReader(in), &n, &n, &s, &t)
	var c, ct ['z' + 1]int
	for _, b := range t {
		ct[b]++
	}
	ans := int64(0)
	for i, b := range s {
		c[b]++
		for c[b] > ct[b] {
			c[s[j]]--
			j++
		}
		ans += int64(i - j + 1)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
