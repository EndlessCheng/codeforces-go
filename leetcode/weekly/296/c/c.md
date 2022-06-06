本题 [视频讲解](https://www.bilibili.com/video/BV1iF41157dG) 已出炉，欢迎三连~

---

## 方法一：正序

注意到 $\textit{nums}$ 的元素互不相同，且执行操作总是将一个元素变成一个不在 $\textit{nums}$ 中的元素，因此每次操作之后 $\textit{nums}$ 的元素仍然互不相同。

用哈希表存储每个元素的位置，然后遍历 $\textit{operations}$ 数组模拟即可。

```python [sol1-Python]
class Solution:
    def arrayChange(self, nums: List[int], operations: List[List[int]]) -> List[int]:
        idx = {num: i for i, num in enumerate(nums)}
        for x, y in operations:
            i = idx[x]
            nums[i] = y
            del idx[x]
            idx[y] = i
        return nums
```

```go [sol1-Go]
func arrayChange(nums []int, operations [][]int) []int {
	idx := make(map[int]int, len(nums))
	for i, num := range nums {
		idx[num] = i
	}
	for _, op := range operations {
		x, y := op[0], op[1]
		i := idx[x]
		nums[i] = y
		delete(idx, x)
		idx[y] = i
	}
	return nums
}
```

## 方法二：逆序

如果没有提示中的那些限制，我们有没有更加通用的做法呢？

我们可以倒着遍历 $\textit{operations}$，设 $x=\textit{operations}[i][0]$，$y=\textit{operations}[i][0]$，倒序遍历的同时用一个哈希表 $\textit{mp}$ 将 $x$ 映射到 $\textit{mp}[y]$ 上，如果 $\textit{mp}[y]$ 不存在则直接映射到 $y$ 上。

这种做法的好处在于，对于 $\textit{operations}$ 靠前的那些 $x$ 到 $y$ 的映射，我们是知道 $x$ **最终**要映射到哪个数字的。 

```Python [sol2-Python3]
class Solution:
    def arrayChange(self, nums: List[int], operations: List[List[int]]) -> List[int]:
        mp = {}
        for x, y in reversed(operations):
            mp[x] = mp.get(y, y)
        return [mp.get(num, num) for num in nums]
```

```go [sol2-Go]
func arrayChange(nums []int, operations [][]int) []int {
	mp := map[int]int{}
	for i := len(operations) - 1; i >= 0; i-- {
		p := operations[i]
		x, y := p[0], p[1]
		if mpY, ok := mp[y]; ok {
			y = mpY
		}
		mp[x] = y
	}
	for i, num := range nums {
		if m, ok := mp[num]; ok {
			nums[i] = m
		}
	}
	return nums
}
```

