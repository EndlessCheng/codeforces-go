package main

import "sort"

// https://space.bilibili.com/206214
type ExamTracker struct {
	times  []int
	preSum []int64
}

func Constructor() ExamTracker {
	// preSum 为什么加个 0，见题目 303. 区域和检索 - 数组不可变
	return ExamTracker{[]int{}, []int64{0}}
}

func (e *ExamTracker) Record(time, score int) {
	e.times = append(e.times, time)
	e.preSum = append(e.preSum, e.preSum[len(e.preSum)-1]+int64(score))
}

func (e *ExamTracker) TotalScore(startTime, endTime int) int64 {
	left := sort.SearchInts(e.times, startTime)
	right := sort.SearchInts(e.times, endTime+1) // 也可以在 e.times[left:] 中二分
	return e.preSum[right] - e.preSum[left]
}
