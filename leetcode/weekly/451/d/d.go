package main

// https://space.bilibili.com/206214
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func lexicographicallySmallestString(s string) (ans string) {
	n := len(s)
	canBeEmpty := make([][]bool, n)
	for i := range canBeEmpty {
		canBeEmpty[i] = make([]bool, n)
	}
	for i := n - 2; i >= 0; i-- {
		canBeEmpty[i+1][i] = true
		for j := i + 1; j < n; j++ {
			// 性质 2
			if isConsecutive(s[i], s[j]) && canBeEmpty[i+1][j-1] {
				canBeEmpty[i][j] = true
				continue
			}
			// 性质 3
			for k := i + 1; k < j-1; k++ {
				if canBeEmpty[i][k] && canBeEmpty[k+1][j] {
					canBeEmpty[i][j] = true
					break
				}
			}
		}
	}

	f := make([]string, n+1)
	for i := n - 1; i >= 0; i-- {
		res := string(s[i]) + f[i+1]
		for j := i + 1; j < n; j++ {
			if canBeEmpty[i][j] {
				res = min(res, f[j+1])
			}
		}
		f[i] = res
	}
	return f[0]
}

func abs(x int) int { if x < 0 { return -x }; return x }
