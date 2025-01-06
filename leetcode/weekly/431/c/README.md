本题是 [2271. 毯子覆盖的最多白色砖块数](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/) 的带权版本，请先完成那题。

回顾一下，在 [2271 我的题解](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/solutions/1496434/by-endlesscheng-kdy9/) 中，做法是把毯子的右端点和瓷砖右端点对齐，然后跑一遍滑动窗口。

对于本题来说，如果出现两个相邻区间，左边区间 $c$ 大，右边区间 $c$ 小的情况，那么和右端点对齐就不是最优的，**和左端点对齐**反而是最优的。

所以在 2271 题的基础上，额外跑一遍和左端点对齐的滑动窗口即可。

代码实现时，把 $\textit{coins}$ 反转，每个区间 $[l,r]$ 改为 $[-r,-l]$，就可以复用和右端点对齐的代码了。

具体请看 [视频讲解](https://www.bilibili.com/video/BV18srKYLEd8/?t=11m04s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    # 2271. 毯子覆盖的最多白色砖块数
    def maximumWhiteTiles(self, tiles: List[List[int]], carpetLen: int) -> int:
        ans = cover = left = 0
        for tl, tr, c in tiles:
            cover += (tr - tl + 1) * c
            while tiles[left][1] < tr - carpetLen + 1:
                cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
                left += 1
            uncover = max((tr - carpetLen + 1 - tiles[left][0]) * tiles[left][2], 0)
            ans = max(ans, cover - uncover)
        return ans

    def maximumCoins(self, coins: List[List[int]], k: int) -> int:
        coins.sort(key=lambda c: c[0])
        ans = self.maximumWhiteTiles(coins, k)

        coins.reverse()
        for t in coins:
            t[0], t[1] = -t[1], -t[0]
        return max(ans, self.maximumWhiteTiles(coins, k))
```

```java [sol-Java]
class Solution {
    public long maximumCoins(int[][] coins, int k) {
        Arrays.sort(coins, (a, b) -> a[0] - b[0]);
        long ans = maximumWhiteTiles(coins, k);

        // 反转数组
        for (int i = 0, j = coins.length - 1; i < j; i++, j--) {
            int[] tmp = coins[i];
            coins[i] = coins[j];
            coins[j] = tmp;
        }
        // 反转每个区间
        for (int[] t : coins) {
            int tmp = t[0];
            t[0] = -t[1];
            t[1] = -tmp;
        }
        return Math.max(ans, maximumWhiteTiles(coins, k));
    }

    // 2271. 毯子覆盖的最多白色砖块数
    private long maximumWhiteTiles(int[][] tiles, int carpetLen) {
        long ans = 0;
        long cover = 0;
        int left = 0;
        for (int[] tile : tiles) {
            int tl = tile[0], tr = tile[1], c = tile[2];
            cover += (long) (tr - tl + 1) * c;
            while (tiles[left][1] + carpetLen - 1 < tr) {
                cover -= (long) (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2];
                left++;
            }
            long uncover = Math.max((long) (tr - carpetLen + 1 - tiles[left][0]) * tiles[left][2], 0);
            ans = Math.max(ans, cover - uncover);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 2271. 毯子覆盖的最多白色砖块数
    long long maximumWhiteTiles(vector<vector<int>>& tiles, int carpetLen) {
        long long ans = 0, cover = 0;
        int left = 0;
        for (auto& tile : tiles) {
            int tl = tile[0], tr = tile[1], c = tile[2];
            cover += (long long) (tr - tl + 1) * c;
            while (tiles[left][1] + carpetLen - 1 < tr) {
                cover -= (long long) (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2];
                left++;
            }
            long long uncover = max((long long) (tr - carpetLen + 1 - tiles[left][0]) * tiles[left][2], 0LL);
            ans = max(ans, cover - uncover);
        }
        return ans;
    }

public:
    long long maximumCoins(vector<vector<int>>& coins, int k) {
        ranges::sort(coins, {}, [](auto& c) { return c[0]; });
        long long ans = maximumWhiteTiles(coins, k);

        ranges::reverse(coins);
        for (auto& t : coins) {
            int tmp = t[0];
            t[0] = -t[1];
            t[1] = -tmp;
        }
        return max(ans, maximumWhiteTiles(coins, k));
    }
};
```

```go [sol-Go]
// 2271. 毯子覆盖的最多白色砖块数
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	cover, left := 0, 0
	for _, tile := range tiles {
		tl, tr, c := tile[0], tile[1], tile[2]
		cover += (tr - tl + 1) * c
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
			left++
		}
		uncover := max((tr-carpetLen+1-tiles[left][0])*tiles[left][2], 0)
		ans = max(ans, cover-uncover)
	}
	return
}

func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int { return a[0] - b[0] })
	ans := maximumWhiteTiles(coins, k)

	slices.Reverse(coins)
	for _, t := range coins {
		t[0], t[1] = -t[1], -t[0]
	}
	return int64(max(ans, maximumWhiteTiles(coins, k)))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{coins}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

更多相似题目，见下面滑动窗口题单中的「**§2.1 求最长/最大**」。

## 变形题

如果题目没有保证区间互不重叠呢？

**解答**：

可以转换成互不重叠的区间。

例如 $\textit{coins}=[[1,3,1],[2,4,1]]$，这两个区间叠加后，**等同于** $[1,1,1],[2,3,2],[4,4,1]$ 这三个互不重叠区间，这样就可以用上面的算法解决了。

如何计算叠加后的区间有哪些？用差分数组（哈希表），具体可以看数据结构题单。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. 【本题相关】[滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
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
