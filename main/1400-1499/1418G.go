package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
)

// https://space.bilibili.com/206214
func CF1418G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	base := make([]uint64, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		base[i] = rand.Uint64()
	}

	cnt := make([]int, n)
	rem := make([]int, n)
	h := make([]uint64, n+1)
	cntH := map[uint64]int{0: 1}

	ans := int64(0)
	l := 0
	for i, v := range a {
		cnt[v]++
		for cnt[v] > 3 {
			cnt[a[l]]--
			cntH[h[l]]--
			l++
		}
		h[i+1] = h[i] + uint64((rem[v]+1)%3-rem[v]%3)*base[v]
		rem[v]++
		ans += int64(cntH[h[i+1]])
		cntH[h[i+1]]++
	}
	Fprint(out, ans)
}

//func main() { CF1418G(os.Stdin, os.Stdout) }
