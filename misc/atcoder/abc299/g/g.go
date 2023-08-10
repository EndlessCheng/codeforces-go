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
	st := []int{}
	inSt := make([]bool, m+1)
	for _, c := range a {
		left[c]--
		if inSt[c] {
			continue
		}
		for len(st) > 0 && c < st[len(st)-1] && left[st[len(st)-1]] > 0 {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			inSt[top] = false
		}
		st = append(st, c)
		inSt[c] = true
	}
	for _, v := range st {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
