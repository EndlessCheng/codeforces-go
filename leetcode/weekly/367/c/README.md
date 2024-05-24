不妨假设 $i$ 在左，$j$ 在右，即 $i\le j - \textit{indexDifference}$。

枚举 $j$，寻找左边的 $i$。要想满足 $|\textit{nums}[i]-\textit{nums}[j]|\ge \textit{valueDifference}$，要找的 $\textit{nums}[i]$ 应当尽量大或者尽量小，这样差的绝对值才能尽量大。

类似 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)，我们可以在枚举 $j$ 的同时，维护 $\textit{nums}[0]$ 到 $\textit{nums}[j - \textit{indexDifference}]$ 中的最大值 $\textit{mx}$ 和最小值 $\textit{mn}$。

例如 $\textit{indexDifference}=2$，从 $j=2$ 开始遍历数组：

- 遍历到 $\textit{nums}[2]$，用 $\textit{nums}[0]$ 更新 $\textit{mx}$ 和 $\textit{mn}$。
- 遍历到 $\textit{nums}[3]$，用 $\textit{nums}[1]$ 更新 $\textit{mx}$ 和 $\textit{mn}$。
- 遍历到 $\textit{nums}[4]$，用 $\textit{nums}[2]$ 更新 $\textit{mx}$ 和 $\textit{mn}$。
- 依此类推。

这个过程可以保证 $\textit{mx}$ 和 $\textit{mn}$ 在数组中的下标 $i$ 满足 $i\le j - \textit{indexDifference}$，即题目的第一个要求。

对于题目的第二个要求，可以转换成如下两个不等式：

- $\textit{mx} -\textit{nums}[j] \ge \textit{valueDifference}$
- $\textit{nums}[j] - mn \ge \textit{valueDifference}$

满足其一即可返回答案。不用算绝对值的原因见下面的答疑。

由于要输出 $\textit{mx}$ 或者 $\textit{mn}$ 在数组中的下标，我们可以记录 $\textit{mx}$ 在数组中的下标 $\textit{maxIdx}$，以及 $\textit{mn}$ 在数组中的下标 $\textit{minIdx}$。由于知道下标就能知道元素值，所以只需记录下标，无需记录 $\textit{mx}$ 和 $\textit{mn}$。

### 答疑

**问**：为什么不用算绝对值？万一 $\textit{mx} < \textit{nums}[j]$，并且 $|\textit{mx} - \textit{nums}[j]| = \textit{nums}[j] - \textit{mx} \ge \textit{valueDifference}$，不就错过答案了吗？

**答**：在上述条件成立的前提下，由于 $\textit{mn} \le \textit{mx}$，得

$$
\textit{nums}[j] - \textit{mn} \ge \textit{nums}[j] - \textit{mx} \ge \textit{valueDifference}
$$ 

所以此时 $\textit{mn}$ 是满足要求的，不会错过答案。

```py [sol-Python3]
class Solution:
    def findIndices(self, nums: List[int], indexDifference: int, valueDifference: int) -> List[int]:
        max_idx = min_idx = 0
        for j in range(indexDifference, len(nums)):
            i = j - indexDifference
            if nums[i] > nums[max_idx]:
                max_idx = i
            elif nums[i] < nums[min_idx]:
                min_idx = i
            if nums[max_idx] - nums[j] >= valueDifference:
                return [max_idx, j]
            if nums[j] - nums[min_idx] >= valueDifference:
                return [min_idx, j]
        return [-1, -1]
```

```java [sol-Java]
class Solution {
    public int[] findIndices(int[] nums, int indexDifference, int valueDifference) {
        int maxIdx = 0;
        int minIdx = 0;
        for (int j = indexDifference; j < nums.length; j++) {
            int i = j - indexDifference;
            if (nums[i] > nums[maxIdx]) {
                maxIdx = i;
            } else if (nums[i] < nums[minIdx]) {
                minIdx = i;
            }
            if (nums[maxIdx] - nums[j] >= valueDifference) {
                return new int[]{maxIdx, j};
            }
            if (nums[j] - nums[minIdx] >= valueDifference) {
                return new int[]{minIdx, j};
            }
        }
        return new int[]{-1, -1};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findIndices(vector<int> &nums, int indexDifference, int valueDifference) {
        int max_idx = 0, min_idx = 0;
        for (int j = indexDifference; j < nums.size(); j++) {
            int i = j - indexDifference;
            if (nums[i] > nums[max_idx]) {
                max_idx = i;
            } else if (nums[i] < nums[min_idx]) {
                min_idx = i;
            }
            if (nums[max_idx] - nums[j] >= valueDifference) {
                return {max_idx, j};
            }
            if (nums[j] - nums[min_idx] >= valueDifference) {
                return {min_idx, j};
            }
        }
        return {-1, -1};
    }
};
```

```c [sol-C]
int* findIndices(int* nums, int numsSize, int indexDifference, int valueDifference, int* returnSize) {
    int max_idx = 0, min_idx = 0;
    int* ans = malloc(2 * sizeof(int));
    *returnSize = 2;
    for (int j = indexDifference; j < numsSize; j++) {
        int i = j - indexDifference;
        if (nums[i] > nums[max_idx]) {
            max_idx = i;
        } else if (nums[i] < nums[min_idx]) {
            min_idx = i;
        }
        if (nums[max_idx] - nums[j] >= valueDifference) {
            ans[0] = max_idx;
            ans[1] = j;
            return ans;
        }
        if (nums[j] - nums[min_idx] >= valueDifference) {
            ans[0] = min_idx;
            ans[1] = j;
            return ans;
        }
    }
    ans[0] = -1;
    ans[1] = -1;
    return ans;
}
```

```go [sol-Go]
func findIndices(nums []int, indexDifference, valueDifference int) []int {
	maxIdx, minIdx := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		} else if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[maxIdx]-nums[j] >= valueDifference {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDifference {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}
```

```js [sol-JavaScript]
var findIndices = function(nums, indexDifference, valueDifference) {
    let maxIdx = 0, minIdx = 0;
    for (let j = indexDifference; j < nums.length; j++) {
        const i = j - indexDifference;
        if (nums[i] > nums[maxIdx]) {
            maxIdx = i;
        } else if (nums[i] < nums[minIdx]) {
            minIdx = i;
        }
        if (nums[maxIdx] - nums[j] >= valueDifference) {
            return [maxIdx, j];
        }
        if (nums[j] - nums[minIdx] >= valueDifference) {
            return [minIdx, j];
        }
    }
    return [-1, -1];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_indices(nums: Vec<i32>, index_difference: i32, value_difference: i32) -> Vec<i32> {
        let mut max_idx = 0;
        let mut min_idx = 0;
        for j in index_difference as usize..nums.len() {
            let i = j - index_difference as usize;
            if nums[i] > nums[max_idx] {
                max_idx = i;
            } else if nums[i] < nums[min_idx] {
                min_idx = i;
            }
            if nums[max_idx] - nums[j] >= value_difference {
                return vec![max_idx as i32, j as i32];
            }
            if nums[j] - nums[min_idx] >= value_difference {
                return vec![min_idx as i32, j as i32];
            }
        }
        vec![-1, -1]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n-\textit{indexDifference})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
