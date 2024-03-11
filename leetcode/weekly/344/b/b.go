package main

// https://space.bilibili.com/206214
type FrequencyTracker struct {
	cnt  map[int]int // number 的出现次数
	freq map[int]int // number 的出现次数的出现次数
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f FrequencyTracker) update(number, delta int) {
	f.freq[f.cnt[number]]-- // 去掉一个旧的 cnt[number]
	f.cnt[number] += delta
	f.freq[f.cnt[number]]++ // 添加一个新的 cnt[number]
}

func (f FrequencyTracker) Add(number int) {
	f.update(number, 1)
}

func (f FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] > 0 {
		f.update(number, -1)
	}
}

func (f FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0 // 至少有一个 number 的出现次数恰好为 frequency
}
