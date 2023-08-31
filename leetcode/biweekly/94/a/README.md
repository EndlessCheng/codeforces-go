题目相当于计算：在 $-1$ 和 $1$ 之间全是 $0$ 的前提下，最多可以有多少个 $0$。

从左到右遍历 $\textit{forts}$，设当前非零元素下标为 $i$，上一个非零元素下标为 $\textit{pre}$。如果 $\textit{forts}[i]\ne \textit{forts}[\textit{pre}]$，说明其中一个是 $1$，另一个是 $-1$，则可以用中间 $0$ 的个数 $i-\textit{pre}-1$ 更新答案的最大值。

```py [sol-Python3]
class Solution:
    def captureForts(self, forts: List[int]) -> int:
        ans = 0
        pre = -1  # 表示不存在
        for i, x in enumerate(forts):
            if x:  # 1 或 -1
                if pre >= 0 and x != forts[pre]:  # 一个是 1，另一个是 -1
                    ans = max(ans, i - pre - 1)
                pre = i
        return ans
```

```java [sol-Java]
class Solution {
    public int captureForts(int[] forts) {
        int ans = 0;
        int pre = -1; // 表示不存在
        for (int i = 0; i < forts.length; i++) {
            if (forts[i] != 0) {
                if (pre >= 0 && forts[i] != forts[pre]) { // 一个是 1，另一个是 -1
                    ans = Math.max(ans, i - pre - 1);
                }
                pre = i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int captureForts(vector<int>& forts) {
        int ans = 0;
        int pre = -1; // 表示不存在
        for (int i = 0; i < forts.size(); i++) {
            if (forts[i]) {
                if (pre >= 0 && forts[i] != forts[pre]) { // 一个是 1，另一个是 -1
                    ans = max(ans, i - pre - 1);
                }
                pre = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func captureForts(forts []int) (ans int) {
	pre := -1 // 表示不存在
	for i, x := range forts {
		if x != 0 {
			if pre >= 0 && forts[i] != forts[pre] { // 一个是 1，另一个是 -1
				ans = max(ans, i-pre-1)
			}
			pre = i
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

```js [sol-JavaScript]
var captureForts = function(forts) {
    let ans = 0;
    let pre = -1; // 表示不存在
    for (let i = 0; i < forts.length; i++) {
        if (forts[i]) {
            if (pre >= 0 && forts[i] !== forts[pre]) { // 一个是 1，另一个是 -1
                ans = Math.max(ans, i - pre - 1);
            }
            pre = i;
        }
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{forts}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干额外变量。

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
