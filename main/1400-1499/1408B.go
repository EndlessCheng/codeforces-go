package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1408B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		same, d := true, 1
		for i := range a {
			Fscan(in, &a[i])
			if a[i] != a[0] {
				same = false
			}
			if i > 0 && a[i] != a[i-1] {
				d++
			}
		}
		if same {
			Fprintln(out, 1)
		} else if k == 1 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, (d-2)/(k-1)+1)
		}
	}
}

//func main() { CF1408B(os.Stdin, os.Stdout) }
