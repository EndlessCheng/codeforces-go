考虑用下标、绝对差构建状态。

尝试定义 $f[i][j]$ 表示以 $x=\textit{nums}[i]$ 结尾的、最后两个数之差为 $j$ 的子序列的最长长度。注意这意味着子序列倒数第二个数是 $x-j$ 或者 $x+j$。

这已经有 $\mathcal{O}(nD)$ 个状态了（$D$ 为 $\textit{nums}$ 的最大值减最小值），再算上枚举转移来源（枚举子序列倒数第二第三的两数之差），时间复杂度就是 $\mathcal{O}(nD^2)$，会超时。

注意到，题目要求绝对差是非递增的，所以我们枚举的倒数第二第三的两数之差，必须 $\ge j$。如果维护 $f[i]$ 的后缀最大值，就可以 $\mathcal{O}(1)$ 转移了。

不妨直接用 $f[i][j]$ 表示后缀最大值，即定义 $f[i][j]$ 表示以 $\textit{nums}[i]$ 结尾的、最后两个数之差**至少**为 $j$ 的子序列的最长长度。

分类讨论：

- $x=\textit{nums}[i]$ 单独形成一个子序列，那么 $f[i][j] = 1$。
- 子序列最后两数之差严格大于 $j$，也就是至少为 $j+1$，那么 $f[i][j] = f[i][j+1]$。
- 子序列最后两数之差恰好等于 $j$，那么子序列倒数第二个数是 $x-j$ 或者 $x+j$，并且子序列倒数第二第三的两数之差至少为 $j$，那么 $f[i][j] = \max(f[\textit{last}[x-j]][j]+1, f[\textit{last}[x+j]][j]+1)$。其中 $\textit{last}[x]$ 表示 $x$ 上一次出现的下标。

> 注：只考虑上一次出现的下标，是因为 $i$ 越大，子序列就越长。

所有情况取最大值，得

$$
f[i][j] = \max(1, f[i][j+1], f[\textit{last}[x-j]][j]+1, f[\textit{last}[x+j]][j]+1)
$$

> 注：由于 $f[i][j]$ 需要从 $f[i][j+1]$ 转移过来，所以 $j$ 要倒序枚举。

最终答案为所有 $f[i][j]$ 的最大值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SzrAYMESJ/?t=10m18s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        mx = max(nums)
        max_d = mx - min(nums)
        f = [[0] * (max_d + 2) for _ in range(len(nums))]
        last = [-1] * (mx + 1)
        for i, x in enumerate(nums):
            for j in range(max_d, -1, -1):
                f[i][j] = max(f[i][j + 1], 1)
                if x - j >= 0 and last[x - j] >= 0:
                    f[i][j] = max(f[i][j], f[last[x - j]][j] + 1)
                if x + j <= mx and last[x + j] >= 0:
                    f[i][j] = max(f[i][j], f[last[x + j]][j] + 1)
            last[x] = i
        return max(map(max, f))
