在选择学生人数固定的时候，选择方案是否唯一呢？

假设恰好选 $k$ 个学生，那么：

- 所有 $\textit{nums}[i] < k$ 的学生都要选；
- 所有 $\textit{nums}[i] > k$ 的学生都不能选；
- 不能出现 $\textit{nums}[i] = k$ 的情况，因为每个学生只有选或不选两种可能。

这意味着**在选择学生人数固定的时候，选择方案是唯一的**。把 $\textit{nums}$ 从小到大排序后，唯一性可以更明显地看出来：

- 以 $k$ 为分界线，左边的都要选，右边的都不能选。
  
排序后：

- 如果选了 $\textit{nums}[i]$，那么比 $\textit{nums}[i]$ 更小的学生也要选。
- 如果不选 $\textit{nums}[i]$，那么比 $\textit{nums}[i]$ 更大的学生也不选。

具体地，如果选 $\textit{nums}[i-1]$ 而不选 $\textit{nums}[i]$，由于数组已排序，我们必须要选下标为 $0,1,2,\cdots,i-1$ 的学生，一共 $i$ 个，而下标 $\ge i$ 的学生都不能选，所以需要满足

$$
\textit{nums}[i-1] < i < \textit{nums}[i]
$$

枚举 $i=1,2,\cdots,n-1$（枚举分界线的位置），如果上式成立，就意味着我们可以选 $i$ 个学生，算作一种方案。

特殊情况：

- 如果 $\textit{nums}[0] > 0$，那么可以一个学生都不选。
- 如果 $\textit{nums}[n-1] < n$，那么可以所有学生都选。由于数据范围保证 $\textit{nums}[i]<n$，所以这种方案一定存在。

见 [视频讲解](https://www.bilibili.com/video/BV1Lm4y1N7mf/) 第二题。

```py [sol-Python3]
class Solution:
    def countWays(self, nums: List[int]) -> int:
        nums.sort()
        ans = nums[0] > 0  # 一个学生都不选
        for i, (x, y) in enumerate(pairwise(nums), 1):
            if x < i < y:
                ans += 1
        return ans + 1  # 一定可以都选
```

```java [sol-Java]
class Solution {
    public int countWays(List<Integer> nums) {
        int[] a = nums.stream().mapToInt(i -> i).toArray();
        Arrays.sort(a);
        int ans = a[0] > 0 ? 1 : 0; // 一个学生都不选
        for (int i = 1; i < a.length; i++) {
            if (a[i - 1] < i && i < a[i]) {
                ans++;
            }
        }
        return ans + 1; // 一定可以都选
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWays(vector<int>& nums) {
        ranges::sort(nums);
        int ans = nums[0] > 0; // 一个学生都不选
        for (int i = 1; i < nums.size(); i++) {
            ans += nums[i - 1] < i && i < nums[i];
        }
        return ans + 1; // 一定可以都选
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int countWays(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int ans = nums[0] > 0; // 一个学生都不选
    for (int i = 1; i < numsSize; i++) {
        ans += nums[i - 1] < i && i < nums[i];
    }
    return ans + 1; // 一定可以都选
}
```

```go [sol-Go]
func countWays(nums []int) (ans int) {
    slices.Sort(nums)
    if nums[0] > 0 { // 一个学生都不选
        ans = 1
    }
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < i && i < nums[i] {
            ans++
        }
    }
    return ans + 1 // 一定可以都选
}
```

```js [sol-JavaScript]
var countWays = function(nums) {
    nums.sort((a, b) => a - b);
    let ans = nums[0] > 0 ? 1 : 0; // 一个学生都不选
    for (let i = 1; i < nums.length; i++) {
        if (nums[i - 1] < i && i < nums[i]) {
            ans++;
        }
    }
    return ans + 1; // 一定可以都选
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_ways(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        let mut ans = (nums[0] > 0) as i32; // 一个学生都不选
        for i in 1..nums.len() {
            let k = i as i32;
            if nums[i - 1] < k && k < nums[i] {
                ans += 1;
            }
        }
        ans + 1 // 一定可以都选
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

注：如果采用计数排序，可以做到 $\mathcal{O}(n)$ 的时间复杂度。

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
