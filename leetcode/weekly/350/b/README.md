由于分区值是 $\textit{nums}$ 中的某两个数的绝对差，所以答案的**理论最小值**，是 $\textit{nums}$ 任意两个元素绝对差的最小值。实际答案能否取到这个最小值呢？

可以的，构造方案如下：

1. 把 $\textit{nums}$ 从小到大排序。
2. 排序后，最小值必然对应两个**相邻**元素，设其为 $\textit{nums}[i-1]$ 和 $\textit{nums}[i]$。
3. 把 $\le \textit{nums}[i-1]$ 的数分到第一个数组中，把 $\ge \textit{nums}[i]$ 的数分到第二个数组中。

例如示例 2 可以分成 $[1]$ 和 $[10,100]$。

所以排序后，答案为 $\textit{nums}[i]-\textit{nums}[i-1]$ 的最小值。

```py [sol-Python3]
class Solution:
    def findValueOfPartition(self, nums: List[int]) -> int:
        nums.sort()
        return min(y - x for x, y in pairwise(nums))
```

```java [sol-Java]
class Solution {
    public int findValueOfPartition(int[] nums) {
        Arrays.sort(nums);
        int ans = Integer.MAX_VALUE;
        for (int i = 1; i < nums.length; i++) {
            ans = Math.min(ans, nums[i] - nums[i - 1]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findValueOfPartition(vector<int>& nums) {
        ranges::sort(nums);
        int ans = INT_MAX;
        for (int i = 1; i < nums.size(); i++) {
            ans = min(ans, nums[i] - nums[i - 1]);
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int findValueOfPartition(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int ans = INT_MAX;
    for (int i = 1; i < numsSize; i++) {
        ans = MIN(ans, nums[i] - nums[i - 1]);
    }
    return ans;
}
```

```go [sol-Go]
func findValueOfPartition(nums []int) int {
	slices.Sort(nums)
	ans := math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}
```

```js [sol-JavaScript]
var findValueOfPartition = function(nums) {
    nums.sort((a, b) => a - b);
    let ans = Infinity;
    for (let i = 1; i < nums.length; i++) {
        ans = Math.min(ans, nums[i] - nums[i - 1]);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_value_of_partition(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        nums.windows(2).map(|w| w[1] - w[0]).min().unwrap()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

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
