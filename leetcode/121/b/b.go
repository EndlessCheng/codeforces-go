package main

import "sort"

type pair struct {
	v int
	s string
}
type TimeMap struct {
	m map[string][]pair
}

/** Initialize your data structure here. */
func Constructor() (t TimeMap) {
	t.m = map[string][]pair{}
	return
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.m[key] = append(t.m[key], pair{timestamp, value})
}

func (t *TimeMap) Get(key string, timestamp int) (ans string) {
	a := t.m[key]
	i := sort.Search(len(a), func(i int) bool { return a[i].v > timestamp })
	if i > 0 {
		return a[i-1].s
	}
	return
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
