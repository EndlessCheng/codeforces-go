下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

$\textit{nums}[i]$ 逐位**异或**任意非负整数，相当于把 $\textit{nums}[i]$ 修改为任意非负整数。

#### 提示 2

$\textit{nums}[i]$ 逐位**与**任意非负整数，相当于把 $\textit{nums}[i]$ 的某些比特位的值，由 $1$ 修改为 $0$，但是不能由 $0$ 修改为 $1$。

#### 提示 3

考虑逐位构造出 $\textit{nums}$ 的最大逐位异或和。如果 $\textit{nums}$ 在某个比特位上有奇数个 $1$，那么这个比特位异或和的结果就能是 $1$。根据提示 2，只需要保证 $\textit{nums}$ 在这个比特位上有至少一个 $1$ 即可。

#### 提示 4

通过逐位**或**运算可以求出 $\textit{nums}$ 在哪些比特位上有 $1$，其结果亦是答案。

```Python [sol1-Python3]
class Solution:
    def maximumXOR(self, nums: List[int]) -> int:
        return reduce(or_, nums)
```

```go [sol1-Go]
func maximumXOR(nums []int) (ans int) {
	for _, num := range nums {
		ans |= num
	}
	return
}
```
