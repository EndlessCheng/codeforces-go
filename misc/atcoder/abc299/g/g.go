package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	left := make([]int, m+1)
	for i := range a {
		Fscan(in, &a[i])
		left[a[i]]++
	}
	ans := []int{}
	inAns := make([]bool, m+1)
	for _, c := range a {
		left[c]--
		if inAns[c] {
			continue
		}
		for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
			top := ans[len(ans)-1]
			ans = ans[:len(ans)-1]
			inAns[top] = false
		}
		ans = append(ans, c)
		inAns[c] = true
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
