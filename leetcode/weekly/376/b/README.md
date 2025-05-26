注意：本题的子数组是**子集**的意思，不是连续子数组。

对于 $\textit{nums}$ 的最小值 $x$ 来说，应该把 $x$ 与次小值 $y$ 和第三小值 $z$ 放在同一组。如果不这样做，其他某些组就不得不与更小的 $y,z$ 在一起了，这可能无法满足要求。 

因此做法是，把数组排序后，三个三个地切分。（注意题目保证数组长度是 $3$ 的倍数）

如果同一组中的最大最小之差大于 $k$，无法满足要求，返回空列表。

```py [sol-Python3]
class Solution:
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        nums.sort()
        ans = []
        for i in range(0, len(nums), 3):
            if nums[i + 2] - nums[i] > k:
                return []
            ans.append(nums[i: i + 3])
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] divideArray(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        int[][] ans = new int[n / 3][3];
        for (int i = 2; i < n; i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return new int[][]{};
            }
            ans[i / 3] = new int[]{nums[i - 2], nums[i - 1], nums[i]};
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> divideArray(vector<int>& nums, int k) {
        ranges::sort(nums);
        vector<vector<int>> ans;
        for (int i = 2; i < nums.size(); i += 3) {
            if (nums[i] - nums[i - 2] > k) {
                return {};
            }
            ans.push_back({nums[i - 2], nums[i - 1], nums[i]});
        }
        return ans;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int** divideArray(int* nums, int numsSize, int k, int* returnSize, int** returnColSize) {
    qsort(nums, numsSize, sizeof(int), cmp);

    #define SIZE 3
    *returnSize = 0;
    for (int i = 0; i < numsSize; i += SIZE) {
        if (nums[i + SIZE - 1] - nums[i] > k) {
            *returnColSize = NULL;
            return NULL;
        }
    }

    int** ans = malloc(numsSize / SIZE * sizeof(int*));
    *returnColSize = malloc(numsSize / SIZE * sizeof(int));
    for (int i = 0; i < numsSize; i += SIZE) {
        ans[*returnSize] = malloc(SIZE * sizeof(int));
        (*returnColSize)[*returnSize] = SIZE;
        // 复制从 nums[i] 开始的 SIZE 个数
        memcpy(ans[(*returnSize)++], &nums[i], SIZE * sizeof(int));
    }
    return ans;
}
```

```go [sol-Go]
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for a := range slices.Chunk(nums, 3) {
		if a[2]-a[0] > k {
			return nil
		}
		ans = append(ans, a)
	}
	return
}
```

```js [sol-JavaScript]
var divideArray = function(nums, k) {
    nums.sort((a, b) => a - b);
    const ans = [];
    for (let i = 0; i < nums.length; i += 3) {
        if (nums[i + 2] - nums[i] > k) {
            return [];
        }
        ans.push(nums.slice(i, i + 3));
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn divide_array(mut nums: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        let mut ans = vec![];
        for a in nums.chunks(3) {
            if a[2] - a[0] > k {
                return vec![];
            }
            ans.push(a.to_vec());
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，返回值不计入。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
