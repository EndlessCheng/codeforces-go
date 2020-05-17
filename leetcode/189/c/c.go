package main

import "sort"

func peopleIndexes(mat [][]string) (ans []int) {
	isSubset := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				return true
			}
			if j == m {
				return false
			}
			x, y := a[i], b[j]
			if x < y {
				return false
			} else if x > y {
				j++
			} else {
				i++
				j++
			}
		}
	}

	sid := map[string]int{}
	for _, row := range mat {
		for _, s := range row {
			if sid[s] == 0 {
				sid[s] = len(sid) + 1
			}
		}
	}
	ids := make([][]int, len(mat))
	for i := range ids {
		for _, s := range mat[i] {
			ids[i] = append(ids[i], sid[s])
		}
		sort.Ints(ids[i])
	}
o:
	for i, a := range ids {
		for j, b := range ids {
			if j != i && isSubset(a, b) {
				continue o
			}
		}
		ans = append(ans, i)
	}
	return
}
