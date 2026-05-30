设 $\textit{nums}$ 所有元素之和为 $\textit{total}$。

根据题意，有

$$
\textit{leftSum}[i] + \textit{nums}[i] + \textit{rightSum}[i] = \textit{total}
$$

即

$$
\textit{rightSum}[i] = \textit{total} - \textit{leftSum}[i] - \textit{nums}[i]
$$

所以

$$
|\textit{leftSum}[i] - \textit{rightSum}[i]| = |2\cdot \textit{leftSum}[i] + \textit{nums}[i] - \textit{total}| 
$$

代码实现时，可以直接把答案保存到 $\textit{nums}$ 中。

> 如果不想修改输入的话，可以额外创建一个 $\textit{answer}$ 数组保存答案。

```py [sol-Python3]
class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        total = sum(nums)

        left_sum = 0
        for i, x in enumerate(nums):
            nums[i] = abs(left_sum * 2 + x - total)
            left_sum += x
        return nums
```

```java [sol-Java]
class Solution {
    public int[] leftRightDifference(int[] nums) {
        int total = 0;
        for (int x : nums) {
            total += x;
        }

        int leftSum = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            nums[i] = Math.abs(leftSum * 2 + x - total);
            leftSum += x;
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> leftRightDifference(vector<int>& nums) {
        int total = reduce(nums.begin(), nums.end());

        int left_sum = 0;
        for (int& x : nums) {
            left_sum += x; // 这样写的话，left_sum 包含 x
            x = abs(left_sum * 2 - x - total);
        }
        return nums;
    }
};
```

```c [sol-C]
int* leftRightDifference(int* nums, int numsSize, int* returnSize) {
    int total = 0;
    for (int i = 0; i < numsSize; i++) {
        total += nums[i];
    }

    int left_sum = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        nums[i] = abs(left_sum * 2 + x - total);
        left_sum += x;
    }

    *returnSize = numsSize;
    return nums;
}
```

```go [sol-Go]
func leftRightDifference(nums []int) []int {
	total := 0
	for _, x := range nums {
		total += x
	}

	leftSum := 0
	for i, x := range nums {
		nums[i] = abs(leftSum*2 + x - total)
		leftSum += x
	}
	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

```js [sol-JavaScript]
var leftRightDifference = function(nums) {
    const total = _.sum(nums);

    let leftSum = 0;
    for (let i = 0; i < nums.length; i++) {
        const x = nums[i];
        nums[i] = Math.abs(leftSum * 2 + x - total);
        leftSum += x;
    }
    return nums;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn left_right_difference(mut nums: Vec<i32>) -> Vec<i32> {
        let total = nums.iter().sum::<i32>();

        let mut left_sum = 0;
        for x in nums.iter_mut() {
            left_sum += *x; // 这样写的话，left_sum 包含 x
            *x = (left_sum * 2 - *x - total).abs();
        }
        nums
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
