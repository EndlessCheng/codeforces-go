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
	const mx int = 1e5

	var n, v int
	cnt := [mx + 1]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}

	ans := 1
	vis := [mx + 1]bool{}
	for i := 2; i <= mx; i++ {
		if vis[i] {
			continue
		}
		tot := 0
		for j := i; j <= mx; j += i {
			tot += cnt[j]
			vis[i] = true
		}
		if tot > ans {
			ans = tot
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
