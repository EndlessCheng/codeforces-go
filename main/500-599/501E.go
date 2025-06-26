package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf501E(in io.Reader, out io.Writer) {
	var n, odd int
	Fscan(in, &n)
	a := make([]int, n)
	tot := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		tot[a[i]]++
	}
	for _, c := range tot {
		odd += c % 2
	}
	if odd > 1 {
		Fprint(out, 0)
		return
	}

	l := 0
	for l < n/2 && a[l] == a[n-1-l] {
		l++
	}
	if l == n/2 {
		Fprint(out, n*(n+1)/2)
		return
	}
	r := n/2 - 1
	for a[r] == a[n-1-r] {
		r--
	}

	clear(tot)
	cnt := make([]int, n+1)
	for i := l; i <= r; i++ {
		cnt[a[i]]++
		tot[a[i]]++
		tot[a[n-1-i]]++
	}
	bal := true
	for i, c := range cnt {
		if c*2 != tot[i] {
			bal = false
			break
		}
	}
	if bal {
		Fprint(out, (l+1)*(n-r+n-1-r-l))
		return
	}

	clear(tot)
	for _, v := range a[l : n-l] {
		tot[v]++
	}

	clear(cnt)
	rr := n - 1 - l
	for ; ; rr-- {
		v := a[rr]
		cnt[v]++
		if cnt[v] > tot[v]/2 {
			break
		}
	}

	clear(cnt)
	ll := l
	for ; ; ll++ {
		v := a[ll]
		cnt[v]++
		if cnt[v] > tot[v]/2 {
			break
		}
	}

	Fprint(out, (l+1)*(n-rr+ll-l))
}

//func main() { cf501E(bufio.NewReader(os.Stdin), os.Stdout) }
