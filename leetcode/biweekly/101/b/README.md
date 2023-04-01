下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

### 前置知识：动态规划入门

见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

### 思路

按照题目要求，把 $s$ 中的字母映射到数字上，设映射后变成了数组 $a$，那么题目相当于求 $a$ 的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)（允许子数组为空）。

定义 $f[i]$ 为以 $a[i]$ 结尾的最大子数组和，考虑是否接在以 $a[i-1]$ 结尾的最大子数组之后：

- 接：$f[i] = f[i-1] + a[i]$
- 不接：$f[i] = a[i]$

取最大值，则有

$$
f[i] = \max(f[i-1],0) + a[i]
$$

答案为 $\max(f)$。

代码实现时，$f$ 可以用一个变量表示，具体可以看「前置知识」中的讲解。

```py [sol1-Python3]
class Solution:
    def maximumCostSubstring(self, s: str, chars: str, vals: List[int]) -> int:
        mapping = dict(zip(chars, vals))
        ans = f = 0
        for c in s:
            f = max(f, 0) + mapping.get(c, ord(c) - ord('a') + 1)
            ans = max(ans, f)
        return ans
```

```java [sol1-Java]
class Solution {
    public int maximumCostSubstring(String s, String chars, int[] vals) {
        var mapping = new int[26];
        for (int i = 0; i < 26; ++i)
            mapping[i] = i + 1;
        for (int i = 0; i < vals.length; ++i)
            mapping[chars.charAt(i) - 'a'] = vals[i];
        // 最大子段和（允许子数组为空）
        int ans = 0, f = 0;
        for (char c : s.toCharArray()) {
            f = Math.max(f, 0) + mapping[c - 'a'];
            ans = Math.max(ans, f);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int maximumCostSubstring(string s, string chars, vector<int> &vals) {
        int mapping[26]{};
        iota(mapping, mapping + 26, 1);
        for (int i = 0; i < chars.length(); ++i)
            mapping[chars[i] - 'a'] = vals[i];
        // 最大子段和（允许子数组为空）
        int ans = 0, f = 0;
        for (char c: s) {
            f = max(f, 0) + mapping[c - 'a'];
            ans = max(ans, f);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumCostSubstring(s, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n+|\Sigma|)$，其中 $n$ 为 $\textit{nums}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。

### 相似题目

- [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)
