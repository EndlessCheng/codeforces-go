按题意模拟即可。

可以额外创建一个 $\textit{vis}$ 数组标记访问过的位置，也可以直接把 $\textit{grid}[x][y]$ 改成 $0$ 或者空。

```py [sol-Python3]
class Solution:
    def lastVisitedCell(self, grid: list[list[str]]) -> list[int]:
        m, n = len(grid), len(grid[0])
        x = y = 0
        while True:
            nx, ny = x, y

            ch = grid[x][y]
            if ch == "U":
                nx -= 1
            elif ch == "D":
                nx += 1
            elif ch == "L":
                ny -= 1
            else:
                ny += 1

            if nx < 0 or nx == m or ny < 0 or ny == n or grid[nx][ny] is None:
                return [x, y]

            grid[x][y] = None  # 访问标记
            x, y = nx, ny
```

```java [sol-Java]
class Solution {
    public int[] lastVisitedCell(char[][] grid) {
        int m = grid.length, n = grid[0].length;
        int x = 0, y = 0;
        while (true) {
            int nx = x, ny = y;
            switch (grid[x][y]) {
                case 'U' -> nx--;
                case 'D' -> nx++;
                case 'L' -> ny--;
                default -> ny++;
            }
            if (nx < 0 || nx == m || ny < 0 || ny == n || grid[nx][ny] == 0) {
                return new int[]{x, y};
            }
            grid[x][y] = 0; // 访问标记
            x = nx;
            y = ny;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> lastVisitedCell(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size();
        int x = 0, y = 0;
        while (true) {
            int nx = x, ny = y;
            switch (grid[x][y]) {
                case 'U': nx--; break;
                case 'D': nx++; break;
                case 'L': ny--; break;
                default: ny++;
            }
            if (nx < 0 || nx == m || ny < 0 || ny == n || grid[nx][ny] == 0) {
                return {x, y};
            }
            grid[x][y] = 0; // 访问标记
            x = nx;
            y = ny;
        }
    }
};
```

```go [sol-Go]
func lastVisitedCell(grid [][]byte) []int {
	m, n := len(grid), len(grid[0])
	x, y := 0, 0
	for {
		nx, ny := x, y
		switch grid[x][y] {
		case 'U': nx--
		case 'D': nx++
		case 'L': ny--
		default: ny++
		}
		if nx < 0 || nx == m || ny < 0 || ny == n || grid[nx][ny] == 0 {
			return []int{x, y}
		}
		grid[x][y] = 0 // 访问标记
		x, y = nx, ny
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。注意题目保证 $mn\le 10^5$。
- 空间复杂度：$\mathcal{O}(1)$。

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
