package main

func groupThePeople(groupSizes []int) (res [][]int) {
	for sz := 1; sz <= 505; sz++ {
		ans := []int{}
		for i, v := range groupSizes {
			if v == sz {
				ans = append(ans, i)
				if len(ans) == sz {
					res = append(res, ans)
					ans = []int{}
				}
			}
		}
	}
	return
}
