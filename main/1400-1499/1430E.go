package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1430E(in io.Reader, out io.Writer) {
	var n int
	var s []byte
	Fscan(in, &n, &s)
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

	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}
	ans := int64(0)
	for i := range n {
		b := s[n-1-i] - 'a'
		p := pos[b][0]
		pos[b] = pos[b][1:]
		ans += int64(p + query(p+1, n) - i)
		add(p + 1)
	}
	Fprint(out, ans)
}

//func main() { CF1430E(bufio.NewReader(os.Stdin), os.Stdout) }
