package main

// github.com/EndlessCheng/codeforces-go
func countVowelStrings(n int) (ans int) {
	const mx = 60
	C := [mx + 1][mx + 1]int{}
	for i := 0; i <= mx; i++ {
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	return C[n+4][4]
}
