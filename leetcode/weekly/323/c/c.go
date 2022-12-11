package main

// https://space.bilibili.com/206214
type Allocator []int

func Constructor(n int) Allocator {
	return make([]int, n)
}

func (a Allocator) Allocate(size, mID int) int {
	cnt := 0
	for i, id := range a {
		if id > 0 {
			cnt = 0
		} else if cnt++; cnt == size {
			for j := i; j > i-size; j-- {
				a[j] = mID
			}
			return i - size + 1
		}
	}
	return -1
}

func (a Allocator) Free(mID int) (ans int) {
	for i, id := range a {
		if id == mID {
			ans++
			a[i] = 0
		}
	}
	return
}
