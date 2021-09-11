package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, cnt2, ans int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, 0)
		return
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i] + v
	}
	if sum[n]%3 != 0 {
		Fprint(out, 0)
		return
	}

	avg := sum[n] / 3
	for _, s := range sum[1:n] {
		if s == avg*2 {
			cnt2++
		}
	}

	// 把等于 avg 的前缀和当成第一组，然后累加前缀和等于 avg*2 的前缀和的个数
	for _, s := range sum[1:n] {
		if s == avg*2 {
			cnt2--
		}
		if s == avg {
			ans += cnt2
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
