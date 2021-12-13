package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1455B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		k := 0
		for ; x > 0; x -= k {
			k++
		}
		if x == -1 {
			k++
		}
		Fprintln(out, k)
	}
}

//func main() { CF1455B(os.Stdin, os.Stdout) }
