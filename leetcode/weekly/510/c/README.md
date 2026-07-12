**核心思路**：构造一个 $k\times 2$ 或 $2\times k$ 的全为 $\texttt{.}$ 的区域，这个区域从左上角移动到右下角，恰好有 $k$ 条路径。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def createGrid(self, m: int, n: int, k: int) -> list[str]:
        # 特判
        if k == 4 and m == 3 and n == 3:
            return ["..#", "...", "#.."]

        if m == 1 or n == 1:
            # 一行或一列，只能有一种方案
            if k > 1:
                return []
            return ['.' * n] * m

        # 至少要有 k 行或 k 列（特殊情况上面已判断）
        if m < k and n < k:
            return []

        # 初始全为 '#'（第一列全为 '.'）
        a = [['.'] + ['#'] * (n - 1) for _ in range(m)]
        a[-1] = '.' * n  # 最后一行全为 '.'
        if n >= k:  # 至少有 k 列
            # 倒数第二行开头 k 个 '.'
            for j in range(1, k):
                a[-2][j] = '.'
        else:  # 至少有 k 行
            # 第二列末尾 k 个 '.'
            for i in range(m - k, m - 1):
                a[i][1] = '.'
        return [''.join(row) for row in a]
```

```java [sol-Java]
class Solution {
    public String[] createGrid(int m, int n, int k) {
        // 特判
        if (k == 4 && m == 3 && n == 3) {
            return new String[]{"..#", "...", "#.."};
        }

        if (m == 1 || n == 1) {
            // 一行或一列，只能有一种方案
            if (k > 1) {
                return new String[0];
            }
            // 全为 '.'
            String[] ans = new String[m];
            Arrays.fill(ans, ".".repeat(n));
            return ans;
        }

        // 至少要有 k 行或 k 列（特殊情况上面已判断）
        if (m < k && n < k) {
            return new String[0];
        }

        // 初始全为 '#'
        char[][] a = new char[m][n];
        for (char[] row : a) {
            Arrays.fill(row, '#');
            row[0] = '.'; // 第一列全为 '.'
        }
        Arrays.fill(a[m - 1], '.'); // 最后一行全为 '.'
        if (n >= k) { // 至少有 k 列
            // 倒数第二行开头 k 个 '.'
            for (int j = 1; j < k; j++) {
                a[m - 2][j] = '.';
            }
        } else { // 至少有 k 行
            // 第二列末尾 k 个 '.'
            for (int i = m - k; i < m - 1; i++) {
                a[i][1] = '.';
            }
        }

        String[] ans = new String[m];
        for (int i = 0; i < m; i++) {
            ans[i] = new String(a[i]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> createGrid(int m, int n, int k) {
        // 特判
        if (k == 4 && m == 3 && n == 3) {
            return {"..#", "...", "#.."};
        }

        if (m == 1 || n == 1) {
            // 一行或一列，只能有一种方案
            if (k > 1) {
                return {};
            }
            // 全为 '.'
            return vector(m, string(n, '.'));
        }

        // 至少要有 k 行或 k 列（特殊情况上面已判断）
        if (m < k && n < k) {
            return {};
        }

        // 初始全为 '#'
        vector a(m, string(n, '#'));
        for (auto& row : a) {
            row[0] = '.'; // 第一列全为 '.'
        }
        a[m - 1] = string(n, '.'); // 最后一行全为 '.'
        if (n >= k) { // 至少有 k 列
            // 倒数第二行开头 k 个 '.'
            for (int j = 1; j < k; j++) {
                a[m - 2][j] = '.';
            }
        } else { // 至少有 k 行
            // 第二列末尾 k 个 '.'
            for (int i = m - k; i < m - 1; i++) {
                a[i][1] = '.';
            }
        }
        return a;
    }
};
```

```go [sol-Go]
func createGrid(m, n, k int) []string {
	// 特判
	if k == 4 && m == 3 && n == 3 {
		return []string{"..#", "...", "#.."}
	}

	if m == 1 || n == 1 {
		// 一行或一列，只能有一种方案
		if k > 1 {
			return nil
		}
		// 全为 '.'
		ans := make([]string, m)
		row := strings.Repeat(".", n)
		for i := range ans {
			ans[i] = row
		}
		return ans
	}

	// 至少要有 k 行或 k 列（特殊情况上面已判断）
	if m < k && n < k {
		return nil
	}

	// 初始全为 '#'
	a := make([][]byte, m)
	for i := range m - 1 {
		a[i] = bytes.Repeat([]byte{'#'}, n)
		a[i][0] = '.' // 第一列全为 '.'
	}
	a[m-1] = bytes.Repeat([]byte{'.'}, n) // 最后一行全为 '.'
	if n >= k { // 至少有 k 列 
		// 倒数第二行开头 k 个 '.'
		for j := 1; j < k; j++ {
			a[m-2][j] = '.'
		}
	} else { // 至少有 k 行
		// 第二列末尾 k 个 '.'
		for _, row := range a[m-k : m-1] {
			row[1] = '.'
		}
	}

	ans := make([]string, m)
	for i, row := range a {
		ans[i] = string(row)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 更一般的构造方法

适用于 $m,n,k$ 都很大的情况。

先初始化全为 $\texttt{#}$ 的网格图，最右边那一列改成 $\texttt{.}$ 号。

然后，在对角线上不重叠地放一些 $2\times 2$ 的全为 $\texttt{.}$ 号的小网格（称作「**倍增器**」）。相邻倍增器间隔一行，用一个 L 型管道（三个 $\texttt{.}$ 号）连起来。

对于这样的网格图，从左上角往右下角移动，每经过一个倍增器，方案数乘以 $2$。经过 $i$ 个倍增器的方案数为 $2^i$。

把 $k$ 视作二进制数，拆分成若干个不同的 $2^i$ 之和。

对于每个 $2^i$，从第 $i$ 个倍增器的第二行往右打一条隧道（改成 $\texttt{.}$ 号）通向最右边那列。这会贡献 $2^i$ 个方案数。

## 专题训练

见下面思维题单的「**六、构造题**」。

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
