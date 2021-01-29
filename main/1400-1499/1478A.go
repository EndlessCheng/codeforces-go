package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1478A(_r io.Reader, _w io.Writer) {
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
		mx := 0
		for i := 0; i < n; {
			st := i
			for ; i < n && a[i] == a[st]; i++ {
			}
			if i-st > mx {
				mx = i - st
			}
		}
		Fprintln(out, mx)
	}
}

//func main() { CF1478A(os.Stdin, os.Stdout) }
