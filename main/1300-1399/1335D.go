package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1335D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		for i := 0; i < 9; i++ {
			Fscan(in, &s)
			j := i%3*3 + i/3
			s[j] = '1' + s[j]&15%9
			Fprintf(out, "%s\n", s)
		}
	}
}

//func main() { CF1335D(os.Stdin, os.Stdout) }
