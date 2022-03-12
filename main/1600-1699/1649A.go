package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1649A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		land := make([]bool, n)
		for i := range land {
			Fscan(in, &land[i])
		}
		i, j := 1, n-2
		for ; i < n && land[i]; i++ {
		}
		for ; j > i && land[j]; j-- {
		}
		Fprintln(out, j-i+2)
	}
}

//func main() { CF1649A(os.Stdin, os.Stdout) }
