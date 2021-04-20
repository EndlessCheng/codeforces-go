package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1092A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := ""
		for i := 0; i < k; i++ {
			c := n / k
			if i < n%k {
				c++
			}
			s += strings.Repeat(string(byte('a'+i)), c)
		}
		Fprintln(out, s)
	}
}

//func main() { CF1092A(os.Stdin, os.Stdout) }
