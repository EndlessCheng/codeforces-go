下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 方法一：哈希表

用哈希表记录每个元素，然后遍历 $\textit{nums}$，看 $\textit{nums}[j]-\textit{diff}$ 和 $\textit{nums}[j]+\textit{diff}$ 是否都在哈希表中。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        s = set(nums)
        return sum(x - diff in s and x + diff in s for x in nums)
```

```go [sol1-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	for _, x := range nums {
		if set[x-diff] && set[x+diff] {
			ans++
		}
	}
	return
}
```

也可以改为遍历 $\textit{nums}[k]$，这样加入哈希表的的同时可以顺带求出算术三元组。

```py [sol12-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        ans, s = 0, set()
        for x in nums:
            if x - diff in s and x - diff * 2 in s:
                ans += 1
            s.add(x)
        return ans
```

```go [sol12-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		if set[x-diff] && set[x-diff*2] {
			ans++
		}
		set[x] = true
	}
	return
}
```

#### 方法二：三指针

由于 $\textit{nums}$ 是严格递增的，遍历 $\textit{nums}[k]$ 时，$i$ 和 $j$ 只增不减，因此可以用三个指针来实现判断逻辑。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。仅需要几个额外的变量。

```py [sol1-Python3]
class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        ans, i, j = 0, 0, 1
        for x in nums:
            while nums[j] + diff < x:
                j += 1
            if nums[j] + diff > x:
                continue
            while nums[i] + diff * 2 < x:
                i += 1
            if nums[i] + diff * 2 == x:
                ans += 1
        return ans
```

```go [sol1-Go]
func arithmeticTriplets(nums []int, diff int) (ans int) {
	i, j := 0, 1
	for _, x := range nums[2:] {
		for nums[j]+diff < x {
			j++
		}
		if nums[j]+diff > x {
			continue
		}
		for nums[i]+diff*2 < x {
			i++
		}
		if nums[i]+diff*2 == x {
			ans++
		}
	}
	return
}
```
