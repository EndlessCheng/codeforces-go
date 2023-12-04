package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf888C(in io.Reader, out io.Writer) {
	s := ""
	Fscan(bufio.NewReader(in), &s)

	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}

	n := len(s)
	ans := n
	for _, ps := range pos {
		ps = append(append([]int{-1}, ps...), n)
		maxD := 0
		for i := 1; i < len(ps); i++ {
			maxD = max(maxD, ps[i]-ps[i-1])
		}
		ans = min(ans, maxD)
	}
	Fprint(out, ans)
}

//func main() { cf888C(os.Stdin, os.Stdout) }
