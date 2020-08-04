package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF519B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var v int64
	Fscan(in, &n)
	s := make([]int64, 3*n)
	for i := 0; i < 3*n-3; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] + v
	}
	a, b, c := s[n], s[2*n-1], s[3*n-3]
	Fprint(out, 2*a-b, "\n", 2*b-a-c)
}

//func main() {
//	CF519B(os.Stdin, os.Stdout)
//}
