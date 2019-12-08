package main

func minFlips(mat [][]int) int {
	copyMat := func(mat [][]int) [][]int {
		n, m := len(mat), len(mat[0])
		dst := make([][]int, n)
		for i, si := range mat {
			dst[i] = make([]int, m)
			copy(dst[i], si)
		}
		return dst
	}
	hash01Mat := func(mat [][]int) int {
		hash := 0
		cnt := uint(0)
		for _, mi := range mat {
			for _, mij := range mi {
				hash |= mij << cnt
				cnt++
			}
		}
		return hash
	}
	dirOffset4 := [...][2]int{{0, 0}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	cache := map[int]int{0: 0}
	type pair struct {
		mat [][]int
		dep int
	}
	n, m := len(mat), len(mat[0])
	mt0 := make([][]int, n)
	for i := range mt0 {
		mt0[i] = make([]int, m)
	}
	queue := []pair{{mt0, 0}}
	for len(queue) > 0 {
		var p pair
		p, queue = queue[0], queue[1:]
		for i := range mat {
			for j := range mat[i] {
				// for all i,j
				mt2 := copyMat(p.mat)
				for _, dir := range dirOffset4 {
					ii := i + dir[0]
					jj := j + dir[1]
					if ii < 0 || ii >= n || jj < 0 || jj >= m {
						continue
					}
					mt2[ii][jj] ^= 1
				}
				hash := hash01Mat(mt2)
				if _, ok := cache[hash]; ok {
					continue
				}
				cache[hash] = p.dep + 1
				queue = append(queue, pair{mt2, p.dep + 1})
			}
		}
	}
	if dep, ok := cache[hash01Mat(mat)]; ok {
		return dep
	}
	return -1
}
