package main

// github.com/EndlessCheng/codeforces-go
func countQuadruplets(a []int) (ans int) {
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			w := a[j]
			for k := j + 1; k < len(a); k++ {
				s := v + w + a[k]
				for _, v := range a[k+1:] {
					if v == s {
						ans++
					}
				}
			}
		}
	}
	return
}
