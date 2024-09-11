**核心思路**：计算每个点被覆盖了多少次。统计覆盖次数大于 $0$ 的点，即为答案。

假设一开始有一个全为 $0$ 的数组 $a$，用来保存每个点被覆盖了多少次。

对于示例 1，我们可以把 $a$ 中下标在 $[3,6]$ 的元素都加一，下标在 $[1,5]$ 的元素都加一，下标在 $[4,7]$ 的元素都加一。

然后，统计 $a[i] > 0$ 的个数，即为答案。

如何快速地「把区间内的数都加一」呢？

这可以用**差分数组**实现。见 [原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

```py [sol-Python3]
class Solution:
    def numberOfPoints(self, nums: List[List[int]]) -> int:
        max_end = max(end for _, end in nums)
        diff = [0] * (max_end + 2)  # 注意下面有 end+1
        for start, end in nums:
            diff[start] += 1
            diff[end + 1] -= 1
        return sum(s > 0 for s in accumulate(diff))
```

```java [sol-Java]
class Solution {
    public int numberOfPoints(List<List<Integer>> nums) {
        int maxEnd = 0;
        for (List<Integer> p : nums) {
            maxEnd = Math.max(maxEnd, p.get(1));
        }

        int[] diff = new int[maxEnd + 2]; // 注意下面有 end+1
        for (List<Integer> interval : nums) {
            diff[interval.get(0)]++;
            diff[interval.get(1) + 1]--;
        }

        int ans = 0;
        int s = 0;
        for (int d : diff) {
            s += d;
            if (s > 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfPoints(vector<vector<int>>& nums) {
        int max_end = ranges::max(nums, {}, [](const auto& a) { return a[1]; })[1];

        vector<int> diff(max_end + 2); // 注意下面有 end+1
        for (auto& interval : nums) {
            diff[interval[0]]++;
            diff[interval[1] + 1]--;
        }

        int ans = 0, s = 0;
        for (int d : diff) {
            s += d;
            ans += s > 0;
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int numberOfPoints(int** nums, int numsSize, int* numsColSize) {
    int max_end = 0;
    for (int i = 0; i < numsSize; i++) {
        max_end = MAX(max_end, nums[i][1]);
    }

    int* diff = calloc(max_end + 2, sizeof(int)); // 注意下面有 end+1
    for (int i = 0; i < numsSize; i++) {
        diff[nums[i][0]]++;
        diff[nums[i][1] + 1]--;
    }

    int ans = 0;
    int s = 0;
    for (int i = 1; i <= max_end; i++) {
        s += diff[i];
        ans += s > 0;
    }

    free(diff);
    return ans;
}
```

```go [sol-Go]
func numberOfPoints(nums [][]int) (ans int) {
    maxEnd := 0
    for _, interval := range nums {
        maxEnd = max(maxEnd, interval[1])
    }

    diff := make([]int, maxEnd+2) // 注意下面有 end+1
    for _, interval := range nums {
        diff[interval[0]]++
        diff[interval[1]+1]--
    }

    s := 0
    for _, d := range diff {
        s += d
        if s > 0 {
            ans++
        }
    }
    return
}
```

```js [sol-JavaScript]
var numberOfPoints = function(nums) {
    const maxEnd = Math.max(...nums.map(interval => interval[1]));

    const diff = Array(maxEnd + 2).fill(0); // 注意下面有 end+1
    for (const [start, end] of nums) {
        diff[start]++;
        diff[end + 1]--;
    }

    let ans = 0, s = 0;
    for (const d of diff) {
        s += d;
        if (s > 0) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn number_of_points(nums: Vec<Vec<i32>>) -> i32 {
        let max_end = nums.iter().map(|interval| interval[1]).max().unwrap() as usize;

        let mut diff = vec![0; max_end + 2]; // 注意下面有 end+1
        for interval in nums {
            diff[interval[0] as usize] += 1;
            diff[(interval[1] + 1) as usize] -= 1;
        }

        let mut ans = 0;
        let mut s = 0;
        for d in diff {
            s += d;
            if s > 0 {
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+U)$。其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max\{\textit{end}_i\}$。
- 空间复杂度：$\mathcal{O}(U)$。

更多相似题目，见下面数据结构题单中的「**差分**」。

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
