package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	sort.Ints(a)
	p := sort.SearchInts(a, 0)
	q := sort.SearchInts(a, 1)

	neg := p * (n - q)
	if k <= neg {
		k = neg - k + 1
		ans := sort.Search(1e18, func(kth int) bool {
			cnt := 0
			j := q
			for _, v := range a[:p] {
				for j < n && a[j]*-v <= kth {
					j++
				}
				cnt += j - q
			}
			return cnt >= k
		})
		Fprint(out, -ans)
		return
	}
	k -= neg

	c0 := q - p
	zero := c0*(n-c0) + c0*(c0-1)/2
	if k <= zero {
		Fprint(out, 0)
		return
	}
	k -= zero

	ans := sort.Search(1e18, func(kth int) bool {
		cnt := 0
		i, j := 0, p-1
		for i < j {
			if a[i]*a[j] > kth {
				i++
			} else {
				cnt += j - i
				j--
			}
		}
		i, j = q, n-1
		for i < j {
			if a[i]*a[j] > kth {
				j--
			} else {
				cnt += j - i
				i++
			}
		}
		return cnt >= k
	})
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
