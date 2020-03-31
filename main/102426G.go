package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF102426G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	blocks := [11]int{}
	fill := func(x int) {
		for i := 0; x > 0; x >>= 1 {
			blocks[i] += x & 1
			i++
		}
		for i, v := range blocks {
			if i > 0 {
				Fprint(out, " ")
			}
			Fprint(out, v)
		}
		Fprintln(out)
	}

	var k, x int
	var s []byte
o:
	for Fscan(in, &k); k > 0; k-- {
		if Fscan(in, &s, &x); s[0] == 'f' {
			fill(x)
		} else {
			for i, cnt := range blocks {
				if 1<<i >= x && cnt > 0 {
					blocks[i]--
					fill(1<<i - x)
					continue o
				}
			}
			Fprintln(out, "ERROR!")
		}
	}
}

//func main() { CF102426G(os.Stdin, os.Stdout) }
