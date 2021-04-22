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
	var ms, ts [1440][101]int
	var n, m, x, y, cnt, ans int
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &x, &y)
		ms[x][y]++
	}
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		ts[x][y]++
	}

	for i := 1439; i > 0; i-- {
		for j := 100; j >= 0; j-- {
			c := 0
			for k, t := j, ts[i][j]; k <= 100; k++ {
				if ms[i][k] > t {
					c += t
					ms[i][k] -= t
					break
				}
				c += ms[i][k]
				t -= ms[i][k]
				ms[i][k] = 0
			}
			cnt += c
			ans += (i*500 + j*2) * c
		}
		for j, c := range ms[i][:] {
			ms[i-1][j] += c
		}
	}
	Fprint(out, cnt, ans)
}

func main() { run(os.Stdin, os.Stdout) }
