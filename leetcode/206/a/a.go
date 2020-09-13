package main

// github.com/EndlessCheng/codeforces-go
func numSpecial(a [][]int) (ans int) {
	for i, row := range a {
	o:
		for j, v := range row {
			if v == 0 {
				continue
			}
			for k, w := range row {
				if k != j && w > 0 {
					continue o
				}
			}
			for k, row2 := range a {
				if k != i && row2[j] > 0 {
					continue o
				}
			}
			ans++
		}
	}
	return
}
