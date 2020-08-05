package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF120H(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	genSubStrs := func(s string) []string {
		a := []string{}
		var f func(s, sub string)
		f = func(s, sub string) {
			a = append(a, sub)
			if len(sub) < 4 {
				for i, b := range s {
					f(s[i+1:], sub+string(b))
				}
			}
		}
		f(s, "")
		a = a[1:]
		sort.Strings(a)
		j := 0
		for i := 1; i < len(a); i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}
	var n int
	var s string
	Fscan(in, &n)
	sid := map[string]int{}
	strs := []string{}
	g := make([][]int, n)
	for i := range g {
		Fscan(in, &s)
		a := genSubStrs(s)
		for _, s := range a {
			id, ok := sid[s]
			if !ok {
				id = len(sid)
				sid[s] = id
				strs = append(strs, s)
			}
			g[i] = append(g[i], id)
		}
	}

	matchL := make([]int, n)
	matchR := make([]int, len(sid))
	for i := range matchR {
		matchR[i] = -1
	}
	var used []bool
	var f func(int) bool
	f = func(v int) bool {
		used[v] = true
		for _, w := range g[v] {
			if lv := matchR[w]; lv == -1 || !used[lv] && f(lv) {
				matchR[w] = v
				matchL[v] = w
				return true
			}
		}
		return false
	}
	cnt := 0
	for v := range g {
		used = make([]bool, n)
		if f(v) {
			cnt++
		}
	}
	if cnt < n {
		Fprint(out, -1)
		return
	}
	for _, id := range matchL {
		Fprintln(out, strs[id])
	}
}

//func main() { r, _ := os.Open("input.txt"); w, _ := os.Create("output.txt"); CF120H(r, w) }
