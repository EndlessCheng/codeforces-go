**前置知识**：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)。

如何想出状态定义？

- 如果 $\textit{nums}$ 的最后两个数相等，那么去掉这两个数，问题变成剩下 $n-2$ 个数能否有效划分。
- 如果 $\textit{nums}$ 的最后三个数相等，那么去掉这三个数，问题变成剩下 $n-3$ 个数能否有效划分。
- 如果 $\textit{nums}$ 的最后三个数是连续递增的，那么去掉这三个数，问题变成剩下 $n-3$ 个数能否有效划分。

我们要解决的问题都形如「$\textit{nums}$ 的前 $i$ 个数能否有效划分」。

于是定义 $f[0] = \texttt{true}$，$f[i+1]$ 表示能否有效划分 $\textit{nums}[0]$ 到 $\textit{nums}[i]$。

根据有效划分的定义，有

$$
f[i+1] = \vee
\begin{cases} 
f[i-1]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1],&i>0\\
f[i-2]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1] = \textit{nums}[i-2],&i>1\\
f[i-2]\ \wedge\ \textit{nums}[i] = \textit{nums}[i-1]+1 = \textit{nums}[i-2]+2,&i>1
\end{cases}
$$

答案为 $f[n]$。

```py [sol-Python3]
class Solution:
    def validPartition(self, nums: List[int]) -> bool:
        n = len(nums)
        f = [True] + [False] * n
        for i, x in enumerate(nums):
            if i > 0 and f[i - 1] and x == nums[i - 1] or \
               i > 1 and f[i - 2] and (x == nums[i - 1] == nums[i - 2] or
                                       x == nums[i - 1] + 1 == nums[i - 2] + 2):
               f[i + 1] = True
        return f[n]
```

```java [sol-Java]
class Solution {
    public boolean validPartition(int[] nums) {
        int n = nums.length;
        boolean[] f = new boolean[n + 1];
        f[0] = true;
        for (int i = 1; i < n; i++) {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool validPartition(vector<int> &nums) {
        int n = nums.size();
        vector<int> f(n + 1);
        f[0] = true;
        for (int i = 1; i < n; i++) {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				                x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
```

```js [sol-JavaScript]
var validPartition = function(nums) {
    const n = nums.length;
    const f = Array(n + 1).fill(false);
    f[0] = true;
    for (let i = 1; i < n; i++) {
        if (f[i - 1] && nums[i] === nums[i - 1] ||
            i > 1 && f[i - 2] && (nums[i] === nums[i - 1] && nums[i] === nums[i - 2] ||
                                  nums[i] === nums[i - 1] + 1 && nums[i] === nums[i - 2] + 2)) {
            f[i + 1] = true;
        }
    }
    return f[n];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn valid_partition(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut f = vec![false; n + 1];
        f[0] = true;
        for i in 1..n {
            if (f[i - 1] && nums[i] == nums[i - 1] ||
                i > 1 && f[i - 2] && (nums[i] == nums[i - 1] && nums[i] == nums[i - 2] ||
                                      nums[i] == nums[i - 1] + 1 && nums[i] == nums[i - 2] + 2)) {
                f[i + 1] = true;
            }
        }
        f[n]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

有多少种有效划分的**方案数**？模 $10^9+7$。

欢迎在评论区发表你的思路/代码。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
