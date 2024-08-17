## 题意解读

选一个最长连续子数组，满足子数组元素依次是偶数，奇数，偶数，奇数，……，且元素值均不超过 $\textit{threshold}$。

例如 $\textit{nums}=[2,1,1,4,3,4,2,8],\textit{threshold}=5$，数组可以分成 $[2,1],1,[4,3,4],[2],8$，其中 $[\cdots]$ 是子数组，其余数字不满足要求。所以最长连续子数组的长度是 $3$。

## 分组循环

**适用场景**：按照题目要求，数组会被分割成若干组，且每一组的判断/处理逻辑是一样的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

时间复杂度乍一看是 $\mathcal{O}(n^2)$，但注意变量 $i$ 只会增加，不会重置也不会减少。所以二重循环总共循环 $\mathcal{O}(n)$ 次，所以时间复杂度是 $\mathcal{O}(n)$。

```py [sol-Python3]
class Solution:
    def longestAlternatingSubarray(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        ans = i = 0
        while i < n:
            if nums[i] > threshold or nums[i] % 2:
                i += 1  # 直接跳过
                continue
            start = i  # 记录这一组的开始位置
            i += 1  # 开始位置已经满足要求，从下一个位置开始判断
            while i < n and nums[i] <= threshold and nums[i] % 2 != nums[i - 1] % 2:
                i += 1
            # 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - start)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestAlternatingSubarray(int[] nums, int threshold) {
        int n = nums.length;
        int ans = 0, i = 0;
        while (i < n) {
            if (nums[i] > threshold || nums[i] % 2 != 0) {
                i++; // 直接跳过
                continue;
            }
            int start = i; // 记录这一组的开始位置
            i++; // 开始位置已经满足要求，从下一个位置开始判断
            while (i < n && nums[i] <= threshold && nums[i] % 2 != nums[i - 1] % 2) {
                i++;
            }
            // 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = Math.max(ans, i - start);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestAlternatingSubarray(vector<int> &nums, int threshold) {
        int n = nums.size();
        int ans = 0, i = 0;
        while (i < n) {
            if (nums[i] > threshold || nums[i] % 2) {
                i++; // 直接跳过
                continue;
            }
            int start = i; // 记录这一组的开始位置
            i++; // 开始位置已经满足要求，从下一个位置开始判断
            while (i < n && nums[i] <= threshold && nums[i] % 2 != nums[i - 1] % 2) {
                i++;
            }
            // 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - start);
        }
        return ans; 
    }
};
```

```go [sol-Go]
func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++ // 直接跳过
			continue
		}
		start := i // 记录这一组的开始位置
		i++ // 开始位置已经满足要求，从下一个位置开始判断
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		// 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-start)
	}
	return
}
```

```js [sol-JavaScript]
var longestAlternatingSubarray = function(nums, threshold) {
    const n = nums.length;
    let ans = 0, i = 0;
    while (i < n) {
        if (nums[i] > threshold || nums[i] % 2 !== 0) {
            i++; // 直接跳过
            continue;
        }
        let start = i; // 记录这一组的开始位置
        i++; // 开始位置已经满足要求，从下一个位置开始判断
        while (i < n && nums[i] <= threshold && nums[i] % 2 !== nums[i - 1] % 2) {
            i++;
        }
        // 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
        ans = Math.max(ans, i - start);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_alternating_subarray(nums: Vec<i32>, threshold: i32) -> i32 {
        let n = nums.len();
        let mut ans = 0;
        let mut i = 0;
        while i < n {
            if nums[i] > threshold || nums[i] % 2 != 0 {
                i += 1; // 直接跳过
                continue;
            }
            let start = i; // 记录这一组的开始位置
            i += 1; // 开始位置已经满足要求，从下一个位置开始判断
            while i < n && nums[i] <= threshold && nums[i] % 2 != nums[i - 1] % 2 {
                i += 1;
            }
            // 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = ans.max(i - start);
        }
        ans as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。时间复杂度乍一看是 $\mathcal{O}(n^2)$，但注意变量 $i$ 只会增加，不会重置也不会减少。所以二重循环总共循环 $\mathcal{O}(n)$ 次，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 练习

一般来说，分组循环的模板如下（可根据题目调整）：

```py
n = len(nums)
i = 0
while i < n:
    start = i
    while i < n and ...:
        i += 1
    # 从 start 到 i-1 是一组
    # 下一组从 i 开始，无需 i += 1
```

学会一个模板是远远不够的，需要大量练习才能灵活运用。

- [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/)
- [1869. 哪种连续子字符串更长](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/)
- [1957. 删除字符使字符串变好](https://leetcode.cn/problems/delete-characters-to-make-fancy-string/)
- [2038. 如果相邻两个颜色均相同则删除当前颜色](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/)
- [1759. 统计同质子字符串的数目](https://leetcode.cn/problems/count-number-of-homogenous-substrings/)
- [2110. 股票平滑下跌阶段的数目](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/)
- [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/)
- [1839. 所有元音按顺序排布的最长子字符串](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/)
- [228. 汇总区间](https://leetcode.cn/problems/summary-ranges/)
- [2765. 最长交替子序列](https://leetcode.cn/problems/longest-alternating-subarray/)

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
