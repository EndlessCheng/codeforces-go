package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 我的憨憨写法（模拟）
func CF1483BList(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	type node struct {
		next *node
		v, i int
	}
	type lst struct{ h, t *node }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ls := []lst{}
		for i, cur := 0, (&node{}); i < n; {
			h := &node{v: a[i], i: i + 1}
			t := h
			cur.next = t
			cur = t
			for i++; i < n && gcd(a[i], a[i-1]) > 1; i++ {
				t = &node{v: a[i], i: i + 1}
				cur.next = t
				cur = t
			}
			ls = append(ls, lst{h, t})
		}
		ans := []interface{}{}
		for len(ls) > 1 {
			tmp := ls
			ls = nil
			cd := false
			for i, n := 0, len(tmp); i < n; {
				p := tmp[i]
				for i++; i < n && gcd(tmp[i-1].t.v, tmp[i].h.v) > 1; i++ {
					tmp[i-1].t.next = tmp[i].h
					p.t = tmp[i].t
				}
				ls = append(ls, p)
				if i < n {
					p := tmp[i]
					ans = append(ans, p.h.i)
					if p.h != p.t {
						tmp[i].h = p.h.next
					} else {
						i++
						if i == n {
							cd = true
						}
					}
				}
			}
			if len(ls) == 1 {
				break
			}
			if cd {
				continue
			}
			if p := ls[0]; gcd(ls[len(ls)-1].t.v, p.h.v) == 1 {
				ans = append(ans, p.h.i)
				if p.h != p.t {
					ls[0].h = p.h.next
				} else {
					ls = ls[1:]
				}
			}
		}
		for h, t := ls[0].h, ls[0].t; gcd(h.v, t.v) == 1; h = h.next {
			ans = append(ans, h.i)
			if h == t {
				break
			}
		}
		Fprint(out, len(ans), " ")
		Fprintln(out, ans...)
	}
}

//func main() { CF1483B(os.Stdin, os.Stdout) }
