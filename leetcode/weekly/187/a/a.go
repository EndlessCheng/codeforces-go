package main

func destCity(paths [][]string) (ans string) {
	all := map[string]bool{}
	out := map[string]bool{}
	for _, p := range paths {
		v, w := p[0], p[1]
		all[v] = true
		all[w] = true
		out[v] = true
	}
	for s := range all {
		if !out[s] {
			return s
		}
	}
	return
}
