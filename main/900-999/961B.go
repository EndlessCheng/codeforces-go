package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf961B(in io.Reader, out io.Writer) {
	var n, k, maxS1 int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make([]int, n)
	for i := range t {
		Fscan(in, &t[i])
		t[i] ^= 1
	}
	s := [2]int{}
	for i, c := range a {
		s[t[i]] += c
		if i < k-1 {
			continue
		}
		maxS1 = max(maxS1, s[1])
		if t[i-k+1] > 0 {
			s[1] -= a[i-k+1]
		}
	}
	Fprint(out, s[0]+maxS1)
}

//func main() { cf961B(bufio.NewReader(os.Stdin), os.Stdout) }
