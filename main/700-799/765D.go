package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf765D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x int
	Fscan(in, &n)
	left := make([]int, n)
	for i := range left {
		left[i] = -1
	}
	right := []int{}
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		x--
		if left[x] < 0 {
			left[x] = len(right)
			left[i] = len(right)
			right = append(right, x)
		} else if right[left[x]] == x {
			left[i] = left[x]
		} else {
			Fprintln(out, -1)
			return
		}
	}
	Fprintln(out, len(right))
	for _, v := range left {
		Fprint(out, v+1, " ")
	}
	Fprintln(out)
	for _, v := range right {
		Fprint(out, v+1, " ")
	}
}

//func main() { cf765D(os.Stdin, os.Stdout) }
