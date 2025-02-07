package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	s := ""
	Fscan(in, &s)
	n := len(s)
	ans := n
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans = min(ans, max(i, n-i))
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
