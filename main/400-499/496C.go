package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF496C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([][]byte, n)
	for j := 0; j < m; j++ {
		b[0] = append(b[0], a[0][j])
		for i := 1; i < n; i++ {
			b[i] = append(b[i], a[i][j])
			if bytes.Compare(b[i], b[i-1]) < 0 {
				for ; i >= 0; i-- {
					b[i] = b[i][:len(b[i])-1]
				}
				break
			}
		}
	}
	Fprint(out, m-len(b[0]))
}

//func main() { CF496C(os.Stdin, os.Stdout) }
