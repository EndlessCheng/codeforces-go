package main

import (
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func generateSentences(synonyms [][]string, text string) (ans []string) {
	// 也可以暴力求 cc
	all := map[string]bool{}
	g := map[string]map[string]bool{}
	for _, e := range synonyms {
		v, w := e[0], e[1]
		all[v] = true
		all[w] = true
		if g[v] == nil {
			g[v] = map[string]bool{}
		}
		g[v][w] = true
		if g[w] == nil {
			g[w] = map[string]bool{}
		}
		g[w][v] = true
	}
	var c map[string]bool
	vis := map[string]bool{}
	var dfs func(string)
	dfs = func(v string) {
		vis[v] = true
		c[v] = true
		for w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
		return
	}
	cc := []map[string]bool{}
	for s := range all {
		if !vis[s] {
			c = map[string]bool{}
			dfs(s)
			cc = append(cc, c)
		}
	}

	a := strings.Split(text, " ")
	res := []string{}
	var f func(int)
	f = func(p int) {
		if p == len(a) {
			ans = append(ans, strings.Join(res, " "))
			return
		}
		for _, c := range cc {
			if c[a[p]] {
				for s := range c {
					res = append(res, s)
					f(p + 1)
					res = res[:len(res)-1]
				}
				return
			}
		}
		res = append(res, a[p])
		f(p + 1)
		res = res[:len(res)-1]
	}
	f(0)
	sort.Strings(ans)
	return
}
