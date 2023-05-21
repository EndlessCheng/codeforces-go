下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

判断 $[1,1000]$ 的每个数字 $i$ 是否符合要求，并预处理 $[1,i]$ 内的符合要求的数字和 $\textit{preSum}$。

对于每个数字 $i$，把它转成字符串 $s$ 后，写一个回溯，枚举第一个子串、第二个子串、……，累加所有子串对应的整数值之和 $\textit{sum}$。如果存在 $\textit{sum}=i$，则说明 $i$ 符合要求。

如果你不清楚怎么写这个回溯，可以看 [回溯算法套路①子集型回溯【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

```python [sol1-Python3]
PRE_SUM = [0] * 1001  # 预处理
for i in range(1, 1001):
    s = str(i * i)
    n = len(s)
    def dfs(p: int, sum: int) -> bool:
        if p == n:  # 递归终点
            return sum == i  # i 符合要求
        x = 0
        for j in range(p, n):  # 从 s[p] 到 s[j] 组成的子串
            x = x * 10 + int(s[j])  # 对应的整数值
            if dfs(j + 1, sum + x):
                return True
        return False
    PRE_SUM[i] = PRE_SUM[i - 1] + (i * i if dfs(0, 0) else 0)

class Solution:
    def punishmentNumber(self, n: int) -> int:
        return PRE_SUM[n]
```

```go [sol1-Go]
var preSum [1001]int

func init() { // 预处理
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n { // 递归终点
				return sum == i // i 符合要求
			}
			x := 0
			for j := p; j < n; j++ { // 从 s[p] 到 s[j] 组成的子串
				x = x*10 + int(s[j]-'0') // 对应的整数值
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) { // i 符合要求
			preSum[i] += i * i // 计算前缀和
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
```

#### 复杂度分析

- 时间复杂度：预处理 $\mathcal{O}(U^2)$，其中 $U=1000$。对于数字 $i$，它转成字符串后的长度为 $m=\mathcal{O}(\log i)$，所以回溯需要 $\mathcal{O}(2^m)=\mathcal{O}(i)$ 的时间，所以整个预处理需要 $\mathcal{O}(U^2)$ 的时间。
- 空间复杂度：预处理 $\mathcal{O}(U)$。
