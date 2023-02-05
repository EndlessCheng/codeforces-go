下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

模拟。翻转新插入的数字，可以做到 $O(1)$ 额外空间。

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
