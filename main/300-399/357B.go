package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF357B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	ans := make([]int, n+1)
	a := [3]int{}
	for ; m > 0; m-- {
		has := 0
		for i := range a {
			Fscan(in, &a[i])
			has |= ans[a[i]]
		}
		i := 1
		for _, v := range a {
			if i == has {
				i++
			}
			if ans[v] == 0 {
				ans[v] = i
				i++
			}
		}
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF357B(os.Stdin, os.Stdout) }
