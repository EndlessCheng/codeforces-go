请看 [视频讲解](https://www.bilibili.com/video/BV1oC4y1o7Tz/) 第三题。

我们需要：

- 先满足子数组分数之和尽量小；
- 再满足分割出的子数组尽量多。

## 提示 1

AND 的性质是，参与 AND 的数越多，结果越小。

## 提示 2

子数组的 AND，不会低于整个 $\textit{nums}$ 数组的 AND。

## 提示 3

如果 $\textit{nums}$ 数组的 AND（记作 $a$）是大于 $0$ 的，那么只能分割出一个数组（即 $\textit{nums}$ 数组）。根据提示 2，如果分割出超过一个数组，那么得分至少为 $2a$，而这是大于 $a$ 的，不满足「子数组分数之和尽量小」的要求。所以在 $a>0$ 的情况下，答案为 $1$。

如果 $\textit{nums}$ 数组的 AND 是等于 $0$ 的，那么可以分割出尽量多的 AND 等于 $0$ 的子数组。怎么分？从左到右遍历数组，只要发现 AND 等于 $0$ 就立刻分割。如果不立刻分割，由于 AND 的数越多越能为 $0$，现在多分了一个数，后面就要少分一个数，可能后面就不能为 $0$ 了。注意，如果最后剩下的一段子数组的 AND 大于 $0$，这一段可以直接合并到前一个 AND 为 $0$ 的子数组中。

```py [sol-Python3]
class Solution:
    def maxSubarrays(self, nums: List[int]) -> int:
        ans = 0
        a = -1  # -1 就是 111...1，和任何数 AND 都等于那个数
        for x in nums:
            a &= x
            if a == 0:
                ans += 1  # 分割
                a = -1
        return max(ans, 1)  # 如果 ans=0 说明所有数的 and>0，答案为 1
```

```java [sol-Java]
class Solution {
    public int maxSubarrays(int[] nums) {
        int ans = 0;
        int a = -1; // -1 就是 111...1，和任何数 AND 都等于那个数
        for (int x : nums) {
            a &= x;
            if (a == 0) {
                ans++; // 分割
                a = -1;
            }
        }
        return Math.max(ans, 1); // 如果 ans=0 说明所有数的 and>0，答案为 1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubarrays(vector<int>& nums) {
        int ans = 0;
        int a = -1; // -1 就是 111...1，和任何数 AND 都等于那个数
        for (int x : nums) {
            a &= x;
            if (a == 0) {
                ans++; // 分割
                a = -1;
            }
        }
        return max(ans, 1); // 如果 ans=0 说明所有数的 and>0，答案为 1
    }
};
```

```go [sol-Go]
func maxSubarrays(nums []int) (ans int) {
	and := -1 // -1 就是 111...1，和任何数 AND 都等于那个数
	for _, x := range nums {
		and &= x
		if and == 0 {
			ans++ // 分割
			and = -1
		}
	}
	return max(ans, 1) // 如果 ans=0 说明所有数的 and>0，答案为 1
}
```

```js [sol-JavaScript]
var maxSubarrays = function(nums) {
    let ans = 0;
    let a = -1; // -1 就是 111...1，和任何数 AND 都等于那个数
    for (const x of nums) {
        a &= x;
        if (a === 0) {
            ans++; // 分割
            a = -1;
        }
    }
    return Math.max(ans, 1); // 如果 ans=0 说明所有数的 and>0，答案为 1
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_subarrays(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut a = -1; // -1 就是 111...1，和任何数 AND 都等于那个数
        for x in nums {
            a &= x;
            if a == 0 {
                ans += 1; // 分割
                a = -1;
            }
        }
        ans.max(1) // 如果 ans=0 说明所有数的 and>0，答案为 1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
