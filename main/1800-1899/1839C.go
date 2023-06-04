package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1839C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			Fscan(in, &a[i])
		}
		if a[0] > 0 {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		for i := 0; i < n; {
			st := i
			for ; i < n && a[i] == 0; i++ {}
			st2 := i
			for ; i < n && a[i] == 1; i++ {}
			Fprint(out, strings.Repeat("0 ", i-st-1), i-st2, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1839C(os.Stdin, os.Stdout) }
