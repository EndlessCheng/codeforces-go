[视频讲解](https://www.bilibili.com/video/BV1kd4y1q7fC) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

定义 $f[i]$ 表示删除后缀 $s[i:]$ 所需的最大操作数。

根据题意，我们可以枚举删除字母的长度 $j$，如果 $s[i:i+j] = s[i+j:i+2j]$，那么可以删除，此时有转移 $f[i] = f[i+j] + 1$。如果不存在两个子串相等的情况，则 $f[i] = 1$。$f[i]$ 取所有情况的最大值。

倒着计算 $f[i]$，答案为 $f[0]$。

最后，我们需要快速判断两个子串是否相同。这可以用 $O(n^2)$ 的 DP 预处理出来，具体见代码。

```py [sol1-Python3]
class Solution:
    def deleteString(self, s: str) -> int:
        n = len(s)
        if len(set(s)) == 1: return n  # 特判全部相同的情况
        lcp = [[0] * (n + 1) for _ in range(n + 1)]  # lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        for i in range(n - 1, -1, -1):
            for j in range(n - 1, i, -1):
                if s[i] == s[j]:
                    lcp[i][j] = lcp[i + 1][j + 1] + 1
        f = [0] * n
        for i in range(n - 1, -1, -1):
            for j in range(1, (n - i) // 2 + 1):
                if lcp[i][i + j] >= j:  # 说明 s[i:i+j] == s[i+j:i+2*j]
                    f[i] = max(f[i], f[i + j])
            f[i] += 1
        return f[0]
```

```java [sol1-Java]
class Solution {
    public int deleteString(String S) {
        var s = S.toCharArray();
        var n = s.length;
        if (allEqual(s)) return n; // 特判全部相同的情况
        var lcp = new int[n + 1][n + 1]; // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        for (var i = n - 1; i >= 0; --i)
            for (var j = n - 1; j > i; --j)
                if (s[i] == s[j])
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
        var f = new int[n];
        for (var i = n - 1; i >= 0; --i) {
            for (var j = 1; i + j * 2 <= n; ++j)
                if (lcp[i][i + j] >= j) // 说明 s[i:i+j] == s[i+j:i+j*2]
                    f[i] = Math.max(f[i], f[i + j]);
            ++f[i];
        }
        return f[0];
    }

    private boolean allEqual(char[] s) {
        for (var i = 1; i < s.length; i++)
            if (s[i] != s[0])
                return false;
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int deleteString(string s) {
        int n = s.length();
        if (equal(s.begin() + 1, s.end(), s.begin())) // 特判全部相同的情况
            return n;
        int lcp[n + 1][n + 1]; // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        memset(lcp, 0, sizeof(lcp));
        for (int i = n - 1; i >= 0; --i)
            for (int j = n - 1; j > i; --j)
                if (s[i] == s[j])
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
        int f[n];
        memset(f, 0, sizeof(f));
        for (int i = n - 1; i >= 0; --i) {
            for (int j = 1; i + j * 2 <= n; ++j)
                if (lcp[i][i + j] >= j) // 说明 s[i:i+j] == s[i+j:i+j*2]
                    f[i] = max(f[i], f[i + j]);
            ++f[i];
        }
        return f[0];
    }
};
```

```go [sol1-Go]
func deleteString(s string) int {
	n := len(s)
	if allEqual(s) { // 特判全部相同的情况
		return n
	}
	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}
	return f[0]
}

func allEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(n^2)$。
