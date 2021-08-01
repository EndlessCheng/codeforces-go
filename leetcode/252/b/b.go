package main

/* 贪心

贪心的方案是，每周可以选择一个剩余任务数最多的项目（且与上一周不同）。

考虑任务最多的项目 $i$，如果它超过了其余任务数之和，则最佳方案是从第一周开始，奇数周选项目 $i$，偶数周选其余项目。若它没有超过其余任务数之和，则可以做完所有任务。

*/

// github.com/EndlessCheng/codeforces-go
func numberOfWeeks(milestones []int) int64 {
	sum, max := 0, 0
	for _, v := range milestones {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum < max*2 {
		return int64((sum-max)*2 + 1)
	}
	return int64(sum)
}
