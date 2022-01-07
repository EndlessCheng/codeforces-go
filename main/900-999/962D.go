package main

import (
	"bufio"
	"container/list"
	. "fmt"
	"io"
)

// 堆的写法见 https://codeforces.com/problemset/submission/962/141866417

// github.com/EndlessCheng/codeforces-go
func CF962D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var v int64
	Fscan(in, &n)
	pos := make(map[int64]*list.Element, n)
	lst := list.New()
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for pos[v] != nil {
			lst.Remove(pos[v])
			delete(pos, v)
			v *= 2
		}
		pos[v] = lst.PushBack(v) // 这样我们可以按照插入的顺序输出
	}
	Fprintln(out, len(pos))
	for o := lst.Front(); o != nil; o = o.Next() {
		Fprint(out, o.Value, " ")
	}
}

//func main() { CF962D(os.Stdin, os.Stdout) }
