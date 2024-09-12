根据题意，首先取出 $\textit{nums}$ 中的最小值和次小值，依次把次小值和最小值加到 $\textit{arr}$ 末尾。然后再次取出 $\textit{nums}$ 中的剩余元素的最小值和次小值，依次把次小值和最小值加到 $\textit{arr}$ 末尾。重复该过程，直到 $\textit{nums}$ 为空。

该操作等价于，把 $\textit{nums}$ 从小到大排序，然后从左到右，两个两个一对交换。

注意题目保证 $\textit{nums}$ 的长度是偶数。

```py [sol-Python3]
class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        nums.sort()
        for i in range(1, len(nums), 2):
            nums[i - 1], nums[i] = nums[i], nums[i - 1]
        return nums
```

```py [sol-Python3 切片]
class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        nums.sort()
        nums[::2], nums[1::2] = nums[1::2], nums[::2]
        return nums
```

```java [sol-Java]
class Solution {
    public int[] numberGame(int[] nums) {
        Arrays.sort(nums);
        for (int i = 1; i < nums.length; i += 2) {
            int tmp = nums[i - 1];
            nums[i - 1] = nums[i];
            nums[i] = tmp;
        }
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> numberGame(vector<int>& nums) {
        ranges::sort(nums);
        for (int i = 1; i < nums.size(); i += 2) {
            swap(nums[i - 1], nums[i]);
        }
        return nums;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return (*(int*)a - *(int*)b);
}

int* numberGame(int* nums, int numsSize, int* returnSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    for (int i = 1; i < numsSize; i += 2) {
        int tmp = nums[i - 1];
        nums[i - 1] = nums[i];
        nums[i] = tmp;
    }
    *returnSize = numsSize;
    return nums;
}
```

```go [sol-Go]
func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}
```

```js [sol-JS]
var numberGame = function(nums) {
    nums.sort((a, b) => a - b);
    for (let i = 1; i < nums.length; i += 2) {
        [nums[i - 1], nums[i]] = [nums[i], nums[i - 1]];
    }
    return nums;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn number_game(mut nums: Vec<i32>) -> Vec<i32> {
        nums.sort_unstable();
        for i in (1..nums.len()).step_by(2) {
            nums.swap(i - 1, i);
        }
        nums
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
