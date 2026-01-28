package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1837C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		pre := byte('0')
		for i, b := range s {
			if b == '?' {
				s[i] = pre
			} else {
				pre = b
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf1837C(bufio.NewReader(os.Stdin), os.Stdout) }
