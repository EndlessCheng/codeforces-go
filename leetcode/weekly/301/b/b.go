package main

// https://space.bilibili.com/206214/dynamic
type SmallestInfiniteSet map[int]bool

func Constructor() SmallestInfiniteSet { return SmallestInfiniteSet{} }

func (s SmallestInfiniteSet) PopSmallest() int {
	for i := 1; ; i++ {
		if !s[i] {
			s[i] = true
			return i
		}
	}
}

func (s SmallestInfiniteSet) AddBack(n int) {
	delete(s, n)
}
