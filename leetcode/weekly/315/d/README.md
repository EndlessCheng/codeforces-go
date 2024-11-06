## 提示 1

首先考虑一个简单的情况，$\textit{nums}$ 的所有元素都在 $[\textit{minK},\textit{maxK}]$ 范围内。

在这种情况下，相当于要统计同时包含 $\textit{minK}$ 和 $\textit{maxK}$ 的子数组的个数。

我们可以枚举子数组的右端点。遍历 $\textit{nums}$，记录 $\textit{minK}$ 上一次出现的位置 $\textit{minI}$，以及 $\textit{maxK}$ 上一次出现的位置 $\textit{maxI}$，当遍历到 $\textit{nums}[i]$ 时，如果
$\textit{minK}$ 和 $\textit{maxK}$ 之前出现过，则左端点 $\le\min(\textit{minI},\textit{maxI})$ 的子数组都是合法的。

以 $i$ 为右端点的合法子数组的个数为 

$$
\min(\textit{minI},\textit{maxI})+1
$$

## 提示 2

回到原问题，由于子数组不能包含在 $[\textit{minK},\textit{maxK}]$ 范围之外的元素，因此我们还需要记录上一个在 $[\textit{minK},\textit{maxK}]$ 范围之外的 $\textit{nums}[i]$ 的下标，记作 $i_0$。

以 $i$ 为右端点的合法子数组的个数为 

$$
\min(\textit{minI},\textit{maxI})-i_0
$$

代码实现时：

- 为方便计算，可以初始化 $\textit{minI},\ \textit{maxI},\ i_0$ 均为 $-1$。
- 如果 $\min(\textit{minI},\textit{maxI})-i_0 < 0$，则表示在 $i_0$ 右侧 $\textit{minK}$ 和 $\textit{maxK}$ 没有同时出现，此时以 $i$ 为右端点的合法子数组的个数为 $0$。

[视频讲解](https://www.bilibili.com/video/BV1Ae4y1i7PM) 第四题。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], min_k: int, max_k: int) -> int:
        ans = 0
        min_i = max_i = i0 = -1
        for i, x in enumerate(nums):
            if x == min_k: min_i = i
            if x == max_k: max_i = i
            if not min_k <= x <= max_k: i0 = i  # 子数组不能包含 nums[i0]
            ans += max(min(min_i, max_i) - i0, 0)
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def countSubarrays(self, nums: List[int], min_k: int, max_k: int) -> int:
        ans = 0
        min_i = max_i = i0 = -1
        for i, x in enumerate(nums):
            if x == min_k: min_i = i
            if x == max_k: max_i = i
            if not min_k <= x <= max_k: i0 = i  # 子数组不能包含 nums[i0]
            j = min_i if min_i < max_i else max_i
            if j > i0: ans += j - i0
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int minK, int maxK) {
        long ans = 0;
        int minI = -1, maxI = -1, i0 = -1;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x == minK) minI = i;
            if (x == maxK) maxI = i;
            if (x < minK || x > maxK) i0 = i; // 子数组不能包含 nums[i0]
            ans += Math.max(Math.min(minI, maxI) - i0, 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int> &nums, int min_k, int max_k) {
        long long ans = 0;
        int n = nums.size(), min_i = -1, max_i = -1, i0 = -1;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x == min_k) min_i = i;
            if (x == max_k) max_i = i;
            if (x < min_k || x > max_k) i0 = i; // 子数组不能包含 nums[i0]
            ans += max(min(min_i, max_i) - i0, 0);
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, minK, maxK int) (ans int64) {
	minI, maxI, i0 := -1, -1, -1
	for i, x := range nums {
		if x == minK {
			minI = i
		}
		if x == maxK {
			maxI = i
		}
		if x < minK || x > maxK {
			i0 = i // 子数组不能包含 nums[i0]
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}
```

```js [sol-JavaScript]
var countSubarrays = function (nums, minK, maxK) {
    let ans = 0, minI = -1, maxI = -1, i0 = -1;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        if (x === minK) minI = i;
        if (x === maxK) maxI = i;
        if (x < minK || x > maxK) i0 = i; // 子数组不能包含 nums[i0]
        ans += Math.max(Math.min(minI, maxI) - i0, 0);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_subarrays(nums: Vec<i32>, min_k: i32, max_k: i32) -> i64 {
        let mut ans = 0i64;
        let mut min_i = -1;
        let mut max_i = -1;
        let mut i0 = -1;
        for (i, &x) in nums.iter().enumerate() {
            let i = i as i32;
            if x == min_k {
                min_i = i;
            }
            if x == max_k {
                max_i = i;
            }
            if x < min_k || x > max_k {
                i0 = i; // 子数组不能包含 nums[i0]
            }
            ans += 0.max(min_i.min(max_i) - i0) as i64;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

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
