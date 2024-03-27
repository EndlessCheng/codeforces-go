题目说「有交集的区间必须在同一个组内」，我们可以先把有交集的区间，合并成一个大区间。

设合并后有 $m$ 个大区间，那么每个大区间都可以分到第一个组或者第二个组，每个大区间都有 $2$ 个方案。由于不同的大区间之间互相独立，根据**乘法原理**，方案数为 $2^m$。

怎么求出 $m$ 呢？

初始化 $m=0,\ \textit{maxR}=-1$。把区间按照**左端点**从小到大排序，遍历区间，同时维护当前合并的大区间右端点 $\textit{maxR}$：

- 如果当前区间的左端点 $l$ 大于 $\textit{maxR}$，由于我们已经按照左端点排序了，那么后面任何区间都不会和之前的区间有交集，换句话说，产生了一个新的大区间，把 $m$ 加一，同时 $\textit{maxR}$ 更新为当前区间右端点 $r$。
- 否则，当前区间要合并到大区间内，用当前区间右端点 $r$ 更新 $\textit{maxR}$ 的最大值。

代码实现时，也可以在遍历的同时，直接计算答案 $\textit{ans}=2^m$，把「$m$ 加一」改成「$\textit{ans}$ 乘 $2$」。在计算中对 $10^9+7$ 取模。

附：[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/) 第三题。

```py [sol-Python3]
class Solution:
    def countWays(self, ranges: List[List[int]]) -> int:
        ranges.sort(key=lambda p: p[0])
        m, max_r = 0, -1
        for l, r in ranges:
            if l > max_r:  # 无法合并
                m += 1  # 新区间
            max_r = max(max_r, r)  # 合并
        return pow(2, m, 1_000_000_007)
```

```java [sol-Java]
class Solution {
    public int countWays(int[][] ranges) {
        Arrays.sort(ranges, (a, b) -> a[0] - b[0]);
        int ans = 1;
        int maxR = -1;
        for (int[] p : ranges) {
            if (p[0] > maxR) { // 无法合并
                ans = ans * 2 % 1_000_000_007; // 新区间
            }
            maxR = Math.max(maxR, p[1]); // 合并
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWays(vector<vector<int>> &ranges) {
        ranges::sort(ranges, [](auto &a, auto &b) { return a[0] < b[0]; });
        int ans = 1, max_r = -1;
        for (auto &p : ranges) {
            if (p[0] > max_r) { // 无法合并
                ans = ans * 2 % 1'000'000'007; // 新区间
            }
            max_r = max(max_r, p[1]); // 合并
        }
        return ans;
    }
};
```

```go [sol-Go]
func countWays(ranges [][]int) int {
	slices.SortFunc(ranges, func(p, q []int) int { return p[0] - q[0] })
	ans, maxR := 1, -1
	for _, p := range ranges {
		if p[0] > maxR { // 无法合并
			ans = ans * 2 % 1_000_000_007 // 新区间
		}
		maxR = max(maxR, p[1]) // 合并
	}
	return ans
}
```

```js [sol-JavaScript]
var countWays = function(ranges) {
    ranges.sort((a, b) => a[0] - b[0]);
    let ans = 1, maxR = -1;
    for (const [l, r] of ranges) {
        if (l > maxR) { // 无法合并
            ans = ans * 2 % 1_000_000_007; // 新区间
        }
        maxR = Math.max(maxR, r); // 合并
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_ways(mut ranges: Vec<Vec<i32>>) -> i32 {
        ranges.sort_unstable_by(|a, b| a[0].cmp(&b[0]));
        let mut ans = 1;
        let mut max_r = -1;
        for p in &ranges {
            if p[0] > max_r { // 无法合并
                ans = ans * 2 % 1_000_000_007; // 新区间
            }
            max_r = max_r.max(p[1]); // 合并
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{ranges}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时的栈开销，仅用到若干额外变量。

## 题单：合并区间

#### 练习 A

- [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
- [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
- [2963. 统计好分割方案的数目](https://leetcode.cn/problems/count-the-number-of-good-partitions/) 1985
- [2584. 分割数组使乘积互质](https://leetcode.cn/problems/split-the-array-to-make-coprime-products/) 2159
- [2655. 寻找最大长度的未覆盖区间](https://leetcode.cn/problems/find-maximal-uncovered-ranges/)（会员题）

#### 练习 B

- [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)
- [1024. 视频拼接](https://leetcode.cn/problems/video-stitching/) 1746
- [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/) 1885

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
