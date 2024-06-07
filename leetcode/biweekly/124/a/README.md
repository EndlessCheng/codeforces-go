设 $s = \textit{nums}[0] + \textit{nums}[1]$。答案初始化为 $1$。

从 $i=3$ 开始，如果 $\textit{nums}[i-1] + \textit{nums}[i] = s$ 就继续，把 $i$ 增加 $2$，答案增加 $1$；否则退出循环。

```py [sol-Python3]
class Solution:
    def maxOperations(self, nums: List[int]) -> int:
        s = nums[0] + nums[1]
        for i in range(3, len(nums), 2):
            if nums[i - 1] + nums[i] != s:
                return i // 2
        return len(nums) // 2
```

```java [sol-Java]
class Solution {
    public int maxOperations(int[] nums) {
        int s = nums[0] + nums[1];
        int ans = 1;
        for (int i = 3; i < nums.length && nums[i - 1] + nums[i] == s; i += 2) {
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxOperations(vector<int> &nums) {
        int s = nums[0] + nums[1];
        int ans = 1;
        for (int i = 3; i < nums.size() && nums[i - 1] + nums[i] == s; i += 2) {
            ans++;
        }
        return ans;
    }
};
```

```c [sol-C]
int maxOperations(int* nums, int numsSize) {
    int s = nums[0] + nums[1];
    int ans = 1;
    for (int i = 3; i < numsSize && nums[i - 1] + nums[i] == s; i += 2) {
        ans++;
    }
    return ans;
}
```

```go [sol-Go]
func maxOperations(nums []int) int {
	s := nums[0] + nums[1]
	ans := 1
	for i := 3; i < len(nums) && nums[i-1]+nums[i] == s; i += 2 {
		ans++
	}
	return ans
}
```

```js [sol-JavaScript]
var maxOperations = function(nums) {
    const s = nums[0] + nums[1];
    let ans = 1;
    for (let i = 3; i < nums.length && nums[i - 1] + nums[i] === s; i += 2) {
        ans++;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_operations(nums: Vec<i32>) -> i32 {
        let s = nums[0] + nums[1];
        for i in (3..nums.len()).step_by(2) {
            if nums[i - 1] + nums[i] != s {
                return (i / 2) as _;
            }
        }
        return (nums.len() / 2) as _;
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

额外输入一个整数 $k$。每一步操作，你需要从数组中选出和为 $k$ 的两个整数，并将它们移出数组。最多可以执行多少次操作？

见 [1679. K 和数对的最大数目](https://leetcode.cn/problems/max-number-of-k-sum-pairs/)。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
