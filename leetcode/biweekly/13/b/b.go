package main

// github.com/EndlessCheng/codeforces-go
func findSmallestRegion(regions [][]string, region1 string, region2 string) (ans string) {
	g := map[string][]string{}
	in := map[string]bool{}
	for _, rs := range regions {
		g[rs[0]] = rs[1:]
		for _, s := range rs[1:] {
			in[s] = true
		}
	}
	path := func(r string) (ps []string) {
		var f func(string) bool
		f = func(v string) bool {
			ps = append(ps, v)
			if v == r {
				return true
			}
			for _, w := range g[v] {
				if f(w) {
					return true
				}
			}
			ps = ps[:len(ps)-1]
			return false
		}
		for _, rs := range regions {
			for _, s := range rs {
				if !in[s] {
					f(s)
				}
			}
		}
		return
	}
	p1, p2 := path(region1), path(region2)
	for i := 0; i < len(p1) && i < len(p2) && p1[i] == p2[i]; i++ {
		ans = p1[i]
	}
	return
}
