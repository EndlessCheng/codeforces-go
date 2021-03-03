package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1494A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
O:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
	o:
		for i := 0; i < 8; i++ {
			c := 0
			for _, b := range s {
				c += i>>(b-'A')&1*2 - 1
				if c < 0 {
					continue o
				}
			}
			if c == 0 {
				Fprintln(out, "YES")
				continue O
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1494A(os.Stdin, os.Stdout) }
