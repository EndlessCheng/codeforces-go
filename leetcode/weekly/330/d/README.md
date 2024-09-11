根据 [132 模式](https://leetcode.cn/problems/132-pattern/) 的取名法，本题可以取名为「**1324 模式**」。

## 方法一

枚举 1324 模式中的 3 和 2，也就是 $j$ 和 $k$ 这两个**中间**的下标。

枚举 3 和 2，那么 4 和 1 有多少个？

- 1324 模式中的 4 的个数：在 $k$ 右侧的比 $x=\textit{nums}[j]$ 大的元素个数，记作 $\textit{great}[k][x]$。
- 1324 模式中的 1 的个数：在 $j$ 左侧的比 $x=\textit{nums}[k]$ 小的元素个数，记作 $\textit{less}[j][x]$。

枚举 $j$ 和 $k$（枚举 1324 模式中的 3 和 2），把满足大小关系的 1 的个数和 4 的个数相乘（乘法原理），即为 1324 模式的个数：

$$
\textit{less}[j][\textit{nums}[k]]\cdot \textit{great}[k][\textit{nums}[j]]
$$

将其加到答案中。

如何计算 $\textit{great}[k][x]$ 和 $\textit{less}[j][x]$？

比如 $\textit{nums}[k]$ 右边有两个数 $3$ 和 $5$，那么：

- 对于 $1$，有两个比 $1$ 大的数，所以 $\textit{great}[k][1] = 2$。
- 对于 $2$，有两个比 $2$ 大的数，所以 $\textit{great}[k][2] = 2$。
- 对于 $3$，有一个比 $3$ 大的数，所以 $\textit{great}[k][3] = 1$。
- 对于 $4$，有一个比 $3$ 大的数，所以 $\textit{great}[k][4] = 1$。

考虑动态规划：

- 如果 $x\ge \textit{nums}[k+1]$，那么「$k$ 右边的比 $x$ 大的数」等同于「$k+1$ 右边的比 $x$ 大的数」，即 $\textit{great}[k][x] = \textit{great}[k+1][x]$。
- 如果 $x< \textit{nums}[k+1]$，那么额外加一，即 $\textit{great}[k][x] = \textit{great}[k+1][x] + 1$。

整理得

$$
\textit{great}[k][x] =
\begin{cases}
\textit{great}[k+1][x], & x\ge \textit{nums}[k+1]     \\
\textit{great}[k+1][x] + 1, & x < \textit{nums}[k+1]     \\
\end{cases}
$$

具体代码怎么写？可以倒序遍历 $\textit{nums}$，遍历到 $k$ 时：

1. 首先把 $\textit{great}[k+1]$ 复制到 $\textit{great}[k]$ 中，这样我们就只需要考虑比 $\textit{nums}[k+1]$ 小的数。
2. 然后把 $\textit{great}[k][1],\textit{great}[k][2],\cdots,\textit{great}[k][\textit{nums}[k+1]-1]$ 都加一。

上面的例子是这样计算的：

1. 对于 $\textit{nums}[n-1]= 5$，把 $\textit{great}[n-2][1],\textit{great}[n-2][2],\textit{great}[n-2][3],\textit{great}[n-2][4]$ 都加一。
2. 对于 $\textit{nums}[n-2]= 3$，首先把 $\textit{great}[n-2]$ 复制到 $\textit{great}[n-3]$ 中，然后把 $\textit{great}[n-3][1],\textit{great}[n-3][2]$ 都加一。更新后，$\textit{great}[n-3]$ 数组的前面几个数分别为 $2,2,1,1$。

对于 $\textit{less}$ 的计算也同理：

$$
\textit{less}[j][x] =
\begin{cases}
\textit{less}[j-1][x], & x\le \textit{nums}[j-1]     \\
\textit{less}[j-1][x] + 1, & x > \textit{nums}[j-1]     \\
\end{cases}
$$

代码实现时，可以在正序枚举 $j$ 的同时计算 $\textit{less}$，这样 $\textit{less}$ 数组的第一个维度可以优化掉。

本题 [视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)。

```py [sol-Python3]
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        n = len(nums)
        great = [None] * n
        great[-1] = [0] * (n + 1)
        for k in range(n - 2, 1, -1):
            great[k] = great[k + 1].copy()  # 也可以写 great[k+1][:]
            for x in range(1, nums[k + 1]):
                great[k][x] += 1

        ans = 0
        less = [0] * (n + 1)
        for j in range(1, n - 1):
            for x in range(nums[j - 1] + 1, n + 1):
                less[x] += 1
            for k in range(j + 1, n - 1):
                if nums[j] > nums[k]:
                    ans += less[nums[k]] * great[k][nums[j]]
        return ans
```

```java [sol-Java]
class Solution {
    public long countQuadruplets(int[] nums) {
        int n = nums.length;
        int[][] great = new int[n][n + 1];
        for (int k = n - 2; k >= 2; k--) {
            great[k] = great[k + 1].clone();
            for (int x = 1; x < nums[k + 1]; x++) {
                great[k][x]++;
            }
        }

        long ans = 0;
        int[] less = new int[n + 1];
        for (int j = 1; j < n - 2; j++) {
            for (int x = nums[j - 1] + 1; x <= n; x++) {
                less[x]++;
            }
            for (int k = j + 1; k < n - 1; k++) {
                if (nums[j] > nums[k]) {
                    ans += less[nums[k]] * great[k][nums[j]];
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countQuadruplets(vector<int> &nums) {
        int n = nums.size();
        vector<vector<int>> great(n, vector<int>(n + 1));
        for (int k = n - 2; k >= 2; k--) {
            great[k] = great[k + 1];
            for (int x = 1; x < nums[k + 1]; x++) {
                great[k][x]++;
            }
        }

        long long ans = 0;
        vector<int> less(n + 1);
        for (int j = 1; j < n - 2; j++) {
            for (int x = nums[j - 1] + 1; x <= n; x++) {
                less[x]++;
            }
            for (int k = j + 1; k < n - 1; k++) {
                if (nums[j] > nums[k]) {
                    ans += less[nums[k]] * great[k][nums[j]];
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countQuadruplets(nums []int) (ans int64) {
    n := len(nums)
    great := make([][]int, n)
    great[n-1] = make([]int, n+1)
    for k := n - 2; k >= 2; k-- {
        great[k] = slices.Clone(great[k+1])
        for x := 1; x < nums[k+1]; x++ {
            great[k][x]++
        }
    }

    less := make([]int, n+1)
    for j := 1; j < n-2; j++ {
        for x := nums[j-1] + 1; x <= n; x++ {
            less[x]++
        }
        for k := j + 1; k < n-1; k++ {
            if nums[j] > nums[k] {
                ans += int64(less[nums[k]] * great[k][nums[j]])
            }
        }
    }
    return
}
```

### 优化

注意 $\textit{nums}$ 是一个 $1$ 到 $n$ 的**排列**，上面的写法并没有利用到这一性质，只利用了 $\textit{nums}$ 的元素范围在 $[1,n]$ 的性质。

设 $x=\textit{nums}[k]$。在 $j$ 右边有 $n-1-j$ 个数，其中 $\textit{great}[j][x]$ 个数比 $x$ 大，由于 $\textit{nums}$ 是一个 $[1,n]$ 的排列，因此 $j$ **右边**有

$$
n-1-j-\textit{great}[j][x]
$$

个不超过 $x$ 的数。

同时，由于总共有 $x$ 个不超过 $x$ 的数，所以 $j$ **左边**有

$$
x - (n-1-j-\textit{great}[j][x])
$$

个不超过 $x$ 的数。又因为 $x$ 在 $j$ 右边，所以上式亦为 $j$ 左边的小于 $x$ 的数的个数。

这样就把 $\textit{less}$ 数组优化掉了。

```py [sol-Python3]
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        n = len(nums)
        great = [None] * n
        great[-1] = [0] * (n + 1)
        for k in range(n - 2, 0, -1):
            great[k] = great[k + 1].copy()
            for x in range(1, nums[k + 1]):
                great[k][x] += 1

        ans = 0
        for j in range(1, n - 1):
            for k in range(j + 1, n - 1):
                x = nums[k]
                if nums[j] > x:
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]]
        return ans
```

```java [sol-Java]
class Solution {
    public long countQuadruplets(int[] nums) {
        int n = nums.length;
        int[][] great = new int[n][n + 1];
        for (int k = n - 2; k > 0; k--) {
            great[k] = great[k + 1].clone();
            for (int x = 1; x < nums[k + 1]; x++) {
                great[k][x]++;
            }
        }

        long ans = 0;
        for (int j = 1; j < n - 2; j++) {
            for (int k = j + 1; k < n - 1; k++) {
                int x = nums[k];
                if (nums[j] > x) {
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]];
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countQuadruplets(vector<int> &nums) {
        int n = nums.size();
        vector<vector<int>> great(n, vector<int>(n + 1));
        for (int k = n - 2; k; k--) {
            great[k] = great[k + 1];
            for (int x = 1; x < nums[k + 1]; x++) {
                great[k][x]++;
            }
        }

        long ans = 0;
        for (int j = 1; j < n - 2; j++) {
            for (int k = j + 1; k < n - 1; k++) {
                if (int x = nums[k]; nums[j] > x) {
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]];
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countQuadruplets(nums []int) (ans int64) {
    n := len(nums)
    great := make([][]int, n)
    great[n-1] = make([]int, n+1)
    for k := n - 2; k > 0; k-- {
        great[k] = slices.Clone(great[k+1])
        for x := 1; x < nums[k+1]; x++ {
            great[k][x]++
        }
    }

    for j := 1; j < n-2; j++ {
        for k := j + 1; k < n-1; k++ {
            x := nums[k]
            if nums[j] > x {
                ans += int64((x - n + 1 + j + great[j][x]) * great[k][nums[j]])
            }
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二

### 核心思想

枚举右边的 4（下标 $l$），问题变成维护左边的 **132 模式的个数**，注意 132 模式中的 3 必须比 $\textit{nums}[l]$ 小。

如果发现 $\textit{nums}[j] < \textit{nums}[l]$，那么就把 $[0,l-1]$ 中的 3 在下标 $j$ 的 132 模式的个数加入答案。

对于 132 模式，可以枚举中间的 3（下标 $j$），问题变成维护 12 模式的个数。

### 12 模式

枚举 2（下标 $k$）和 1（下标 $i$）：

- 如果 $\textit{nums}[i] < \textit{nums}[k]$，则找到了一个 2 在下标 $k$ 的 12 模式。

### 132 模式

枚举 2（下标 $k$）和 3（下标 $j$）：

- 定义 $\textit{cnt}_3[j]$ 表示 3 的下标为 $j$ 时的 132 模式的个数。
- 定义 $\textit{cnt}_2$ 表示 2 的下标为 $k$ 时的 12 模式的个数。

分类讨论：

- 如果 $\textit{nums}[j] > \textit{nums}[k]$，把 $\textit{cnt}_2$ 个 12 模式加到 $\textit{cnt}_3[j]$ 中。
- 把 $j$ 当作 $i$，如果 $\textit{nums}[i]<\textit{nums}[k]$，我们找到了一个 12 模式，把 $\textit{cnt}_2$ 加一。

⚠**注意**：我们并不需要单独计算 12 模式的个数，而是把 12 模式的计算过程整合到 132 模式的计算过程中。

### 1324 模式

枚举 $l$ 和 $j$，分类讨论：

- 如果 $\textit{nums}[j] < \textit{nums}[l]$，把 $\textit{cnt}_3[j]$ 个 132 模式加到答案中。
- 把 $j$ 当作 $i$，把 $l$ 当作 $k$，如果 $\textit{nums}[i] < \textit{nums}[k]$，我们找到了一个 12 模式，把 $\textit{cnt}_2$ 加一。
- 把 $l$ 当作 $k$，如果 $\textit{nums}[j] > \textit{nums}[k]$，把 $\textit{cnt}_2$ 个 12 模式加到 $\textit{cnt}_3[j]$ 中。

⚠**注意**：我们并不需要单独计算 132 模式的个数，而是把 132 模式的计算过程整合到 1324 模式的计算过程中。

```py [sol-Python3]
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        cnt4 = 0
        cnt3 = [0] * len(nums)
        for l in range(2, len(nums)):
            cnt2 = 0
            for j in range(l):
                if nums[j] < nums[l]:  # 3 < 4
                    cnt4 += cnt3[j]
                    # 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
                    cnt2 += 1
                else:  # 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
                    cnt3[j] += cnt2
        return cnt4
```

```java [sol-Java]
class Solution {
    public long countQuadruplets(int[] nums) {
        long cnt4 = 0;
        int[] cnt3 = new int[nums.length];
        for (int l = 2; l < nums.length; l++) {
            int cnt2 = 0;
            for (int j = 0; j < l; j++) {
                if (nums[j] < nums[l]) { // 3 < 4
                    cnt4 += cnt3[j];
                    // 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
                    cnt2++;
                } else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
                    cnt3[j] += cnt2;
                }
            }
        }
        return cnt4;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countQuadruplets(vector<int>& nums) {
        long long cnt4 = 0;
        vector<int> cnt3(nums.size());
        for (int l = 2; l < nums.size(); l++) {
            int cnt2 = 0;
            for (int j = 0; j < l; j++) {
                if (nums[j] < nums[l]) { // 3 < 4
                    cnt4 += cnt3[j];
                    // 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
                    cnt2++;
                } else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
                    cnt3[j] += cnt2;
                }
            }
        }
        return cnt4;
    }
};
```

```go [sol-Go]
func countQuadruplets(nums []int) (cnt4 int64) {
    cnt3 := make([]int, len(nums))
    for l := 2; l < len(nums); l++ {
        cnt2 := 0
        for j := 0; j < l; j++ {
            if nums[j] < nums[l] { // 3 < 4
                cnt4 += int64(cnt3[j])
                // 把 j 当作 i，把 l 当作 k，现在 nums[i] < nums[k]，即 1 < 2
                cnt2++
            } else { // 把 l 当作 k，现在 nums[j] > nums[k]，即 3 > 2
                cnt3[j] += cnt2
            }
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 总结

在处理这类「多元素问题」时，有「**枚举中间**」和「**枚举右维护左**」两种思路。

关于「枚举中间」，有如下题目：

- [3128. 直角三角形](https://leetcode.cn/problems/right-triangles/) 1541
- [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/) 2304
- [2867. 统计树中的合法路径数目](https://leetcode.cn/problems/count-valid-paths-in-a-tree/) 2428
- [3257. 放三个车的价值之和最大 II](https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/) 2553

关于「枚举右维护左」，见下面数据结构题单的第零章。

读者可以用这两种思路解决 [456. 132 模式](https://leetcode.cn/problems/132-pattern/)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
