package main

// github.com/EndlessCheng/codeforces-go
type LockingTree struct{}

var g [][]int
var who, pa []int

func Constructor(parent []int) (_ LockingTree) {
	n := len(parent)
	g = make([][]int, n)
	for w := 1; w < n; w++ {
		v := parent[w]
		g[v] = append(g[v], w)
	}
	who, pa = make([]int, n), parent
	return
}

func (LockingTree) Lock(num, user int) bool {
	if who[num] > 0 {
		return false
	}
	who[num] = user
	return true
}

func (LockingTree) Unlock(num, user int) bool {
	if who[num] != user {
		return false
	}
	who[num] = 0
	return true
}

// 判断 v 的子孙是否有锁
func hasLock(v int) bool {
	for _, w := range g[v] {
		if who[w] > 0 || hasLock(w) {
			return true
		}
	}
	return false
}

// 解锁 v 的所有子孙
func unlock(v int) {
	for _, w := range g[v] {
		who[w] = 0
		unlock(w)
	}
}

func (LockingTree) Upgrade(num, user int) bool {
	for v := num; v >= 0; v = pa[v] {
		if who[v] > 0 {
			return false
		}
	}
	if !hasLock(num) {
		return false
	}
	who[num] = user
	unlock(num)
	return true
}
