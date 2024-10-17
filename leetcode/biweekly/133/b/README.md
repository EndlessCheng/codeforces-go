## 题意

给定一个 $01$ 数组 $\textit{nums}$，每次操作，你可以：

- 选一个 $[0,n-3]$ 中的下标 $i$，把 $\textit{nums}[i],\textit{nums}[i+1],\textit{nums}[i+2]$ 都反转，即异或 $1$。

返回把 $\textit{nums}$ 全变成 $1$ 的最小操作次数，如果无法做到则返回 $-1$。

## 思路

讨论是否需要对 $i=0$ 执行操作：

- 如果 $\textit{nums}[0]=1$，不需要操作，问题变成剩下 $n-1$ 个数的子问题。
- 如果 $\textit{nums}[0]=0$，一定要操作，问题变成剩下 $n-1$ 个数的子问题。

接下来，讨论是否需要对 $i=1$ 执行操作，处理方式同上。

依此类推，一直到 $i=n-3$ 处理完后，还剩下 $\textit{nums}[n-2]$ 和 $\textit{nums}[n-1]$，这两个数必须都等于 $1$，否则无法达成题目要求。

## 正确性

**问**：为什么这样做是对的？

**答**：

1. 先操作 $i$ 再操作 $j$（$i\ne j$），和先操作 $j$ 再操作 $i$ 的结果是一样的，所以操作顺序不影响答案。既然操作顺序无影响，我们可以从左到右操作。或者说，假设某种操作顺序是最优的，那么总是可以把这个操作顺序重排成从左到右操作。
2. 对于同一个 $i$，操作两次等于没有操作，所以同一个 $i$ 至多操作一次。注：操作 $i$ 指的是反转 $i,i+1,i+2$ 这三个位置。
3. 结合上述两点，既然同一个 $i$ 至多操作一次，那么从左到右操作的过程中，遇到 $1$ 一定不能操作，遇到 $0$ 一定要操作，所以**从左到右的操作方式有且仅有一种**。
4. 既然操作方式是唯一的，我们只需**模拟**这个过程。

**问**：题目要求的「最少」体现在哪里？

**答**：对同一个 $i$ 至多操作一次，就可以做到最少的操作次数。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        for i in range(len(nums) - 2):
            if nums[i] == 0:  # 必须操作
                nums[i + 1] ^= 1
                nums[i + 2] ^= 1
                ans += 1
        return ans if nums[-2] and nums[-1] else -1
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n - 2; i++) {
            if (nums[i] == 0) { // 必须操作
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans++;
            }
        }
        return nums[n - 2] != 0 && nums[n - 1] != 0 ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n - 2; i++) {
            if (nums[i] == 0) { // 必须操作
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans++;
            }
        }
        return nums[n - 2] && nums[n - 1] ? ans : -1;
    }
};
```

```c [sol-C]
int minOperations(int* nums, int n) {
    int ans = 0;
    for (int i = 0; i < n - 2; i++) {
        if (nums[i] == 0) { // 必须操作
            nums[i + 1] ^= 1;
            nums[i + 2] ^= 1;
            ans++;
        }
    }
    return nums[n - 2] && nums[n - 1] ? ans : -1;
}
```

```go [sol-Go]
func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums[:n-2] {
		if x == 0 { // 必须操作
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ans++
		}
	}
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
	}
	return
}
```

```js [sol-JavaScript]
var minOperations = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 0; i < n - 2; i++) {
        if (nums[i] === 0) { // 必须操作
            nums[i + 1] ^= 1;
            nums[i + 2] ^= 1;
            ans++;
        }
    }
    return nums[n - 2] && nums[n - 1] ? ans : -1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_operations(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut ans = 0;
        for i in 0..n - 2 {
            if nums[i] == 0 { // 必须操作
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans += 1;
            }
        }

        if nums[n - 2] != 0 && nums[n - 1] != 0 {
            ans
        } else {
            -1
        }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度，$k=3$ 为每次操作反转的元素个数。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

把题目中的 $3$ 替换成 $k$，其中 $1\le k \le n$，你能想出一个与 $k$ 无关的 $\mathcal{O}(n)$ 做法吗？

见 [CF1955E](https://codeforces.com/problemset/problem/1955/E)，枚举 $k$，变成上述思考题。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
