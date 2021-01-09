package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1467A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	s := "8901234567"

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		n--
		Fprintln(out, "9"+strings.Repeat(s, n/10)+s[:n%10])
	}
}

//func main() { CF1467A(os.Stdin, os.Stdout) }
