package main

// github.com/EndlessCheng/codeforces-go
func getCollisionTimes(cars [][]int) []float64 {
	ans := make([]float64, len(cars))
	st := []int{} // 栈底用时长，栈顶用时短
	for i := len(cars) - 1; i >= 0; i-- {
		p, v := cars[i][0], cars[i][1]
		for len(st) > 0 {
			j := st[len(st)-1]
			w := cars[j][1]
			if v > w {
				t := float64(cars[j][0]-p) / float64(v-w)
				if ans[j] < 0 || t <= ans[j] { // 车 i 与车 j 相遇
					ans[i] = t
					break
				}
			}
			// 车 j 已合并（删除）
			st = st[:len(st)-1]
		}
		if len(st) == 0 { // 车 i 不会与下一辆车相遇
			ans[i] = -1
		}
		st = append(st, i)
	}
	return ans
}
