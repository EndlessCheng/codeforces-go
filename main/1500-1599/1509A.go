package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1509A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		a := [2][]interface{}{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			a[v&1] = append(a[v&1], v)
		}
		Fprintln(out, a[0]...)
		Fprintln(out, a[1]...)
	}
}

//func main() { CF1509A(os.Stdin, os.Stdout) }
