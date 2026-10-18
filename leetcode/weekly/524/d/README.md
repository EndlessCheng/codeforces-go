从左到右依次处理（覆盖）每个目标区间 $[\textit{tarL},\textit{tarR}]$。

做法类似 [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)，[我的题解](https://leetcode.cn/problems/jump-game-ii/solutions/2926993/tu-jie-yi-zhang-tu-miao-dong-tiao-yue-yo-h2d4/)。

举个例子，假设我们要覆盖区间 $[5,9]$：

1. 首先，遍历所有左端点 $\le 5$ 的可用区间，求出这些区间的右端点的最大值 $\textit{nextEnd}$。
2. 如果 $\textit{nextEnd} < 5$，那么没法覆盖 $5$，返回 $-1$。
3. 现在，剩下需要覆盖的点变成 $(\textit{nextEnd},9]$。重复上述过程，直到 $\textit{nextEnd}\ge 9$ 为止。此外，如果第一步没有发现新的可用区间，即 $\textit{nextEnd}$ 不变，那么无法继续向后覆盖，此时也应返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV1N2nu6xEPW/?t=18m9s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minSelection(self, availableIntervals: list[list[int]], targetIntervals: list[list[int]]) -> int:
        availableIntervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序
        targetIntervals.sort(key=lambda p: p[0])  # 按照左端点从小到大排序

        cur_end = 0  # 已建造的桥的右端点
        ans = i = 0

        for tar_l, tar_r in targetIntervals:
            while cur_end < tar_r:
                # 当前最靠左的未覆盖位置
                uncover = max(tar_l, cur_end)

                # 寻找下一座桥的最远点
                next_end = cur_end
                while i < len(availableIntervals) and availableIntervals[i][0] <= uncover:
                    next_end = max(next_end, availableIntervals[i][1])
                    i += 1

                # 无法覆盖 uncover，或者无法继续向右延伸
                if next_end < uncover or next_end == cur_end:
                    return -1

                cur_end = next_end  # 建桥后，最远可以到达 next_end
                ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    public int minSelection(int[][] availableIntervals, int[][] targetIntervals) {
        Arrays.sort(availableIntervals, (p, q) -> p[0] - q[0]); // 按照左端点从小到大排序
        Arrays.sort(targetIntervals, (p, q) -> p[0] - q[0]); // 按照左端点从小到大排序

        int curEnd = 0; // 已建造的桥的右端点
        int i = 0;
        int ans = 0;
        for (int[] target : targetIntervals) {
            int tarL = target[0], tarR = target[1];
            while (curEnd < tarR) {
                // 当前最靠左的未覆盖位置
                int uncover = Math.max(tarL, curEnd);

                // 寻找下一座桥的最远点
                int nextEnd = curEnd;
                while (i < availableIntervals.length && availableIntervals[i][0] <= uncover) {
                    nextEnd = Math.max(nextEnd, availableIntervals[i][1]);
                    i++;
                }

                // 无法覆盖 uncover，或者无法继续向右延伸
                if (nextEnd < uncover || nextEnd == curEnd) {
                    return -1;
                }

                curEnd = nextEnd; // 建桥后，最远可以到达 nextEnd
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSelection(vector<vector<int>>& availableIntervals, vector<vector<int>>& targetIntervals) {
        ranges::sort(availableIntervals, {}, [](auto& a) { return a[0]; }); // 按照左端点从小到大排序
        ranges::sort(targetIntervals, {}, [](auto& a) { return a[0]; }); // 按照左端点从小到大排序

        int cur_end = 0; // 已建造的桥的右端点
        int i = 0;
        int ans = 0;
        for (auto& target : targetIntervals) {
            int tar_l = target[0], tar_r = target[1];
            while (cur_end < tar_r) {
                // 当前最靠左的未覆盖位置
                int uncover = max(tar_l, cur_end);

                // 寻找下一座桥的最远点
                int next_end = cur_end;
                while (i < availableIntervals.size() && availableIntervals[i][0] <= uncover) {
                    next_end = max(next_end, availableIntervals[i][1]);
                    i++;
                }

                // 无法覆盖 uncover，或者无法继续向右延伸
                if (next_end < uncover || next_end == cur_end) {
                    return -1;
                }

                cur_end = next_end; // 建桥后，最远可以到达 next_end
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minSelection(availableIntervals, targetIntervals [][]int) (ans int) {
	// 按照左端点从小到大排序
	slices.SortFunc(availableIntervals, func(a, b []int) int { return a[0] - b[0] })
	slices.SortFunc(targetIntervals, func(a, b []int) int { return a[0] - b[0] })

	curEnd := 0 // 已建造的桥的右端点
	i := 0
	for _, target := range targetIntervals {
		tarL, tarR := target[0], target[1]
		for curEnd < tarR {
			// 当前最靠左的未覆盖位置
			uncover := max(tarL, curEnd)

			// 寻找下一座桥的最远点
			nextEnd := curEnd
			for i < len(availableIntervals) && availableIntervals[i][0] <= uncover {
				nextEnd = max(nextEnd, availableIntervals[i][1])
				i++
			}

			// 无法覆盖 uncover，或者无法继续向右延伸
			if nextEnd < uncover || nextEnd == curEnd {
				return -1
			}

			curEnd = nextEnd // 建桥后，最远可以到达 nextEnd
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 是 $\textit{availableIntervals}$ 的长度，$m$ 是 $\textit{targetIntervals}$ 的长度。瓶颈在排序上。虽然写了个三重循环，但最内层的 `i++` 的次数至多 $n$ 次，所以三重循环的时间复杂度为 $\mathcal{O}(n+m)$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面贪心题单的「**§2.5 合并区间**」和「**§2.4 区间覆盖**」。

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
