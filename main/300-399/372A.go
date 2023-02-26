package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol372A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	i := 0
	for _, x := range a[(n+1)/2:] {
		if a[i]*2 <= x {
			i++
		}
	}
	Fprint(out, n-i)
}

//func main() { Sol372A(os.Stdin, os.Stdout) }
