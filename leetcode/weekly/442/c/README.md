为了计算酿造药水的时间，定义 $\textit{lastFinish}[i]$ 表示巫师 $i$ 完成上一瓶药水的时间。

示例 1 在处理完 $\textit{mana}[0]$ 后，有

$$
\textit{lastFinish} = [5,30,40,60]
$$

如果接着 $\textit{lastFinish}$ 继续酿造下一瓶药水 $\textit{mana}[1]=1$，完成时间是多少？注意开始酿造的时间不能早于 $\textit{lastFinish}[i]$。

| $i$  | $\textit{skill}[i]$  | $\textit{lastFinish}[i]$  | 完成时间  |
|---|---|---|---|
| $0$  | $1$  | $5$  |  $5+1=6$ |
| $1$  | $5$  | $30$  |  $\max(6,30)+5=35$ |
| $2$  | $2$  | $40$  |  $\max(35,40)+2=42$ |
| $3$  | $4$  | $60$  |  $\max(42,60)+4=64$ |

题目要求「药水在当前巫师完成工作后必须立即传递给下一个巫师并开始处理」，也就是说，酿造药水的过程中是**不能有停顿**的。

从 $64$ 开始**倒推**，可以得到每名巫师的**实际完成时间**。比如倒数第二位巫师的完成时间，就是 $64$ 减去最后一名巫师花费的时间 $4\cdot 1$，得到 $60$。

| $i$  | $\textit{skill}[i+1]$  | 实际完成时间  |
|---|---|---|
| $3$  |  - |  $64$ |
| $2$  | $4$  |  $64-4\cdot 1=60$ |
| $1$  | $2$  |  $60-2\cdot 1=58$ |
| $0$  | $5$  |  $58-5\cdot 1=53$ |

按照上述过程处理每瓶药水，最终答案为 $\textit{lastFinish}[n-1]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV12eXYYVE5H/?t=7m48s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        last_finish = [0] * n  # 第 i 名巫师完成上一瓶药水的时间
        for m in mana:
            # 按题意模拟
            sum_t = 0
            for x, last in zip(skill, last_finish):
                if last > sum_t: sum_t = last  # 手写 max
                sum_t += x * m
            # 倒推：如果酿造药水的过程中没有停顿，那么 last_finish[i] 应该是多少
            last_finish[-1] = sum_t
            for i in range(n - 2, -1, -1):
                last_finish[i] = last_finish[i + 1] - skill[i + 1] * m
        return last_finish[-1]
```

```java [sol-Java]
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        long[] lastFinish = new long[n]; // 第 i 名巫师完成上一瓶药水的时间
        for (int m : mana) {
            // 按题意模拟
            long sumT = 0;
            for (int i = 0; i < n; i++) {
                sumT = Math.max(sumT, lastFinish[i]) + skill[i] * m;
            }
            // 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
            lastFinish[n - 1] = sumT;
            for (int i = n - 2; i >= 0; i--) {
                lastFinish[i] = lastFinish[i + 1] - skill[i + 1] * m;
            }
        }
        return lastFinish[n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size();
        vector<long long> last_finish(n); // 第 i 名巫师完成上一瓶药水的时间
        for (int m : mana) {
            // 按题意模拟
            long long sum_t = 0;
            for (int i = 0; i < n; i++) {
                sum_t = max(sum_t, last_finish[i]) + skill[i] * m;
            }
            // 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
            last_finish[n - 1] = sum_t;
            for (int i = n - 2; i >= 0; i--) {
                last_finish[i] = last_finish[i + 1] - skill[i + 1] * m;
            }
        }
        return last_finish[n - 1];
    }
};
```

```go [sol-Go]
func minTime(skill, mana []int) int64 {
	n := len(skill)
	lastFinish := make([]int, n) // 第 i 名巫师完成上一瓶药水的时间
	for _, m := range mana {
		// 按题意模拟
		sumT := 0
		for i, x := range skill {
			sumT = max(sumT, lastFinish[i]) + x*m
		}
		// 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
		lastFinish[n-1] = sumT
		for i := n - 2; i >= 0; i-- {
			lastFinish[i] = lastFinish[i+1] - skill[i+1]*m
		}
	}
	return int64(lastFinish[n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1840. 最高建筑高度](https://leetcode.cn/problems/maximum-building-height/)

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
