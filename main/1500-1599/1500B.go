package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1500B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, loop int
	var k, cntDiff int64
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	pos := make([]int, max(n, m)*2+1)
	for i := range pos {
		pos[i] = -1
	}
	diff := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = i
		diff[i] = m
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
		if pos[b[i]] >= 0 {
			diff[((pos[b[i]]-i)%n+n)%n]--
		}
	}

	i := 0
	for {
		cntDiff += int64(diff[i])
		loop++
		i = (i + m) % n
		if i == 0 {
			break
		}
	}
	ans := (k - 1) / cntDiff * int64(loop) * int64(m)
	k -= (k - 1) / cntDiff * cntDiff
	for ; k > int64(diff[i]); i = (i + m) % n {
		k -= int64(diff[i])
		ans += int64(m)
	}
	for j := 0; k > 0; j++ {
		if a[i] != b[j] {
			k--
		}
		ans++
		i = (i + 1) % n
	}
	Fprint(out, ans)
}

//func main() { CF1500B(os.Stdin, os.Stdout) }
