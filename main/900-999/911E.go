package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF911E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, k, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	s := []int{n + 1}
	cur := 1
	for i := 0; i < n; i++ {
		if i >= k {
			a = append(a, s[len(s)-1]-1)
		}
		s = append(s, a[i])
		for len(s) > 0 && s[len(s)-1] == cur {
			s = s[:len(s)-1]
			cur++
		}
	}
	if len(s) > 0 {
		Fprint(out, -1)
	} else {
		for _, v := range a {
			Fprint(out, v, " ")
		}
	}
}

//func main() { CF911E(os.Stdin, os.Stdout) }
