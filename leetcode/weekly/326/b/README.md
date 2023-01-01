遍历每个数，分解质因数，把质因数加到哈希表中。分解质因数的方法见我在第 324 场周赛中讲过了，可以看看。

答案为哈希表的大小。

```py [sol1-Python3]
class Solution:
    def distinctPrimeFactors(self, nums: List[int]) -> int:
        s = set()
        for x in nums:
            i = 2
            while i * i <= x:
                if x % i == 0:
                    s.add(i)
                    x //= i
                    while x % i == 0:
                        x //= i
                i += 1
            if x > 1:
                s.add(x)
        return len(s)
```

```go [sol1-Go]
func distinctPrimeFactors(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				set[i] = struct{}{}
				for x /= i; x%i == 0; x /= i {
				}
			}
		}
		if x > 1 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$O(n\sqrt{U})$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O\left(\dfrac{U}{\log U}\right)$。至多有 $O\left(\dfrac{U}{\log U}\right)$ 个质因数。
