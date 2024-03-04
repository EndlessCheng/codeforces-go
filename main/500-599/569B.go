package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf569B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := [1e5 + 1]int{}
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	miss := []int{}
	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			miss = append(miss, i)
		}
	}
	j := 0
	for _, v := range a {
		if v > n || cnt[v] > 1 {
			cnt[v]--
			v = miss[j]
			j++
		}
		Fprint(out, v, " ")
	}
}

//func main() { cf569B(os.Stdin, os.Stdout) }
