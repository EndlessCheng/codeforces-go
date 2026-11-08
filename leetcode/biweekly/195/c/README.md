推荐先完成「子串 $\texttt{abc}$」的版本 [4102. 统计有效子字符串 I](https://leetcode.cn/problems/count-valid-substrings-i/)，[我的题解](https://leetcode.cn/problems/count-valid-substrings-i/solutions/4054604/mei-ju-you-wei-hu-zuo-pythonjavacgo-by-e-77ly/)。

对于 4102 题（子串），我们只需检查 $i-2,i-1,i$ 三个位置。

对于本题（子序列）：

- 先找到 $i$ 及其左边最近的 $\texttt{c}$ 的位置 $i_2$。
- 然后找到 $i_2$ 左边最近的 $\texttt{b}$ 的位置 $i_1$。
- 最后找到 $i_1$ 左边最近的 $\texttt{a}$ 的位置 $i_0$。如果不存在则 $i_0 = -1$。

那么有效子串的最小左端点为 $i_0+1$，把 $i-i_0$ 加到答案中。

[本题视频讲解](https://www.bilibili.com/video/BV152W76MEX3/?t=3m52s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countValidSubstrings(self, s: str) -> int:
        prev = [[0] * 3 for _ in range(len(s) + 1)]
        last = [0] * 3
        ans = 0
        for i, ch in enumerate(s, 1):  # 下标从 1 开始，这样可以用 0 表示不存在
            last[ord(ch) - ord('a')] = i  # 维护最近的每种字母的下标
            prev[i][:] = last
            left = prev[prev[last[2]][1]][0]  # 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
            ans += i - left  # 子串左端点最小值为 left+1
        return ans
```

```java [sol-Java]
class Solution {
    public long countValidSubstrings(String S) {
        char[] s = S.toCharArray();
        int[][] prev = new int[s.length + 1][3];
        int[] last = new int[3];
        long ans = 0;

        // 下标从 1 开始，这样可以用 0 表示不存在
        for (int i = 1; i <= s.length; i++) {
            last[s[i - 1] - 'a'] = i; // 维护最近的每种字母的下标
            System.arraycopy(last, 0, prev[i], 0, 3);
            int left = prev[prev[last[2]][1]][0]; // 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
            ans += i - left; // 子串左端点最小值为 left+1
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countValidSubstrings(string s) {
        int n = s.size();
        vector<array<int, 3>> prev(n + 1);
        array<int, 3> last{};
        long long ans = 0;

        // 下标从 1 开始，这样可以用 0 表示不存在
        for (int i = 1; i <= n; i++) {
            last[s[i - 1] - 'a'] = i; // 维护最近的每种字母的下标
            prev[i] = last;
            int left = prev[prev[last[2]][1]][0]; // 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
            ans += i - left; // 子串左端点最小值为 left+1
        }
        return ans;
    }
};
```

```go [sol-Go]
func countValidSubstrings(s string) (ans int64) {
	prev := make([][3]int, len(s)+1)
	last := [3]int{}
	for i, ch := range s {
		i++ // 下标从 1 开始，这样可以用 0 表示不存在
		last[ch-'a'] = i // 维护最近的每种字母的下标
		prev[i] = last
		left := prev[prev[last[2]][1]][0] // 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
		ans += int64(i - left) // 子串左端点最小值为 left+1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=3$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。

## 专题训练

见下面字符串题单的「**九、子序列自动机**」。

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
