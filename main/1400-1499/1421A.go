package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1421A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		Fprintln(out, a^b)
	}
}

//func main() { CF1421A(os.Stdin, os.Stdout) }
