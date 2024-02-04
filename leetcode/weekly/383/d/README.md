请看 [视频讲解](https://www.bilibili.com/video/BV1it421W7D8/)。包含 Z 函数（扩展 KMP）的理论讲解。

[可视化网站](https://personal.utdallas.edu/~besp/demo/John2010/z-algorithm.htm)

下文将 $\textit{word}$ 简记为 $s$。

如果只操作一次，就能让 $s$ 恢复成其初始值，意味着什么？

由于可以往 $s$ 的末尾添加任意字符，这意味着只要 $s[k:]$ 是 $s$ 的前缀，就能让 $s$ 恢复成其初始值，其中 $s[k:]$ 表示从 $s[k]$ 开始的后缀。

例如示例 2 的 $s=\text{abacaba},\ k=4$，由于后缀 $s[4:]=\text{aba}$ 是 $s$ 的前缀，所以只需操作一次。

如果操作一次不行，我们就看 $s[2k:]$ 是否为 $s$ 的前缀。依此类推。

如果任意非空 $s[xk:]$（$x>0$）都不是 $s$ 的前缀（例如示例 3），那么只能操作 $\left\lceil\dfrac{n}{k}\right\rceil$ 次，把 $s$ 的字符全部删除，由于可以添加任意字符，我们可以直接生成一个新的 $s$。

我们通过计算 $s$ 后缀与 $s$ 的 LCP（最长公共前缀）长度，即 Z 函数（扩展 KMP）来判断，如果 LCP 长度大于等于后缀长度，就说明对应操作可以让 $s$ 恢复成其初始值。

```py [sol-Python3]
class Solution:
    def minimumTimeToInitialState(self, s: str, k: int) -> int:
        n = len(s)
        z = [0] * n
        l = r = 0
        for i in range(1, n):
            if i <= r:
                z[i] = min(z[i - l], r - i + 1)
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                l, r = i, i + z[i]
                z[i] += 1
            if i % k == 0 and z[i] >= n - i:
                return i // k
        return (n - 1) // k + 1
```

```java [sol-Java]
class Solution {
    public int minimumTimeToInitialState(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] z = new int[n];
        int l = 0, r = 0;
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = Math.min(z[i - l], r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                l = i;
                r = i + z[i];
                z[i]++;
            }
            if (i % k == 0 && z[i] >= n - i) {
                return i / k;
            }
        }
        return (n - 1) / k + 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumTimeToInitialState(string s, int k) {
        int n = s.size();
        vector<int> z(n);
        int l = 0, r = 0;
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = min(z[i - l], r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                l = i;
                r = i + z[i];
                z[i]++;
            }
            if (i % k == 0 && z[i] >= n - i) {
                return i / k;
            }
        }
        return (n - 1) / k + 1;
    }
};
```

```go [sol-Go]
func minimumTimeToInitialState(s string, k int) int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		if i%k == 0 && z[i] >= n-i {
			return i / k
		}
	}
	return (n-1)/k + 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [2223. 构造字符串的总得分和](https://leetcode.cn/problems/sum-of-scores-of-built-strings/) 2220

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
