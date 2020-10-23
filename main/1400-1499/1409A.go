package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1409A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		if a == b {
			Fprintln(out, 0)
		} else {
			Fprintln(out, (abs(a-b)-1)/10+1)
		}
	}
}

//func main() { CF1409A(os.Stdin, os.Stdout) }
