package main

// 事件扫描线

// github.com/EndlessCheng/codeforces-go
func splitPainting(segments [][]int) (ans [][]int64) {
	// 按位置顺序，记录每个线段左右端点相对应的线段颜色
	events := make([][]int, 1e5+1)
	for _, s := range segments {
		l, r, c := s[0], s[1], s[2]
		events[l] = append(events[l], c)
		events[r] = append(events[r], -c)
	}

	// 按位置顺序扫描每个事件
	sum := 0
	for i, e := range events {
		if e == nil {
			continue
		}
		// 若该位置有区间端点，则记录到答案中
		if len(ans) > 0 {
			if sum > 0 {
				ans[len(ans)-1][1] = int64(i)
				ans[len(ans)-1][2] = int64(sum)
			} else {
				ans = ans[:len(ans)-1] // 没有被涂色的部分不出现在结果中
			}
		}
		ans = append(ans, []int64{int64(i), 0, 0})
		for _, c := range e {
			sum += c
		}
	}
	return ans[:len(ans)-1]
}
