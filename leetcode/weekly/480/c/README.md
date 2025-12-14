设 $\textit{balance}$ 中的负数下标是 $i$（题目保证至多一个负数）。

贪心地，移动离 $i$ 近的数，操作次数更少：

- 先把离 $i$ 最近的 $\textit{balance}[i-1]$ 和 $\textit{balance}[i+1]$ 移到 $i$。每个 $1$ 移动 $1$ 次。
- 再把离 $i$ 为 $2$ 的 $\textit{balance}[i-2]$ 和 $\textit{balance}[i+2]$ 移到 $i$。每个 $1$ 移动 $2$ 次。
- 依此类推。
- 设 $s = \textit{balance}[i-d] + \textit{balance}[i+d]$，如果 $s \ge -\textit{balance}[i]$，那么只需把 $-\textit{balance}[i]$ 个 $1$ 移到下标 $i$，每个 $1$ 移动 $d$ 次，结束操作。否则，继续把更远的数移动到 $i$。

特殊情况：

- 如果 $\textit{balance}$ 的总和是负数，无法让所有数都变成非负数，返回 $-1$。
- 如果 $\textit{balance}$ 没有负数，无需操作，返回 $0$。

[本题视频讲解](https://www.bilibili.com/video/BV1a1meBiETs/?t=5m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minMoves(self, balance: List[int]) -> int:
        if sum(balance) < 0:  # 总和必须非负
            return -1

        neg_idx = next((i for i, x in enumerate(balance) if x < 0), -1)
        if neg_idx < 0:  # 没有负数，无需操作
            return 0

        n = len(balance)
        need = -balance[neg_idx]
        ans = 0
        for dis in count(1):  # 把与 neg_idx 相距 dis 的数移到 neg_idx
            s = balance[neg_idx - dis] + balance[(neg_idx + dis) % n]
            if s >= need:
                ans += need * dis  # need 个 1 移动 dis 次
                return ans
            ans += s * dis  # s 个 1 移动 dis 次
            need -= s
```

```java [sol-Java]
class Solution {
    public long minMoves(int[] balance) {
        long total = 0;
        int negIdx = -1;
        for (int i = 0; i < balance.length; i++) {
            int x = balance[i];
            total += x;
            if (x < 0) {
                negIdx = i;
            }
        }

        if (total < 0) { // 总和必须非负
            return -1;
        }
        if (negIdx < 0) { // 没有负数，无需操作
            return 0;
        }

        int n = balance.length;
        int need = -balance[negIdx];
        long ans = 0;
        for (int dis = 1; ; dis++) { // 把与 negIdx 相距 dis 的数移到 negIdx
            int s = balance[(negIdx - dis + n) % n] + balance[(negIdx + dis) % n];
            if (s >= need) {
                ans += (long) need * dis; // need 个 1 移动 dis 次
                return ans;
            }
            ans += (long) s * dis; // s 个 1 移动 dis 次
            need -= s;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minMoves(vector<int>& balance) {
        long long total = 0;
        int neg_idx = -1;
        for (int i = 0; i < balance.size(); i++) {
            int x = balance[i];
            total += x;
            if (x < 0) {
                neg_idx = i;
            }
        }

        if (total < 0) { // 总和必须非负
            return -1;
        }
        if (neg_idx < 0) { // 没有负数，无需操作
            return 0;
        }

        int n = balance.size();
        int need = -balance[neg_idx];
        long long ans = 0;
        for (int dis = 1; ; dis++) { // 把与 neg_idx 相距 dis 的数移到 neg_idx
            int s = balance[(neg_idx - dis + n) % n] + balance[(neg_idx + dis) % n];
            if (s >= need) {
                ans += 1LL * need * dis; // need 个 1 移动 dis 次
                return ans;
            }
            ans += 1LL * s * dis; // s 个 1 移动 dis 次
            need -= s;
        }
    }
};
```

```go [sol-Go]
func minMoves(balance []int) int64 {
	total := 0
	negIdx := -1
	for i, x := range balance {
		total += x
		if x < 0 {
			negIdx = i
		}
	}

	if total < 0 { // 总和必须非负
		return -1
	}
	if negIdx < 0 { // 没有负数，无需操作
		return 0
	}

	n := len(balance)
	need := -balance[negIdx]
	ans := 0
	for dis := 1; ; dis++ { // 把与 negIdx 相距 dis 的数移到 negIdx
		s := balance[(negIdx-dis+n)%n] + balance[(negIdx+dis)%n]
		if s >= need {
			ans += need * dis // need 个 1 移动 dis 次
			return int64(ans)
		}
		ans += s * dis // s 个 1 移动 dis 次
		need -= s
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{balance}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心题单的「**§1.4 从最左/最右开始贪心**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
