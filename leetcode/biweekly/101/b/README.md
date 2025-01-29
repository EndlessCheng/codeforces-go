## 前置知识：动态规划入门

请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

## 思路

按照题目要求，把 $s$ 中的字母映射到数字上，设映射后变成了数组 $a$，那么题目相当于求 $a$ 的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)（允许子数组为空）。

定义 $f[i]$ 为以 $a[i]$ 结尾的最大子数组和，考虑是否接在以 $a[i-1]$ 结尾的最大子数组之后：

- 接：$f[i] = f[i-1] + a[i]$
- 不接：$f[i] = a[i]$

二者取最大值，得

$$
f[i] = \max(f[i-1],0) + a[i]
$$

答案为 $\max(f)$。

代码实现时，$f$ 可以用一个变量表示，具体可以看「前置知识」中的讲解。

[本题视频讲解](https://www.bilibili.com/video/BV1Ga4y1M72A/)

```py [sol-Python3]
class Solution:
    def maximumCostSubstring(self, s: str, chars: str, vals: List[int]) -> int:
        mapping = dict(zip(chars, vals))
        ans = f = 0
        for c in s:
            f = max(f, 0) + mapping.get(c, ord(c) - ord('a') + 1)
            ans = max(ans, f)
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def maximumCostSubstring(self, s: str, chars: str, vals: List[int]) -> int:
        mapping = dict(zip(ascii_lowercase, range(1, 27))) | dict(zip(chars, vals))
        ans = f = 0
        for c in s:
            f = max(f, 0) + mapping[c]
            ans = max(ans, f)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumCostSubstring(String s, String chars, int[] vals) {
        int[] mapping = new int[26];
        for (int i = 0; i < 26; i++) {
            mapping[i] = i + 1;
        }
        for (int i = 0; i < vals.length; i++) {
            mapping[chars.charAt(i) - 'a'] = vals[i];
        }

        // 最大子数组和（允许子数组为空）
        int ans = 0;
        int f = 0;
        for (char c : s.toCharArray()) {
            f = Math.max(f, 0) + mapping[c - 'a'];
            ans = Math.max(ans, f);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumCostSubstring(string s, string chars, vector<int>& vals) {
        int mapping[26]{};
        iota(mapping, mapping + 26, 1);
        for (int i = 0; i < chars.length(); i++) {
            mapping[chars[i] - 'a'] = vals[i];
        }

        // 最大子数组和（允许子数组为空）
        int ans = 0, f = 0;
        for (char c : s) {
            f = max(f, 0) + mapping[c - 'a'];
            ans = max(ans, f);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumCostSubstring(s, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}

    // 最大子数组和（允许子数组为空）
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n+|\Sigma|)$，其中 $n$ 为 $\textit{nums}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
