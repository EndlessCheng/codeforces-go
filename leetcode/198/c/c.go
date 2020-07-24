package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxNumOfSubstrings(s string) (ans []string) {
	pos := [26][]int{}
	for i, b := range s {
		pos[b-'a'] = append(pos[b-'a'], i)
	}
	var g, rg [26][]int
	for i, pi := range pos {
		if pi == nil {
			continue
		}
		for j, pj := range pos {
			if pj == nil || i == j {
				continue
			}
			if l := sort.SearchInts(pj, pi[0]); l < len(pj) && pi[len(pi)-1] > pj[l] {
				g[i] = append(g[i], j)
				rg[j] = append(rg[j], i)
			}
		}
	}

	vs := make([]int, 0, 26)
	vis := [26]bool{}
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
		vs = append(vs, v)
	}
	for i, ps := range pos {
		if ps != nil && !vis[i] {
			dfs(i)
		}
	}

	vis = [26]bool{}
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !vis[w] {
				rdfs(w)
			}
		}
	}
o:
	for i := len(vs) - 1; i >= 0; i-- {
		if v := vs[i]; !vis[v] {
			comp = []int{}
			rdfs(v)
			for _, v := range comp {
				for _, w := range g[v] {
					if !vis[w] {
						continue o
					}
				}
			}
			v0 := comp[0]
			l, r := pos[v0][0], pos[v0][len(pos[v0])-1]
			for _, v := range comp[1:] {
				if p := pos[v][0]; p < l {
					l = p
				}
				if p := pos[v][len(pos[v])-1]; p > r {
					r = p
				}
			}
			ans = append(ans, s[l:r+1])
		}
	}
	return
}
