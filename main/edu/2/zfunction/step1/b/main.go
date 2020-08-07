package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		c, n := 0, len(s)
		for i := range s {
			for j := i + 1; j <= n; j++ {
				t := s[i:j]
				if x, y := strings.HasPrefix(s, t), strings.HasSuffix(s, t); x && !y || !x && y {
					c++
				}
			}
		}
		Fprintln(out, c)
	}
}

func main() { run(os.Stdin, os.Stdout) }
