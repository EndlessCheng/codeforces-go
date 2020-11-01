package main

// github.com/EndlessCheng/codeforces-go
func kthSmallestPath(destination []int, k int) (ans string) {
	const mx = 60
	C := [mx + 1][mx + 1]int{}
	for i := 0; i <= mx; i++ {
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	n, m := destination[0], destination[1]
	i, j := 0, 0
	for i < n || j < m {
		if i == n {
			ans += "H"
			j++
		} else if j == m {
			ans += "V"
			i++
		} else {
			c := C[n-i+m-j-1][n-i]
			if k > c {
				k -= c
				ans += "V"
				i++
			} else {
				ans += "H"
				j++
			}
		}
	}
	return
}
