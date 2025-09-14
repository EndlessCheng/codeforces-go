题目要求：同一支队伍不能在连续两场比赛中。对于连续两场比赛的队伍 $(a,b)$ 和 $(c,d)$，$a,b,c,d$ 必须互不相同，所以 $n$ 至少是 $4$。

对于 $n=4$ 的情况，可以程序验证所有 $12!$ 的比赛排列，证明这是无解的。

下面说明 $n\ge 5$ 的构造方案。

**核心思路**：枚举 $d=1,2,3,\dots,n-1$，枚举 $a=0,1,2,\dots,n-1$，构造比赛 $(a,(a+d)\bmod n)$。单独处理 $d=1$ 和 $d=n-1$ 的特殊情况。

对于 $n$ 为奇数的情况，以 $n=5$ 为例：

- $d=1$：构造方案为 $(0,1),(2,3),(4,0),(1,2),(3,4)$。即先安排 $a$ 为偶数的比赛，再安排 $a$ 为奇数的比赛。
- $d=2$：构造方案为 $(0,2),(1,3),(2,4),(3,0),(4,1)$。
- $d=3$：构造方案为 $(0,3),(1,4),(2,0),(3,1),(4,2)$。
- $d=4$：构造方案为 $(1,0),(3,2),(0,4),(2,1),(4,3)$。即先安排 $a$ 为奇数的比赛，再安排 $a$ 为偶数的比赛。这样可以不与 $d=3$ 的最后一场比赛冲突。

对于 $n$ 为偶数的情况，需要微调 $d=1$ 和 $d=n-1$ 中的各一对比赛。以 $n=6$ 为例：

- 对于 $d=1$，如果按照上述方案构造，我们会得到 $(0,1),(2,3),(4,5),(1,2),(3,4),(5,0)$，末尾的 $(5,0)$ 和 $d=2$ 的第一场比赛 $(0,2)$ 冲突了。解决方案：交换 $d=1$ 的最后两场比赛，得到 $(0,1),(2,3),(4,5),(1,2),(5,0),(3,4)$。
- 对于 $d=n-1=5$，如果按照上述方案构造，我们会得到 $(1,0),(3,2),(5,4),(0,5),(2,1),(4,3)$，其中 $a$ 为奇数时的最后一场比赛 $(5,4)$ 与 $a$ 为偶数时的第一场比赛 $(0,5)$ 冲突了。解决方案：交换 $a$ 为奇数时的最后两场比赛，得到 $(1,0),(5,4),(3,2),(0,5),(2,1),(4,3)$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def generateSchedule(self, n: int) -> List[List[int]]:
        if n < 5:
            return []

        ans = []

        # 单独处理 d=1
        for i in range(0, n, 2):
            ans.append([i, (i + 1) % n])
        for i in range(1, n, 2):
            ans.append([i, (i + 1) % n])
        if n % 2 == 0:  # 保证 d=1 的最后一场比赛与 d=2 的第一场比赛无冲突
            ans[-1], ans[-2] = ans[-2], ans[-1]

        # 处理 d=2,3,...,n-2
        for d in range(2, n - 1):
            for i in range(n):
                ans.append([i, (i + d) % n])

        # 单独处理 d=n-1
        for i in range(1, n, 2):
            ans.append([i, (i + n - 1) % n])
        if n % 2 == 0:  # 保证 i 为奇数时的最后一场比赛与 i 为偶数时的第一场比赛无冲突
            ans[-1], ans[-2] = ans[-2], ans[-1]
        for i in range(0, n, 2):
            ans.append([i, (i + n - 1) % n])

        return ans
