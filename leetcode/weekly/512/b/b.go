package main

// https://space.bilibili.com/206214
func aggregateTimeSeries(series1, series2 [][]int) (ans [][]int) {
	n, m := len(series1), len(series2)
	i, j := 0, 0
	for i < n && j < m {
		t1, t2 := series1[i][0], series2[j][0]
		sum := series1[i][1] + series2[j][1]
		if t1 < t2 {
			ans = append(ans, []int{t1, sum})
			i++
		} else if t1 > t2 {
			ans = append(ans, []int{t2, sum})
			j++
		} else { // 相等
			ans = append(ans, []int{t1, sum})
			i++
			j++
		}
	}
	ans = append(ans, series1[i:]...)
	ans = append(ans, series2[j:]...)
	return
}
