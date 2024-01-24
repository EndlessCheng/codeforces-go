package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf445A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	var s []byte
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		for j, b := range s {
			if b == '.' {
				s[j] = "BW"[(i+j)%2]
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf445A(os.Stdin, os.Stdout) }
