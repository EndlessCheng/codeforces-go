## 性质

**前置知识**：[模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)

考察等式

$$
(a+b)\bmod k = (b+c)\bmod k
$$ 

根据前置知识中的恒等式，移项得

$$
(a+b-(b+c)) \bmod k = 0
$$

化简得

$$
(a-c)\bmod k = 0
$$

这意味着 $a$ 与 $c$ 关于模 $k$ **同余**。

根据该性质，有效子序列的奇数项都关于模 $k$ 同余，偶数项都关于模 $k$ 同余。

因此，如果把每个 $\textit{nums}[i]$ 都变成 $\textit{nums}[i]\bmod k$，那么原问题等价于：

- 寻找一个最长的子序列，满足子序列奇数项都相同，偶数项都相同。

## 方法一：考察子序列的最后两项

从左到右遍历 $\textit{nums}$，遍历的同时，维护一个二维数组 $f[y][x]$，表示最后两项模 $k$ 分别为 $y$ 和 $x$ 的子序列的长度。

对于 $x=\textit{nums}[i]\bmod k$，我们可以在「最后两项模 $k$ 分别为 $x$ 和 $y$ 的子序列」的末尾添加 $\textit{nums}[i]$，那么「最后两项模 $k$ 分别为 $y$ 和 $x$ 的子序列」的长度会增加 $1$，即

$$
f[y][x] = f[x][y] + 1
$$

最后答案为 $f[i][j]$ 的最大值。

### 答疑

**问**：如何理解这个递推？它和记忆化搜索的区别是什么？

**答**：对比二者的**计算顺序**。如果用记忆化搜索来做，需要单独计算「最左（或者最右）两项模 $k$ 分别为 $x$ 和 $y$ 的子序列」的长度，这是「单线程」，必须**查找下一个元素的位置**。而递推的计算顺序是，（假设我们先遍历到了元素 $2$，然后遍历到了元素 $4$，两个元素属于不同的子序列）一会计算一下「最后两项模 $k$ 分别为 $y$ 和 $2$ 的子序列」，一会又计算一下「最后两项模 $k$ 分别为 $y$ 和 $4$ 的子序列」，这是「多线程」，**没有查找元素位置的过程，遇到谁就处理谁**。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16w4m1e7y3/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int], k: int) -> int:
        f = [[0] * k for _ in range(k)]
        for x in nums:
            x %= k
            for y, fxy in enumerate(f[x]):
                f[y][x] = fxy + 1
        return max(map(max, f))
```

```py [sol-Python3 写法二]
class Solution:
    def maximumLength(self, nums: List[int], k: int) -> int:
        f = [0] * (k * k)
        for x in nums:
            x %= k
            # f[x * k: (x + 1) * k] 是二维写法的第 x 行
            # f[x::k] 是二维写法的第 x 列
            f[x::k] = [v + 1 for v in f[x * k: (x + 1) * k]]
        return max(f)
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums, int k) {
        int ans = 0;
        int[][] f = new int[k][k];
        for (int x : nums) {
            x %= k;
            for (int y = 0; y < k; y++) {
                f[y][x] = f[x][y] + 1;
                ans = Math.max(ans, f[y][x]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int>& nums, int k) {
        int ans = 0;
        vector<vector<int>> f(k, vector<int>(k));
        for (int x : nums) {
            x %= k;
            for (int y = 0; y < k; y++) {
                f[y][x] = f[x][y] + 1;
                ans = max(ans, f[y][x]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumLength(nums []int, k int) (ans int) {
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			ans = max(ans, f[y][x])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k^2 + nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。注意创建大小为 $k^2$ 的二维数组需要 $\mathcal{O}(k^2)$ 的时间。
- 空间复杂度：$\mathcal{O}(k^2)$。

## 方法二：枚举余数，考察子序列的最后一项

枚举子序列相邻两项模 $k$ 的结果为 $m=0,1,2,\cdots, k-1$。

如果知道了子序列的最后一项（假设是 $x$），那么子序列的倒数第二项也就确定了，即

$$
(m - x\bmod k + k) \bmod k
$$

加 $k$ 再模 $k$ 是为了在 $m < x\bmod k$ 时，保证计算结果非负。

类似方法一，从左到右遍历 $\textit{nums}$ 的同时，维护一个数组 $f[x]$，表示最后一项模 $k$ 为 $x$ 的子序列的长度。

对于 $x=\textit{nums}[i]\bmod k$，我们可以在「最后一项模 $k$ 为 $(m - x\bmod k + k) \bmod k$ 的子序列」的末尾添加 $\textit{nums}[i]$，那么「最后一项模 $k$ 为 $x$ 的子序列」的长度会增加 $1$，即

$$
f[x] = f[(m - x\bmod k + k) \bmod k] + 1
$$

> Python 更简单，由于允许负数下标，可以直接用 $f[m-x\bmod k]$ 作为转移来源。

遍历结束后（或者遍历中），用 $f[i]$ 更新答案的最大值。

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int], k: int) -> int:
        ans = 0
        for m in range(k):
            f = [0] * k
            for x in nums:
                x %= k
                f[x] = f[m - x] + 1
            ans = max(ans, max(f))
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums, int k) {
        int ans = 0;
        for (int m = 0; m < k; m++) {
            int[] f = new int[k];
            for (int x : nums) {
                x %= k;
                f[x] = f[(m - x + k) % k] + 1;
                ans = Math.max(ans, f[x]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int>& nums, int k) {
        int ans = 0;
        for (int m = 0; m < k; m++) {
            vector<int> f(k);
            for (int x : nums) {
                x %= k;
                f[x] = f[(m - x + k) % k] + 1;
                ans = max(ans, f[x]);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumLength(nums []int, k int) (ans int) {
	f := make([]int, k)
	for m := 0; m < k; m++ {
		clear(f)
		for _, x := range nums {
			x %= k
			f[x] = f[(m-x+k)%k] + 1
			ans = max(ans, f[x])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k(k+n))$，其中 $n$ 是 $\textit{nums}$ 的长度。注意创建大小为 $k$ 的数组需要 $\mathcal{O}(k)$ 的时间。
- 空间复杂度：$\mathcal{O}(k)$。

#### 相似题目

见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**§7.2 特殊子序列**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
