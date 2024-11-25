字符串相当于值域为 $[0,25]$ 的数组。如果把本题改成值域为 $[0,10^5]$ 的数组，暴力做法就无法通过了，如何解决这种更一般的情况呢？

我们可以预处理 $\textit{nextCost}$ 和 $\textit{previousCost}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) $\textit{nxtSum}$ 和 $\textit{preSum}$，从而加速操作代价和的计算。

考虑到字母表是环形的，可以把前缀和数组延长一倍，从而变成非环形的。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1uzBxYoEJC/?t=1m24s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def shiftDistance(self, s: str, t: str, nextCost: List[int], previousCost: List[int]) -> int:
        SIGMA = 26
        nxt_sum = [0] * (SIGMA * 2 + 1)
        pre_sum = [0] * (SIGMA * 2 + 1)
        for i in range(SIGMA * 2):
            nxt_sum[i + 1] = nxt_sum[i] + nextCost[i % SIGMA]
            pre_sum[i + 1] = pre_sum[i] + previousCost[i % SIGMA]

        ans = 0
        ord_a = ord('a')
        for x, y in zip(s, t):
            x = ord(x) - ord_a
            y = ord(y) - ord_a
            ans += min(nxt_sum[y + SIGMA if y < x else y] - nxt_sum[x],
                       pre_sum[(x + SIGMA if x < y else x) + 1] - pre_sum[y + 1])
        return ans
```

```java [sol-Java]
class Solution {
    public long shiftDistance(String s, String t, int[] nextCost, int[] previousCost) {
        final int SIGMA = 26;
        long[] nxtSum = new long[SIGMA * 2 + 1];
        long[] preSum = new long[SIGMA * 2 + 1];
        for (int i = 0; i < SIGMA * 2; i++) {
            nxtSum[i + 1] = nxtSum[i] + nextCost[i % SIGMA];
            preSum[i + 1] = preSum[i] + previousCost[i % SIGMA];
        }

        long ans = 0;
        for (int i = 0; i < s.length(); i++) {
            int x = s.charAt(i) - 'a';
            int y = t.charAt(i) - 'a';
            ans += Math.min(nxtSum[y < x ? y + SIGMA : y] - nxtSum[x],
                            preSum[(x < y ? x + SIGMA : x) + 1] - preSum[y + 1]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long shiftDistance(string s, string t, vector<int>& nextCost, vector<int>& previousCost) {
        const int SIGMA = 26;
        long long nxt_sum[SIGMA * 2 + 1]{}, pre_sum[SIGMA * 2 + 1]{};
        for (int i = 0; i < SIGMA * 2; i++) {
            nxt_sum[i + 1] = nxt_sum[i] + nextCost[i % SIGMA];
            pre_sum[i + 1] = pre_sum[i] + previousCost[i % SIGMA];
        }

        long long ans = 0;
        for (int i = 0; i < s.length(); i++) {
            int x = s[i] - 'a', y = t[i] - 'a';
            ans += min(nxt_sum[y < x ? y + SIGMA : y] - nxt_sum[x],
                       pre_sum[(x < y ? x + SIGMA : x) + 1] - pre_sum[y + 1]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func shiftDistance(s, t string, nextCost, previousCost []int) (ans int64) {
	const sigma = 26
	var nxtSum, preSum [sigma*2 + 1]int64
	for i := 0; i < sigma*2; i++ {
		nxtSum[i+1] = nxtSum[i] + int64(nextCost[i%sigma])
		preSum[i+1] = preSum[i] + int64(previousCost[i%sigma])
	}
	for i := range s {
		x := s[i] - 'a'
		y := t[i] - 'a'
		if y < x {
			y += sigma
		}
		res1 := nxtSum[y] - nxtSum[x]
		y = t[i] - 'a'
		if x < y {
			x += sigma
		}
		res2 := preSum[x+1] - preSum[y+1]
		ans += min(res1, res2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
