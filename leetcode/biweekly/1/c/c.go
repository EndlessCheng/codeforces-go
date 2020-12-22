package main

// github.com/EndlessCheng/codeforces-go
func assignBikes(workers [][]int, bikes [][]int) (ans int) {
	const inf int = 1e9

	n := len(bikes)
	wt := make([][]int, n)
	for i := range wt {
		wt[i] = make([]int, n)
		if i < len(workers) {
			w := workers[i]
			for j, b := range bikes {
				wt[i][j] = -abs(w[0]-b[0]) - abs(w[1]-b[1])
			}
		}
	}

	match := make([]int, n)
	for i := range match {
		match[i] = -1
	}
	la := make([]int, n)
	for i, r := range wt {
		la[i] = r[0]
		for _, w := range r[1:] {
			if w > la[i] {
				la[i] = w
			}
		}
	}
	lb := make([]int, n)
	slack := make([]int, n)
	for i := 0; i < n; i++ {
		for {
			va := make([]bool, n)
			vb := make([]bool, n)
			for j := range slack {
				slack[j] = inf
			}
			var f func(int) bool
			f = func(v int) bool {
				va[v] = true
				for w, b := range vb {
					if !b {
						if d := la[v] + lb[w] - wt[v][w]; d == 0 {
							vb[w] = true
							if match[w] == -1 || f(match[w]) {
								match[w] = v
								return true
							}
						} else if d < slack[w] {
							slack[w] = d
						}
					}
				}
				return false
			}
			if f(i) {
				break
			}
			d := inf
			for j, b := range vb {
				if !b && slack[j] < d {
					d = slack[j]
				}
			}
			for j := 0; j < n; j++ {
				if va[j] {
					la[j] -= d
				}
				if vb[j] {
					lb[j] += d
				}
			}
		}
	}
	for w, v := range match {
		ans -= wt[v][w]
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
