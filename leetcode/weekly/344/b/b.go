package main

// https://space.bilibili.com/206214
type FrequencyTracker struct {
	cnt  map[int]int
	freq map[int]int
}

func Constructor() (_ FrequencyTracker) {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f FrequencyTracker) Add(number int) {
	f.freq[f.cnt[number]]-- // 直接减，因为下面询问的不会涉及到 frequency=0
	f.cnt[number]++
	f.freq[f.cnt[number]]++
}

func (f FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] == 0 {
		return // 不删除任何内容
	}
	f.freq[f.cnt[number]]--
	f.cnt[number]--
	f.freq[f.cnt[number]]++
}

func (f FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0
}
