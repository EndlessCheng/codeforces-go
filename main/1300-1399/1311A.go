package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1311A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		c := 0
		if a != b {
			c = 1
			if a < b != ((b-a)&1 > 0) {
				c = 2
			}
		}
		Fprintln(out, c)
	}
}

//func main() { CF1311A(os.Stdin, os.Stdout) }
