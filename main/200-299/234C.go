package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf234C(in io.Reader, out io.Writer) {
	var n, pre, suf int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] >= 0 {
			pre++
		}
	}
	ans := n
	for i := n - 1; i > 0; i-- {
		v := a[i]
		if v <= 0 {
			suf++
		}
		if v >= 0 {
			pre--
		}
		ans = min(ans, pre+suf)
	}
	Fprint(out, ans)
}

//func main() { r, _ := os.Open("input.txt"); w, _ := os.Create("output.txt"); cf234C(bufio.NewReader(r), w) }
