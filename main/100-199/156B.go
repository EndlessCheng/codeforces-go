package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF156B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, cm int
	Fscan(in, &n, &m)
	cnt := make([]int, n+1)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < 0 {
			cnt[-a[i]]--
			m--
		} else {
			cnt[a[i]]++
		}
	}
	for _, c := range cnt[1:] {
		if c == m {
			cm++
		}
	}
	for _, v := range a {
		if v > 0 {
			if cnt[v] != m {
				Fprintln(out, "Lie")
			} else if cm == 1 {
				Fprintln(out, "Truth")
			} else {
				Fprintln(out, "Not defined")
			}
		} else {
			if cnt[-v] != m {
				Fprintln(out, "Truth")
			} else if cm == 1 {
				Fprintln(out, "Lie")
			} else {
				Fprintln(out, "Not defined")
			}
		}
	}
}

//func main() { CF156B(os.Stdin, os.Stdout) }
