package main

// github.com/EndlessCheng/codeforces-go
func countGoodTriplets(arr []int, a int, b int, c int) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			w := arr[j]
			for k := j + 1; k < len(arr); k++ {
				if x := arr[k]; abs(v-w) <= a && abs(w-x) <= b && abs(v-x) <= c {
					ans++
				}
			}
		}
	}
	return
}
