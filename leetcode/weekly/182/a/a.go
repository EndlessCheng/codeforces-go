package main

func findLucky1(arr []int) int {
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

func findLucky(arr []int) int {
	n := len(arr)
	cnt := make([]int, n+1)
	for _, x := range arr {
		if x <= n {
			cnt[x]++
		}
	}

	for i := n; i >= 1; i-- {
		if cnt[i] == i {
			return i
		}
	}
	return -1
}
