package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF412D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	has := make(map[[2]int]bool, m)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		has[[2]int{v, w}] = true
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = i + 1
		for j := i - 1; j >= 0 && has[[2]int{ans[j], ans[j+1]}]; j-- {
			ans[j], ans[j+1] = ans[j+1], ans[j]
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF412D(os.Stdin, os.Stdout) }
