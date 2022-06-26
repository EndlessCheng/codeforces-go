package main

// https://space.bilibili.com/206214/dynamic
const mod int = 1e9 + 7

var f = [1e4 + 1][6][6]int{}

func init() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if j != i && gcd(j+1, i+1) == 1 {
				f[2][i][j] = 1
			}
		}
	}
	for i := 2; i < 1e4; i++ {
		for j := 0; j < 6; j++ {
			for last := 0; last < 6; last++ {
				if last != j && gcd(last+1, j+1) == 1 {
					for last2 := 0; last2 < 6; last2++ {
						if last2 != j {
							f[i+1][j][last] = (f[i+1][j][last] + f[i][last][last2]) % mod
						}
					}
				}
			}
		}
	}
}

func distinctSequences(n int) (ans int) {
	if n == 1 {
		return 6
	}
	for _, row := range f[n] {
		for _, v := range row {
			ans = (ans + v) % mod
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
