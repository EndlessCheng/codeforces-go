package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1517A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2050 > 0 {
			Fprintln(out, -1)
			continue
		}
		s := 0
		for n /= 2050; n > 0; n /= 10 {
			s += int(n % 10)
		}
		Fprintln(out, s)
	}
}

//func main() { CF1517A(os.Stdin, os.Stdout) }
