思路是用最大堆来模拟，每次将堆顶减半，累加每次减半的值，直到不低于总和的一半。

虽然浮点数可以通过本题，但是本着能不用浮点就不用浮点的想法（毕竟浮点数计算会产生误差），将每个数都乘上一个 $2$ 的幂次（比如 $2^{20}$），因为可以 [证明](https://leetcode.cn/problems/minimum-operations-to-halve-array-sum/solution/onsuan-fa-by-hqztrue-jalf/) 每个数除 $2$ 的次数不会超过 $20$。

这样就可以愉快地用整数 + 堆来模拟了。

```py [sol-Python3]
class Solution:
    def halveArray(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            nums[i] = -(nums[i] << 20)  # 取负号变最大堆
        heapify(nums)

        ans = 0
        half = sum(nums) // 2
        while half < 0:
            half -= nums[0] // 2
            heapreplace(nums, nums[0] // 2)
            ans += 1
        return ans
```

```go [sol-Go]
func halveArray(nums []int) (ans int) {
	half := 0
	for i := range nums {
		nums[i] <<= 20
		half += nums[i]
	}

	h := hp{nums}
	heap.Init(&h)
	for half /= 2; half > 0; ans++ {
		half -= h.IntSlice[0] / 2
		h.IntSlice[0] /= 2
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice } // 继承 sort.IntSlice 的方法
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

- 时间复杂度：$\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。原地修改，只用到常数额外空间。
