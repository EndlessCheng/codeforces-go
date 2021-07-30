package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1554D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n == 1 {
			Fprintln(out, "a")
		} else {
			Fprintln(out, strings.Repeat("a", n/2)+"b"+strings.Repeat("a", n&^1-n/2-1)+strings.Repeat("c", n&1))
		}
	}
}

//func main() { CF1554D(os.Stdin, os.Stdout) }
