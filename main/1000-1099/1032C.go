package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1032C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(a, a[n-1]) // 末尾加个哨兵，方便处理只有一个元素 or 最后两个元素相等的情况
	b := make([]int, n)
	for i := 0; i < n; {
		if a[i] == a[i+1] {
			if b[i] == 0 {
				b[i] = 2
				if i > 0 && b[i-1] == 2 {
					b[i] = 3
				}
			}
			i++
			continue
		}
		st := i
		// 处理连续下降段或连续上升段
		for i += 2; i < n && a[i] != a[i-1] && a[i] < a[i-1] == (a[i-1] < a[i-2]); i++ {
		}
		if a[st] > a[st+1] {
			b[st] = 5
			if st > 0 && b[st-1] == 5 {
				b[st] = 4
			}
			for st++; st < i; st++ {
				b[st] = b[st-1] - 1
			}
		} else {
			b[st] = 1
			if st > 0 && b[st-1] == 1 {
				b[st] = 2
			}
			for st++; st < i; st++ {
				b[st] = b[st-1] + 1
			}
		}
		i--
		if b[i] < 1 || b[i] > 5 {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { CF1032C(os.Stdin, os.Stdout) }
