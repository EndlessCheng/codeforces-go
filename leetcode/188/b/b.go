package main

func countTriplets(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] ^ v
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if sum[j]^sum[i] == sum[k+1]^sum[j] {
					ans++
				}
			}
		}
	}
	return
}
