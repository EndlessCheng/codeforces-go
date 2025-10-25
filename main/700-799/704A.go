package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf704A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, tp, v, unr, cur int
	Fscan(in, &n, &q)
	read := make([]bool, q)
	qs := make([][]int, n+1)
	global := []int{}
	for i := range q {
		Fscan(in, &tp, &v)
		if tp == 1 {
			qs[v] = append(qs[v], i)
			global = append(global, i)
			unr++
		} else if tp == 2 {
			for _, j := range qs[v] {
				if !read[j] {
					read[j] = true
					unr--
				}
			}
			qs[v] = nil
		} else if v > cur {
			for _, j := range global[cur:v] {
				if !read[j] {
					read[j] = true
					unr--
				}
			}
			cur = v
		}
		Fprintln(out, unr)
	}
}

//func main() { cf704A(bufio.NewReader(os.Stdin), os.Stdout) }
