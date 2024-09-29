思路和周赛第三题一样，采用**前后缀分解**解决。

现在变成了两个问题：

- 第一个问题：对于每个从 $s[i]$ 开始的字符串 $s[i..]$，计算它能匹配 $\textit{pattern}$ 多长的**前缀**。
- 第二个问题：对于每个以 $s[j]$ 结尾的字符串 $s[..j]$，计算它能匹配 $\textit{pattern}$ 多长的**后缀**。

比如示例 1，$s=\texttt{abcdefg},\ \textit{pattern}=\texttt{bcdffg}$，其中：

- $s[1..]$ 可以匹配 $\textit{pattern}$ 长为 $3$ 的前缀。
- $s[..6]$ 可以匹配 $\textit{pattern}$ 长为 $2$ 的后缀。

那么 $3+2=5$ 等于 $\textit{pattern}$ 的长度减一，我们可以修改一个字母使得 $s[1..6]$ 与 $\textit{pattern}$ 相等。

对于第一个问题，我们可以构造字符串 $\textit{pattern} + s$，计算其 Z 数组 $\textit{preZ}$。那么 $s[i..]$ 与 $\textit{pattern}$ 前缀可以匹配的最长长度为 $\textit{preZ}[i+m]$，其中 $m$ 为 $\textit{pattern}$ 的长度。

对于第二个问题，我们可以构造字符串 $\text{rev}(\textit{pattern}) + \text{rev}(s)$，计算其 Z 数组，再反转 Z 数组，得到 $\textit{sufZ}$。其中 $\text{rev}(s)$ 表示 $s$ 反转后的字符串。那么 $s[..j]$ 与 $\textit{pattern}$ 后缀可以匹配的最长长度为 $\textit{sufZ}[j]$。

设 $n$ 为 $s$ 的长度，$m$ 为 $\textit{pattern}$ 的长度。

回到原问题，我们枚举 $i=0,1,\cdots,n-m$，那么当前需要匹配的子串为 $s[i..i+m-1]$，对应的 Z 数组元素为 $\textit{preZ}[i+m]$ 和 $\textit{sufZ}[i+m-1]$。

如果 

$$
\textit{preZ}[i+m] + \textit{sufZ}[i+m-1]\ge m-1
$$

那么答案为 $i$。

代码实现时，也可以枚举 $i=m,m+1,\cdots,n$，这样上面的式子可以简化为

$$
\textit{preZ}[i] + \textit{sufZ}[i-1]\ge m-1
$$

答案为 $i-m$。

如果没有找到匹配，返回 $-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1bjxyewEQV/) 第四题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def calc_z(self, s: str) -> list[int]:
        n = len(s)
        z = [0] * n
        box_l = box_r = 0  # z-box 左右边界
        for i in range(1, n):
            if i <= box_r:
                z[i] = min(z[i - box_l], box_r - i + 1)  # 改成手动 if 可以加快速度
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        return z

    def minStartingIndex(self, s: str, pattern: str) -> int:
        pre_z = self.calc_z(pattern + s)
        suf_z = self.calc_z(pattern[::-1] + s[::-1])
        suf_z.reverse()  # 也可以不反转，下面写 suf_z[-i]
        m = len(pattern)
        for i in range(m, len(s) + 1):
            if pre_z[i] + suf_z[i - 1] >= m - 1:
                return i - m
        return -1
```

```java [sol-Java]
class Solution {
    public int minStartingIndex(String s, String pattern) {
        int[] preZ = calcZ(pattern + s);
        int[] sufZ = calcZ(rev(pattern) + rev(s));
        // 可以不反转 sufZ，下面写 sufZ[sufZ.length - i]
        int n = s.length();
        int m = pattern.length();
        for (int i = m; i <= n; i++) {
            if (preZ[i] + sufZ[sufZ.length - i] >= m - 1) {
                return i - m;
            }
        }
        return -1;
    }

    private int[] calcZ(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] z = new int[n];
        int boxL = 0;
        int boxR = 0; // z-box 左右边界
        for (int i = 1; i < n; i++) {
            if (i <= boxR) {
                z[i] = Math.min(z[i - boxL], boxR - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                boxL = i;
                boxR = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

    private String rev(String s) {
        return new StringBuilder(s).reverse().toString();
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc_z(string s) {
        int n = s.length();
        vector<int> z(n);
        int box_l = 0, box_r = 0; // z-box 左右边界
        for (int i = 1; i < n; i++) {
            if (i <= box_r) {
                z[i] = min(z[i - box_l], box_r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                box_l = i;
                box_r = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

public:
    int minStartingIndex(string s, string pattern) {
        vector<int> pre_z = calc_z(pattern + s);
        ranges::reverse(pattern);
        ranges::reverse(s);
        vector<int> suf_z = calc_z(pattern + s);
        ranges::reverse(suf_z); // 也可以不反转，下面写 suf_z[suf_z.size() - i]
        int m = pattern.length();
        for (int i = m; i <= s.length(); i++) {
            if (pre_z[i] + suf_z[i - 1] >= m - 1) {
                return i - m;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func rev(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}

func minStartingIndex(s, pattern string) int {
	preZ := calcZ(pattern + s)
	sufZ := calcZ(rev(pattern) + rev(s))
	slices.Reverse(sufZ) // 也可以不反转，下面写 sufZ[len(sufZ)-i]
	m := len(pattern)
	for i := m; i <= len(s); i++ {
		if preZ[i]+sufZ[i-1] >= m-1 {
			return i - m
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

题目最下面有个**进阶问题**（连续改至多 $k$ 个），把上式中的 $\ge m-1$ 改成 $\ge m-k$ 即可。

类似的问题还有：使答案存在的最小的 $k$ 是多少？

欢迎在评论区分享你的思路/代码。

## 相似题目

- [3036. 匹配模式数组的子数组数目 II](https://leetcode.cn/problems/number-of-subarrays-that-match-a-pattern-ii/) 1895
- [3008. 找出数组中的美丽下标 II](https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii/) 2016
- [2223. 构造字符串的总得分和](https://leetcode.cn/problems/sum-of-scores-of-built-strings/) 2220
- [3031. 将单词恢复初始状态所需的最短时间 II](https://leetcode.cn/problems/minimum-time-to-revert-word-to-initial-state-ii/) 2278
- [3045. 统计前后缀下标对 II](https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii/) 2328

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