```

```java [sol-Java]
class Solution {
    public int[][] generateSchedule(int n) {
        if (n < 5) {
            return new int[][]{};
        }

        int[][] ans = new int[n * (n - 1)][2];
        int idx = 0;

        // 单独处理 d=1
        for (int i = 0; i < n; i += 2) {
            ans[idx++] = new int[]{i, (i + 1) % n};
        }
        for (int i = 1; i < n; i += 2) {
            ans[idx++] = new int[]{i, (i + 1) % n};
        }
        if (n % 2 == 0) { // 保证 d=1 的最后一场比赛与 d=2 的第一场比赛无冲突
            swap(ans, idx - 1, idx - 2);
        }

        // 处理 d=2,3,...,n-2
        for (int d = 2; d < n - 1; d++) {
            for (int i = 0; i < n; i++) {
                ans[idx++] = new int[]{i, (i + d) % n};
            }
        }

        // 单独处理 d=n-1
        for (int i = 1; i < n; i += 2) {
            ans[idx++] = new int[]{i, (i + n - 1) % n};
        }
        if (n % 2 == 0) { // 保证 i 为奇数时的最后一场比赛与 i 为偶数时的第一场比赛无冲突
            swap(ans, idx - 1, idx - 2);
        }
        for (int i = 0; i < n; i += 2) {
            ans[idx++] = new int[]{i, (i + n - 1) % n};
        }

        return ans;
    }

    private void swap(int[][] a, int i, int j) {
        int[] tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> generateSchedule(int n) {
        if (n < 5) {
            return {};
        }

        vector<vector<int>> ans;

        // 单独处理 d=1
        for (int i = 0; i < n; i += 2) {
            ans.push_back({i, (i + 1) % n});
        }
        for (int i = 1; i < n; i += 2) {
            ans.push_back({i, (i + 1) % n});
        }
        if (n % 2 == 0) { // 保证 d=1 的最后一场比赛与 d=2 的第一场比赛无冲突
            swap(ans.back(), ans[ans.size() - 2]);
        }

        // 处理 d=2,3,...,n-2
        for (int d = 2; d < n - 1; d++) {
            for (int i = 0; i < n; i++) {
                ans.push_back({i, (i + d) % n});
            }
        }

        // 单独处理 d=n-1
        for (int i = 1; i < n; i += 2) {
            ans.push_back({i, (i + n - 1) % n});
        }
        if (n % 2 == 0) { // 保证 i 为奇数时的最后一场比赛与 i 为偶数时的第一场比赛无冲突
            swap(ans.back(), ans[ans.size() - 2]);
        }
        for (int i = 0; i < n; i += 2) {
            ans.push_back({i, (i + n - 1) % n});
        }

        return ans;
    }
};
```

```go [sol-Go]
func generateSchedule(n int) [][]int {
	if n < 5 {
		return nil
	}

	ans := make([][]int, 0, n*(n-1)) // 预分配空间

	// 单独处理 d=1
	for i := 0; i < n; i += 2 {
		ans = append(ans, []int{i, (i + 1) % n})
	}
	for i := 1; i < n; i += 2 {
		ans = append(ans, []int{i, (i + 1) % n})
	}
	if n%2 == 0 { // 保证 d=1 的最后一场比赛与 d=2 的第一场比赛无冲突
		ans[len(ans)-1], ans[len(ans)-2] = ans[len(ans)-2], ans[len(ans)-1]
	}

	// 处理 d=2,3,...,n-2
	for d := 2; d < n-1; d++ {
		for i := range n {
			ans = append(ans, []int{i, (i + d) % n})
		}
	}

	// 单独处理 d=n-1
	for i := 1; i < n; i += 2 {
		ans = append(ans, []int{i, (i + n - 1) % n})
	}
	if n%2 == 0 { // 保证 i 为奇数时的最后一场比赛与 i 为偶数时的第一场比赛无冲突
		ans[len(ans)-1], ans[len(ans)-2] = ans[len(ans)-2], ans[len(ans)-1]
	}
	for i := 0; i < n; i += 2 {
		ans = append(ans, []int{i, (i + n - 1) % n})
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

本题构造方式不止一种，欢迎在评论区分享你的构造方案。

## 专题训练

见下面贪心与思维题单的「**六、构造题**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
