统计每个字符的出现次数。

对于小写英文字母：

- 计算 $\texttt{a}$ 和 $\texttt{z}$ 的出现次数的绝对差。
- 计算 $\texttt{b}$ 和 $\texttt{y}$ 的出现次数的绝对差。
- ……
- 计算 $\texttt{m}$ 和 $\texttt{n}$ 的出现次数的绝对差。

对于数字：

- 计算 $\texttt{0}$ 和 $\texttt{9}$ 的出现次数的绝对差。
- 计算 $\texttt{1}$ 和 $\texttt{8}$ 的出现次数的绝对差。
- ……
- 计算 $\texttt{4}$ 和 $\texttt{5}$ 的出现次数的绝对差。

累加这些绝对差，即为答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def mirrorFrequency(self, s: str) -> int:
        cnt = Counter(s)

        ans = 0
        for i in range(13):
            ans += abs(cnt[ascii_lowercase[i]] - cnt[ascii_lowercase[-1 - i]])
        for i in range(5):
            ans += abs(cnt[digits[i]] - cnt[digits[-1 - i]])
        return ans
```

```java [sol-Java]
class Solution {
    public int mirrorFrequency(String s) {
        int[] cnt = new int['z' + 1];
        for (char ch : s.toCharArray()) {
            cnt[ch]++;
        }

        int ans = 0;
        for (int i = 0; i < 13; i++) {
            ans += Math.abs(cnt['a' + i] - cnt['z' - i]);
        }
        for (int i = 0; i < 5; i++) {
            ans += Math.abs(cnt['0' + i] - cnt['9' - i]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int mirrorFrequency(string s) {
        int cnt['z' + 1]{};
        for (char ch : s) {
            cnt[ch]++;
        }

        int ans = 0;
        for (int i = 0; i < 13; i++) {
            ans += abs(cnt['a' + i] - cnt['z' - i]);
        }
        for (int i = 0; i < 5; i++) {
            ans += abs(cnt['0' + i] - cnt['9' - i]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func mirrorFrequency(s string) (ans int) {
	cnt := ['z' + 1]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	for i := range 13 {
		ans += abs(cnt['a'+i] - cnt['z'-i])
	}
	for i := range 5 {
		ans += abs(cnt['0'+i] - cnt['9'-i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$。取决于实现。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