```

```java [sol-Java]
class Solution {
    public int longestSubsequence(int[] nums) {
        int mx = Arrays.stream(nums).max().getAsInt();
        int maxD = mx - Arrays.stream(nums).min().getAsInt();
        int[][] f = new int[nums.length][maxD + 2];
        int[] last = new int[mx + 1];
        Arrays.fill(last, -1);

        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            for (int j = maxD; j >= 0; j--) {
                f[i][j] = Math.max(f[i][j + 1], 1);
                if (x - j >= 0 && last[x - j] >= 0) {
                    f[i][j] = Math.max(f[i][j], f[last[x - j]][j] + 1);
                }
                if (x + j <= mx && last[x + j] >= 0) {
                    f[i][j] = Math.max(f[i][j], f[last[x + j]][j] + 1);
                }
                ans = Math.max(ans, f[i][j]);
            }
            last[x] = i;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        int n = nums.size();
        int mx = ranges::max(nums);
        int max_d = mx - ranges::min(nums);
        vector f(n, vector<int>(max_d + 2));
        vector<int> last(mx + 1, -1);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            for (int j = max_d; j >= 0; j--) {
                f[i][j] = max(f[i][j + 1], 1);
                if (x - j >= 0 && last[x - j] >= 0) {
                    f[i][j] = max(f[i][j], f[last[x - j]][j] + 1);
                }
                if (x + j <= mx && last[x + j] >= 0) {
                    f[i][j] = max(f[i][j], f[last[x + j]][j] + 1);
                }
                ans = max(ans, f[i][j]);
            }
            last[x] = i;
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubsequence(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, len(nums))
	for i := range f {
		f[i] = make([]int, maxD+2)
	}
	last := make([]int, mx+1)
	for i := range last {
		last[i] = -1
	}

	for i, x := range nums {
		for j := maxD; j >= 0; j-- {
			f[i][j] = max(f[i][j+1], 1)
			if x-j >= 0 && last[x-j] >= 0 {
				f[i][j] = max(f[i][j], f[last[x-j]][j]+1)
			}
			if x+j <= mx && last[x+j] >= 0 {
				f[i][j] = max(f[i][j], f[last[x+j]][j]+1)
			}
			ans = max(ans, f[i][j])
		}
		last[x] = i
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nD)$。

## 优化

$\textit{last}$ 可以去掉，改成定义 $f[x][j]$ 表示以元素 $x$ 结尾的、最后两个数之差**至少**为 $j$ 的子序列的最长长度。

状态转移方程改成

$$
f[x][j] = \max(1, f[x][j+1], f[x-j][j]+1, f[x+j][j]+1)
$$

⚠**注意**：$j=0$ 的时候要单独处理，否则转移方程中的 $+1$ 会重复累加。更简单的做法是，用一个变量 $\textit{fx}$ 表示 $f[x][j]$，具体见代码。

```py [sol-Python3]
class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        mx = max(nums)
        max_d = mx - min(nums)
        f = [[0] * (max_d + 1) for _ in range(mx + 1)]
        for x in nums:
            fx = 1
            for j in range(max_d, -1, -1):
                if x - j >= 0:
                    fx = max(fx, f[x - j][j] + 1)
                if x + j <= mx:
                    fx = max(fx, f[x + j][j] + 1)
                f[x][j] = fx
        return max(map(max, f))
```

```py [sol-Python3 优化]
class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        mx = max(nums)
        max_d = mx - min(nums)
        f = [[0] * (max_d + 1) for _ in range(mx + 1)]
        for x in nums:
            fx = 1
            for j in range(max_d, -1, -1):
                if x - j >= 0 and f[x - j][j] + 1 > fx:
                    fx = f[x - j][j] + 1
                if x + j <= mx and f[x + j][j] + 1 > fx:
                    fx = f[x + j][j] + 1
                f[x][j] = fx
        return max(map(max, f))
```

```java [sol-Java]
class Solution {
    public int longestSubsequence(int[] nums) {
        int mx = Arrays.stream(nums).max().getAsInt();
        int maxD = mx - Arrays.stream(nums).min().getAsInt();
        int[][] f = new int[mx + 1][maxD + 1];

        int ans = 0;
        for (int x : nums) {
            int fx = 1;
            for (int j = maxD; j >= 0; j--) {
                if (x - j >= 0) {
                    fx = Math.max(fx, f[x - j][j] + 1);
                }
                if (x + j <= mx) {
                    fx = Math.max(fx, f[x + j][j] + 1);
                }
                f[x][j] = fx;
                ans = Math.max(ans, fx);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        int mx = ranges::max(nums);
        int max_d = mx - ranges::min(nums);
        vector f(mx + 1, vector<int>(max_d + 1));

        int ans = 0;
        for (int x : nums) {
            int fx = 1;
            for (int j = max_d; j >= 0; j--) {
                if (x - j >= 0) {
                    fx = max(fx, f[x - j][j] + 1);
                }
                if (x + j <= mx) {
                    fx = max(fx, f[x + j][j] + 1);
                }
                f[x][j] = fx;
                ans = max(ans, fx);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubsequence(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, mx+1)
	for i := range f {
		f[i] = make([]int, maxD+1)
	}

	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				fx = max(fx, f[x-j][j]+1)
			}
			if x+j <= mx {
				fx = max(fx, f[x+j][j]+1)
			}
			f[x][j] = fx
			ans = max(ans, fx)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(D^2)$。

更多相似题目，见下面动态规划题单中的「**§7.5 多维 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
