模拟，枚举每个数的数位。

附：[视频讲解](https://www.bilibili.com/video/BV1rM4y1X7z9/)

### 一行写法

```Python [sol1-Python3]
class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        return [d for x in nums for d in map(int, str(x))]
```

### O(1) 额外空间写法

从低到高枚举数位，翻转新插入的数字。

```Python [sol1-Python3]
class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        ans = []
        for x in nums:
            i0 = len(ans)
            while x:
                ans.append(x % 10)
                x //= 10
            ans[i0:] = ans[i0:][::-1]  # 忽略切片开销（毕竟你可以手动反转）
        return ans
```

```go [sol1-Go]
func separateDigits(nums []int) (ans []int) {
	for _, x := range nums {
		i0 := len(ans)
		for ; x > 0; x /= 10 {
			ans = append(ans, x%10)
		}
		b := ans[i0:]
		for i, n := 0, len(b); i < n/2; i++ {
			b[i], b[n-1-i] = b[n-1-i], b[i]
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$。不考虑返回值，仅用到若干额外变量。
