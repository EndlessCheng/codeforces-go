package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF978E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, w, min, max, s, v, ans int
	Fscan(in, &n, &w)
	for ; n > 0; n-- {
		Fscan(in, &v)
		s += v
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
	}
	if max-min <= w {
		ans = w - max + min + 1
	}
	Fprint(out, ans)
}

//func main() { CF978E(os.Stdin, os.Stdout) }
