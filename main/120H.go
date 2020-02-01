package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF120H(_r io.Reader, _w io.Writer) {
	var _ss []string
	var _gen func(s, sub string)
	_gen = func(s, sub string) {
		_ss = append(_ss, sub)
		if len(sub) < 4 {
			for i := range s {
				_gen(s[i+1:], sub+string(s[i]))
			}
		}
	}
	genSubStrs := func(s string) []string {
		_ss = []string{}
		_gen(s, "")
		a := _ss[1:]
		n := len(a)
		sort.Strings(a)
		res := make([]string, 1, n)
		res[0] = a[0]
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				res = append(res, a[i])
			}
		}
		return res
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	strID := map[string]int{}
	idToStr := []string{}
	g := make([][]int, n)
	for i := range g {
		var s string
		Fscan(in, &s)
		ss := genSubStrs(s)
		for _, s := range ss {
			id, ok := strID[s]
			if !ok {
				id = len(strID)
				strID[s] = id
				idToStr = append(idToStr, s)
			}
			g[i] = append(g[i], id)
		}
	}

	matchL := make([]int, n)
	matchR := make([]int, len(strID))
	for i := range matchR {
		matchR[i] = -1
	}
	var used []bool
	var f func(v int) bool
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
		Fprintln(out, idToStr[id])
	}
}

//func main() {
//	r, _ := os.Open("input.txt")
//	w, _ := os.Create("output.txt")
//	CF120H(r, w)
//}
