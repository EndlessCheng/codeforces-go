package main

import "github.com/emirpasic/gods/trees/redblacktree"

// https://space.bilibili.com/206214
type fenwickTree [][2]int

// op=1，添加一个 size
// op=-1，移除一个 size
func (t fenwickTree) update(size, op int) {
	for i := len(t) - size; i < len(t); i += i & -i {
		t[i][0] += op
		t[i][1] += op * size
	}
}

// 返回 >= size 的元素个数，元素和
func (t fenwickTree) query(size int) (cnt, sum int) {
	for i := len(t) - size; i > 0; i &= i - 1 {
		cnt += t[i][0]
		sum += t[i][1]
	}
	return
}

func numberOfAlternatingGroups(a []int, queries [][]int) (ans []int) {
	n := len(a)
	set := redblacktree.New[int, struct{}]()
	t := make(fenwickTree, n+1)

	// op=1，添加一个结束位置 i
	// op=-1，移除一个结束位置 i
	update := func(i, op int) {
		prev, ok := set.Floor(i)
		if !ok {
			prev = set.Right()
		}
		pre := prev.Key

		next, ok := set.Ceiling(i)
		if !ok {
			next = set.Left()
		}
		nxt := next.Key

		t.update((nxt-pre+n-1)%n+1, -op) // 移除/添加旧长度
		t.update((i-pre+n)%n, op)
		t.update((nxt-i+n)%n, op) // 添加/移除新长度
	}

	// 添加一个结束位置 i
	add := func(i int) {
		if set.Empty() {
			t.update(n, 1)
		} else {
			update(i, 1)
		}
		set.Put(i, struct{}{})
	}

	// 移除一个结束位置 i
	del := func(i int) {
		set.Remove(i)
		if set.Empty() {
			t.update(n, -1)
		} else {
			update(i, -1)
		}
	}

	for i, c := range a {
		if c == a[(i+1)%n] {
			add(i) // i 是一个结束位置
		}
	}
	for _, q := range queries {
		if q[0] == 1 {
			if set.Empty() {
				ans = append(ans, n) // 每个长为 size 的子数组都符合要求
			} else {
				cnt, sum := t.query(q[1])
				ans = append(ans, sum-cnt*(q[1]-1))
			}
		} else {
			i := q[1]
			if a[i] == q[2] { // 无影响
				continue
			}
			pre, nxt := (i-1+n)%n, (i+1)%n
			// 修改前，先去掉结束位置
			if a[pre] == a[i] {
				del(pre)
			}
			if a[i] == a[nxt] {
				del(i)
			}
			a[i] ^= 1
			// 修改后，添加新的结束位置
			if a[pre] == a[i] {
				add(pre)
			}
			if a[i] == a[nxt] {
				add(i)
			}
		}
	}
	return
}
