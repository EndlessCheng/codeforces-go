package main

// https://space.bilibili.com/206214
func mergeArrays(a, b [][]int) (ans [][]int) {
	i, n := 0, len(a)
	j, m := 0, len(b)
	for {
		if i == n {
			return append(ans, b[j:]...)
		}
		if j == m {
			return append(ans, a[i:]...)
		}
		if a[i][0] < b[j][0] {
			ans = append(ans, a[i])
			i++
		} else if a[i][0] > b[j][0] {
			ans = append(ans, b[j])
			j++
		} else {
			a[i][1] += b[j][1]
			ans = append(ans, a[i])
			i++
			j++
		}
	}
}
