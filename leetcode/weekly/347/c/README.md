请看 [视频讲解](https://www.bilibili.com/video/BV1fo4y1T7MQ/) 第三题，欢迎点赞！

### 提示 1

如果 $s[i-1]\ne s[i]$，那么必须反转，不然没法都相等：

- 要么翻转 $s[0]$ 到 $s[i-1]$，成本为 $i$；
- 要么翻转 $s[i]$ 到 $s[n-1]$，成本为 $n-i$。

这两种情况取最小值。

### 提示 2

从左到右遍历 $s$，如果 $s[i-1]\ne s[i]$，那么必须反转，规则同上。

反转后：

- $s[i]$ 及其左边的字符，都已经相等了。
- $s[i]$ 右边的每对相邻字符，**反转前不同的，反转后仍然不同**。所以要继续反转。

所以累加每次反转的成本，即为答案。

```py [sol-Python3]
class Solution:
    def minimumCost(self, s: str) -> int:
        return sum(min(i, len(s) - i) for i, (x, y) in enumerate(pairwise(s), 1) if x != y)
```

```java [sol-Java]
class Solution {
    public long minimumCost(String S) {
        long ans = 0;
        char[] s = S.toCharArray();
        int n = s.length;
        for (int i = 1; i < n; i++)
            if (s[i - 1] != s[i])
                ans += Math.min(i, n - i);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(string s) {
        long long ans = 0;
        int n = s.length();
        for (int i = 1; i < n; i++)
            if (s[i - 1] != s[i])
                ans += min(i, n - i);
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(s string) (ans int64) {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i-1] != s[i] {
			ans += int64(min(i, n-i))
		}
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
```

```js [sol-JavaScript]
var minimumCost = function (s) {
    const n = s.length;
    let ans = 0;
    for (let i = 1; i < n; i++)
        if (s[i - 1] !== s[i])
            ans += Math.min(i, n - i);
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
