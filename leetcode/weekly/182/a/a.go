package main

func findLucky(arr []int) int {
	cnt := map[int]int{}
	for _, x := range arr {
		cnt[x]++
	}

	ans := -1
	for x, c := range cnt {
		if x == c {
			ans = max(ans, x)
		}
	}
	return ans
}
