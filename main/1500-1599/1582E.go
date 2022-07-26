package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1582E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		sufMax := make([]int64, n+1)
		for i := n - 1; i >= 0; i-- {
			sufMax[i] = max(sufMax[i+1], a[i])
			a[i] += a[i+1]
		}
		sufMax2 := make([]int64, n)
		k := 2
		for l := n - 1; l >= k; {
			sufMax2[l-k+1] = 0
			for i := l - k; i >= 0; i-- {
				s := a[i] - a[i+k]
				if sufMax2[i+1] < s && s < sufMax[i+k] {
					sufMax2[i] = s
				} else {
					sufMax2[i] = sufMax2[i+1]
				}
			}
			if sufMax2[0] == 0 {
				break
			}
			sufMax, sufMax2 = sufMax2, sufMax
			l -= k
			k++
		}
		Fprintln(out, k-1)
	}
}

//func main() { CF1582E(os.Stdin, os.Stdout) }
