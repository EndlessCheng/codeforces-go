package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1430E(_r io.Reader, out io.Writer) {
	var n int
	var s []byte
	Fscan(bufio.NewReader(_r), &n, &s)
	tree := make([]int, n+1)
	add := func(i int) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) }

	posS := [26][]int{}
	t := make([]byte, n)
	for i, b := range s {
		b -= 'a'
		t[n-1-i] = b
		posS[b] = append(posS[b], i)
	}
	ans := int64(0)
	for i, b := range t {
		p := posS[b][0]
		posS[b] = posS[b][1:]
		ans += int64(p + query(p+1, n) - i)
		add(p + 1)
	}
	Fprint(out, ans)
}

//func main() { CF1430E(os.Stdin, os.Stdout) }
