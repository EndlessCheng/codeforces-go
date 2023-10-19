package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF864D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := make([]int, n+2)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	skip := make([]bool, n+1)
	cur := 1
	for i, v := range a {
		if cnt[v] == 1 {
			continue
		}
		for cnt[cur] > 0 {
			cur++
		}
		if cur > v && !skip[v] { // 跳过一次（相同数字改后面的）
			skip[v] = true
			continue
		}
		ans++
		cnt[v]--
		a[i] = cur
		cur++
	}
	Fprintln(out, ans)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF864D(os.Stdin, os.Stdout) }
