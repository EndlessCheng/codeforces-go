package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF733C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, tar int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	ans := []int{}
	p := 1
	Fscan(in, &m)
o:
	for i := 1; i <= m; i++ {
		Fscan(in, &tar)
		s, mx, st := 0, 0, p
		for ; p <= n && s+a[p] <= tar; p++ {
			s += a[p]
			if a[p] > mx {
				mx = a[p]
			}
		}
		if s < tar {
			Fprint(out, "NO")
			return
		}
		if p-st == 1 {
			continue
		}
		shift := st - i
		for j := st; j < p; j++ {
			if j > st && a[j] == mx && a[j-1] < mx {
				for k := j; k > st; k-- {
					ans = append(ans, -(k - shift))
				}
				for k := j; k < p-1; k++ {
					ans = append(ans, i)
				}
				continue o
			}
			if j < p-1 && a[j] == mx && a[j+1] < mx {
				for k := j; k < p-1; k++ {
					ans = append(ans, j-shift)
				}
				for k := j; k > st; k-- {
					ans = append(ans, -(k - shift))
				}
				continue o
			}
		}
		Fprint(out, "NO")
		return
	}
	if p <= n { // 漏判了这种情况 WA 了一发
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, v := range ans {
		if v < 0 {
			Fprintln(out, -v, "L")
		} else {
			Fprintln(out, v, "R")
		}
	}
}

//func main() { CF733C(os.Stdin, os.Stdout) }
