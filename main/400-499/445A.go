package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf445A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	var s []byte
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		for j, b := range s {
			if b == '.' {
				if (i+j)%2 > 0 {
					s[j] = 'W'
				} else {
					s[j] = 'B'
				}
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf445A(os.Stdin, os.Stdout) }
