package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1382B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		i := 0
		for ; i < n-1 && a[i] == 1; i++ {
		}
		if i&1 == 0 {
			Fprintln(out, "First")
		} else {
			Fprintln(out, "Second")
		}
	}
}

//func main() { CF1382B(os.Stdin, os.Stdout) }
