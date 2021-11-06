package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1603B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, y int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x > y {
			Fprintln(out, x+y)
		} else {
			Fprintln(out, y-y%x/2)
		}
	}
}

//func main() { CF1603B(os.Stdin, os.Stdout) }
