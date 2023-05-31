package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF190D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, left int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := int64(0)
	cnt := map[int]int{}
	reachK := false
	for _, v := range a {
		cnt[v]++
		for cnt[v] > k || cnt[v] == k && a[left] != v {
			cnt[a[left]]--
			left++
		}
		if cnt[v] == k {
			reachK = true
		}
		if reachK {
			ans += int64(left + 1)
		}
	}
	Fprint(out, ans)
}

//func main() { CF190D(os.Stdin, os.Stdout) }
