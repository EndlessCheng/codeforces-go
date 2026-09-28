如果不操作，答案为 $\textit{nums}$ 中的相邻相等数对的个数，记作 $\textit{base}$。

操作一次，可以让 $\textit{base}$ 增大多少？

把 $\textit{nums}$ 中的 $x$ 替换成 $y$（$x\ne y$）后，只有与 $x$ 相邻的数对的相等关系会发生变化：

- 如果替换前 $x$ 和 $x$ 相邻，那么替换后变成 $y$ 和 $y$ 相邻，不影响答案。
- 如果替换前 $x$ 和 $y$ 相邻，那么替换后变成 $y$ 和 $y$ 相邻，答案增大。
- 如果替换前 $x$ 和 $z$ 相邻（$z\ne x$ 且 $z\ne y$），那么替换后变成 $y$ 和 $z$ 相邻，不影响答案。

所以操作带来的增量恰好等于：

- $\textit{nums}$ 中的相邻且值为 $(x,y)$ 的数对的个数 $\textit{cnt}(x,y)$。

统计相邻不等的数对个数，计算出现次数最大值，最终答案为

$$
\textit{base} + \max_{x < y} \textit{cnt}(x,y)
$$

[本题视频讲解](https://www.bilibili.com/video/BV12gah6UE9b/?t=5m32s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxEqualAdjacentPairs(self, nums: list[int]) -> int:
        base = 0
        cnt = defaultdict(int)

        for x, y in pairwise(nums):
            if x == y:
                base += 1
            else:
                # 把 (x,y) 和 (y,x) 都统一为 (x,y)
                if x > y:
                    x, y = y, x
                # 统计相邻且不相等的数对个数
                cnt[(x, y)] += 1

        max_cnt = max(cnt.values(), default=0)
        return base + max_cnt
```

```java [sol-Java]
class Solution {
    public int maxEqualAdjacentPairs(int[] nums) {
        int base = 0;
        Map<Long, Integer> cnt = new HashMap<>();
        int maxCnt = 0;

        for (int i = 1; i < nums.length; i++) {
            int x = nums[i - 1];
            int y = nums[i];
            if (x == y) {
                base++;
            } else {
                // 把 (x,y) 和 (y,x) 都统一为 (x,y)
                if (x > y) {
                    int tmp = x; // 交换 x 和 y
                    x = y;
                    y = tmp;
                }
                // 统计相邻且不相等的数对个数
                long key = (long) 2e9 * x + y; // 两个 int 合并为一个 long
                int c = cnt.merge(key, 1, Integer::sum); // c = ++cnt[key]
                maxCnt = Math.max(maxCnt, c);
            }
        }

        return base + maxCnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxEqualAdjacentPairs(vector<int>& nums) {
        int base = 0, max_cnt = 0;
        unordered_map<long long, int> cnt;

        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i - 1], y = nums[i];
            if (x == y) {
                base++;
            } else {
                // 把 (x,y) 和 (y,x) 都统一为 (x,y)
                if (x > y) {
                    swap(x, y);
                }
                // 统计相邻且不相等的数对个数
                // 两个 int 合并为一个 long long
                long long key = 1LL * x << 32 | y;
                max_cnt = max(max_cnt, ++cnt[key]);
            }
        }

        return base + max_cnt;
    }
};
```

```go [sol-Go]
func maxEqualAdjacentPairs(nums []int) int {
	base := 0
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	maxCnt := 0

	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x == y {
			base++
		} else {
			// 把 (x,y) 和 (y,x) 都统一为 (x,y)
			if x > y {
				x, y = y, x
			}
			// 统计相邻且不相等的数对个数
			cnt[pair{x, y}]++
			maxCnt = max(maxCnt, cnt[pair{x, y}])
		}
	}

	return base + maxCnt
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
