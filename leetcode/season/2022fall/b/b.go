package main

// https://space.bilibili.com/206214
func transportationHub(path [][]int) int {
	nodes := map[int]struct{}{}
	outDeg := map[int]int{}
	inDeg := map[int]int{}
	for _, p := range path {
		x, y := p[0], p[1]
		nodes[x] = struct{}{}
		nodes[y] = struct{}{}
		outDeg[x]++
		inDeg[y]++
	}
	for x := range nodes {
		if outDeg[x] == 0 && inDeg[x] == len(nodes)-1 {
			return x
		}
	}
	return -1
}
