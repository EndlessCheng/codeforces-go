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

	var n, q, op, x, y int
	Fscan(in, &n, &q)
	fa := make([]int, n+q+1)
	newBox := make([]int, n+q+1)
	originBox := make([]int, n+q+1)
	for i := range fa {
		fa[i] = i
		newBox[i] = i
		originBox[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ballNum := n
	emptyBox := n + q
	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		if op == 1 {
			Fscan(in, &y)
			fa[newBox[y]] = newBox[x] // 无需 find，因为每个盒子用完就换新的了
			newBox[y] = emptyBox      // 找一个新的编号当作「空盒子」
			originBox[emptyBox] = y   // 为了操作 3 能正确输出答案
			emptyBox--
		} else if op == 2 {
			ballNum++
			fa[ballNum] = newBox[x]
		} else {
			Fprintln(out, originBox[find(x)])
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
