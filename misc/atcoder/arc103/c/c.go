package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	n := len(s)
	if s[0] == '0' || s[n-1] == '1' {
		Fprint(out, -1)
		return
	}
	for i, j := 0, n-2; i < j; i++ {
		if s[i] != s[j] {
			Fprint(out, -1)
			return
		}
		j--
	}
	Fprintln(out, 1, 2)
	rt := 2
	i := 1
	for ; i < n/2; i++ {
		v := i + 2
		Fprintln(out, rt, v)
		if s[i] == '1' {
			rt = v
		}
	}
	for ; i+2 <= n; i++ {
		Fprintln(out, rt, i+2)
	}
}

func main() { run(os.Stdin, os.Stdout) }
