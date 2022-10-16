分别计算小时和分钟的有效个数，根据乘法原理，答案为这两个个数的乘积。

```py [sol1-Python3]
def count(time: str, limit: int) -> int:
    ans = 0
    for i in range(limit):
        if all(t == '?' or t == c for t, c in zip(time, f"{i:02d}")):
            ans += 1
    return ans

class Solution:
    def countTime(self, time: str) -> int:
        return count(time[:2], 24) * count(time[3:], 60)
```

```go [sol1-Go]
func count(time string, limit int) (ans int) {
next:
	for i := 0; i < limit; i++ {
		for j, c := range fmt.Sprintf("%02d", i) {
			if time[j] != '?' && byte(c) != time[j] {
				continue next
			}
		}
		ans++
	}
	return
}

func countTime(time string) int {
	return count(time[:2], 24) * count(time[3:], 60)
}
```

#### 复杂度分析

- 时间复杂度：$O(24+60)=O(1)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
