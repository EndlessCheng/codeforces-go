package main

// github.com/EndlessCheng/codeforces-go
const mod = 1_000_000_007

func numOfWays(nums []int) int {
	// 0 和 n+1 作为哨兵节点
	n := len(nums)
	prev := make([]int, n+2) // 双向链表的前驱
	next := make([]int, n+2) // 双向链表的后继
	size := make([]int, n+2) // 子树大小（节点个数）
	for i := range prev {
		prev[i] = i - 1
		next[i] = i + 1
	}
	fac, prod := 1, 1 // 分子分母

	// 自底向上，计算每棵子树的大小
	for i := n - 1; i >= 0; i-- {
		fac = fac * (i + 1) % mod

		x := nums[i]
		size[x] = 1

		// x 左右两侧最近的未被删除节点，即为 x 的左右儿子
		left := prev[x]
		if size[left] > 0 {
			// 加上左子树的大小
			size[x] += size[left]
			// 删除 left：在 left 的前一个节点与 x 之间建立双向链接
			pl := prev[left]
			prev[x] = pl
			next[pl] = x
		}

		right := next[x]
		if size[right] > 0 {
			// 加上右子树的大小
			size[x] += size[right]
			// 删除 right：在 right 的下一个节点与 x 之间建立双向链接
			nr := next[right]
			prev[nr] = x
			next[x] = nr
		}

		prod = prod * size[x] % mod
	}

	return (fac*pow(prod, mod-2) - 1 + mod) % mod // +mod 保证结果非负
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
