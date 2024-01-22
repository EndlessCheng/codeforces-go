package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf224B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, left, s int
	Fscan(in, &n, &k)
	minS, l, r := int(1e18), -1, -1
	cnt := map[int]int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
		cnt[a[i]]++
		for len(cnt) == k {
			if s < minS {
				minS = s
				l, r = left+1, i+1
			}
			v := a[left]
			s -= v
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
			left++
		}
	}
	Fprint(out, l, r)
}

//func main() { cf224B(os.Stdin, os.Stdout) }
