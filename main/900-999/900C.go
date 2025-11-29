package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf900C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	pos := []int{}
	mn := n + 1
	for i := range a {
		Fscan(in, &a[i])
		if len(pos) == 0 || a[i] > a[pos[len(pos)-1]] {
			pos = append(pos, i)
		} else {
			mn = min(mn, a[i])
		}
	}
	if len(pos) == n {
		Fprint(out, 1)
		return
	}

	mx := len(pos)
	for k, i := range pos {
		pre := 0
		if k > 0 {
			pre = a[pos[k-1]]
		}
		nxt := n
		if k < len(pos)-1 {
			nxt = pos[k+1]
		}
		cnt := len(pos) - 1
		for _, v := range a[i+1 : nxt] {
			if v > pre {
				pre = v
				cnt++
			}
		}
		if cnt > mx {
			mx, mn = cnt, a[i]
		}
	}
	Fprint(out, mn)
}

//func main() { cf900C(bufio.NewReader(os.Stdin), os.Stdout) }
