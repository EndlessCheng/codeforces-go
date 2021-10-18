package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1389C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		save := 2
		d := [2]byte{}
		for d[0] = 0; d[0] <= '9'; d[0]++ {
			for d[1] = 0; d[1] <= '9'; d[1]++ {
				c, p := 0, 0
				for _, b := range s {
					if b == d[p] {
						c++
						p ^= 1
					}
				}
				if d[0] != d[1] {
					c &^= 1
				}
				if c > save {
					save = c
				}
			}
		}
		Fprintln(out, len(s)-save)
	}
}

//func main() { CF1389C(os.Stdin, os.Stdout) }
