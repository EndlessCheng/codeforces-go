为方便描述，下文把 $\textit{skill}$ 简称为 $s$，把 $\textit{station}$ 简称为 $t$。

题目要我们找到一个 $t$ 的子序列，恰好等于 $s$。设子序列在 $t$ 中的下标为 $j_0,j_1\ldots,j_{n-1}$，我们的目标是找到一个最大的差值 $j_{i+1} - j_i$。

考虑枚举 $i$。那么 $j_i$ 越小越好，$j_{i+1}$ 越大越好。

- 「$j_i$ 越小越好」意味着子序列 $j_0,j_1,\ldots,j_i$ 在 $t$ 中的位置尽量靠左。
- 「$j_{i+1}$ 越大越好」意味着子序列 $j_{i+1},j_{i+2},\ldots,j_{n-1}$ 在 $t$ 中的位置尽量靠右。

这是两个独立的问题：

- 计算 $t$ 的最短前缀 $[0,\textit{pre}[i]]$，使得 $s$ 的前缀 $s[0,i]$ 是 $t[0,\textit{pre}[i]]$ 的子序列。
- 计算 $t$ 的最短后缀 $[\textit{suf}[i+1], |t|-1]$，使得 $s$ 的后缀 $s[i+1,|s|-1]$ 是 $t[\textit{suf}[i+1], |t|-1]$ 的子序列。

答案为 

$$
\min_{i=0}^{|s|-2}\textit{suf}[i+1] - \textit{pre}[i]
$$

如何计算 $\textit{pre}$ 和 $\textit{suf}$？见 [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/)，[我的题解](https://leetcode.cn/problems/is-subsequence/solution/jian-ji-xie-fa-pythonjavaccgojsrust-by-e-mz22/)。  

[本题视频讲解](https://www.bilibili.com/video/BV1gRbD6zECR/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumGap(self, s: str, t: str) -> int:
        n = len(s)
        suf = [0] * n  # s[i:] 是 t[suf[i]:] 的子序列
        j = len(t) - 1
        for i in range(n - 1, 0, -1):
            while t[j] != s[i]:  # 题目保证 s 是 t 的子序列，下标不会越界
                j -= 1
            suf[i] = j
            j -= 1

        ans = pre = 0
        for i in range(n - 1):
            while t[pre] != s[i]:
                pre += 1
            # 此时 s[:i+1] 是 t[:pre+1] 的子序列
            # 此时 s[i+1:] 是 t[suf[i+1]:] 的子序列
            ans = max(ans, suf[i + 1] - pre)
            pre += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumGap(String skill, String station) {
        char[] s = skill.toCharArray();
        char[] t = station.toCharArray();

        int n = s.length;
        int[] suf = new int[n]; // s[i,|s|-1] 是 t[suf[i],|t|-1] 的子序列
        int j = t.length - 1;
        for (int i = n - 1; i > 0; i--) {
            while (t[j] != s[i]) { // 题目保证 s 是 t 的子序列，下标不会越界
                j--;
            }
            suf[i] = j;
            j--;
        }

        int ans = 0;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            while (t[pre] != s[i]) {
                pre++;
            }
            // 此时 s[0,i] 是 t[0,pre] 的子序列
            // 此时 s[i+1,|s|-1] 是 t[suf[i+1],|t|-1] 的子序列
            ans = Math.max(ans, suf[i + 1] - pre);
            pre++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumGap(string s, string t) {
        int n = s.size();
        vector<int> suf(n); // s[i,|s|-1] 是 t[suf[i],|t|-1] 的子序列
        int j = t.size() - 1;
        for (int i = n - 1; i > 0; i--) {
            while (t[j] != s[i]) { // 题目保证 s 是 t 的子序列，下标不会越界
                j--;
            }
            suf[i] = j;
            j--;
        }

        int ans = 0;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            while (t[pre] != s[i]) {
                pre++;
            }
            // 此时 s[0,i] 是 t[0,pre] 的子序列
            // 此时 s[i+1,|s|-1] 是 t[suf[i+1],|t|-1] 的子序列
            ans = max(ans, suf[i + 1] - pre);
            pre++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumGap(s, t string) (ans int) {
	n := len(s)
	suf := make([]int, n) // s[i:] 是 t[suf[i]:] 的子序列
	j := len(t) - 1
	for i := n - 1; i > 0; i-- {
		for t[j] != s[i] { // 题目保证 s 是 t 的子序列，下标不会越界
			j--
		}
		suf[i] = j
		j--
	}

	pre := 0
	for i, ch := range s[:n-1] {
		for t[pre] != byte(ch) {
			pre++
		}
		// 此时 s[:i+1] 是 t[:pre+1] 的子序列
		// 此时 s[i+1:] 是 t[suf[i+1]:] 的子序列
		ans = max(ans, suf[i+1]-pre)
		pre++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[3983. 一次替换后的子序列](https://leetcode.cn/problems/subsequence-after-one-replacement/)

## 专题训练

1. 双指针题单的「**§4.2 判断子序列**」。
2. 动态规划题单的「**专题：前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
