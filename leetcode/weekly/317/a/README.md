能被 $3$ 和 $2$ 整除的数，就是能被最小公倍数 $\text{LCM}(3,2)=6$ 整除的数。

遍历一遍统计这些数的和，以及这些数的个数，就能算出答案。

```py [sol1-Python3]
class Solution:
    def averageValue(self, nums: List[int]) -> int:
        s = c = 0
        for x in nums:
            if x % 6 == 0:
                s += x
                c += 1
        return s // c if c else 0
```

```go [sol1-Go]
func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, x := range nums {
		if x%6 == 0 {
			sum += x
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
