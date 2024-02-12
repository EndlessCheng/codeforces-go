package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf5E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, mx, cnt, mx2, c2 int
	Fscan(in, &n)
	a := make([]int, n, n*2)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		if v > mx {
			mx2, c2 = mx, cnt
			mx, cnt = v, 1
		} else if v == mx {
			cnt++
		} else if v > mx2 {
			mx2, c2 = v, 1
		} else if v == mx2 {
			c2++
		}
	}
	if cnt == 1 {
		cnt += c2
		mx = mx2
	}

	a = append(a, a...)

	ans := cnt * (cnt - 1) / 2
	st := [][]int{}
	for i := n*2 - 1; i >= 0; i-- {
		if len(st) > 0 && st[0][0] >= i+n {
			st[0] = st[0][1:]
			if len(st[0]) == 0 {
				st = st[1:]
			}
		}
		for len(st) > 0 && a[st[len(st)-1][0]] < a[i] {
			if i < n && a[st[len(st)-1][0]] < mx {
				ans += len(st[len(st)-1])
			}
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			if a[st[len(st)-1][0]] == a[i] {
				if i < n && a[i] < mx {
					ans += len(st[len(st)-1])
					if len(st) > 1 {
						ans++
					}
				}
				st[len(st)-1] = append(st[len(st)-1], i)
			} else {
				if i < n && a[i] < mx {
					ans++
				}
				st = append(st, []int{i})
			}
		} else {
			st = append(st, []int{i})
		}
	}
	Fprint(out, ans)
}

//func main() { cf5E(os.Stdin, os.Stdout) }
