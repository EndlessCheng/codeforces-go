package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1312A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		Fprintln(out, map[bool]string{true: "YES", false: "NO"} [n%m == 0])
	}
}

//func main() { CF1312A(os.Stdin, os.Stdout) }
