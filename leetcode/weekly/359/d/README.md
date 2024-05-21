### 前置知识：滑动窗口

请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

### 思路

把相同元素分组，相同元素的下标记录到哈希表（或者数组）$\textit{posLists}$ 中。

例如示例 1，元素 $3$ 在 $\textit{nums}$ 中的下标有 $1,3,5$，那么 $\textit{posLists}[3] = [1,3,5]$。

遍历 $\textit{posLists}$ 中的每个下标列表 $\textit{pos}$，例如遍历 $\textit{pos}=[1,3,5]$。

请记住，$\textit{pos}$ 中保存的是下标，这些下标在 $\textit{nums}$ 中的对应元素都相同。

然后用**滑动窗口**计算。设窗口左右端点为 $\textit{left}$ 和 $\textit{right}$。

假设 $\textit{nums}$ 的等值子数组的元素下标从 $\textit{pos}[\textit{left}]$ 到 $\textit{pos}[\textit{right}]$，那么在删除前，子数组的长度为

$$
\textit{pos}[\textit{right}] - \textit{pos}[\textit{left}] + 1
$$

这个子数组有

$$
\textit{right} - \textit{left} + 1
$$

个数都是相同的，无需删除，其余元素都需要删除，那么需要删除的元素个数就是

$$
\textit{pos}[\textit{right}] - \textit{pos}[\textit{left}] - (\textit{right} - \textit{left})
$$

如果上式大于 $k$，说明要删除的数太多了，那么移动左指针 $\textit{left}$，直到上式小于等于 $k$，此时用 $\textit{right}-\textit{left}+1$ 更新答案的最大值。

代码实现时，为简化上式，$\textit{pos}$ 实际保存的是 $\textit{pos}[i]-i$，也就是把上面的每个 $\textit{pos}[i]$ 都减去其在 $\textit{pos}$ 中的下标 $i$，于是需要删除的元素个数简化为

$$
\textit{pos}[\textit{right}] - \textit{pos}[\textit{left}]
$$

```py [sol-Python3]
class Solution:
    def longestEqualSubarray(self, nums: List[int], k: int) -> int:
        pos_lists = defaultdict(list)
        for i, x in enumerate(nums):
            pos_lists[x].append(i - len(pos_lists[x]))

        ans = 0
        for pos in pos_lists.values():
            if len(pos) <= ans:
                continue  # 无法让 ans 变得更大
            left = 0
            for right, p in enumerate(pos):
                while p - pos[left] > k:  # 要删除的数太多了
                    left += 1
                ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestEqualSubarray(List<Integer> nums, int k) {
        int n = nums.size();
        List<Integer>[] posLists = new ArrayList[n + 1];
        Arrays.setAll(posLists, i -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            int x = nums.get(i);
            posLists[x].add(i - posLists[x].size());
        }

        int ans = 0;
        for (List<Integer> pos : posLists) {
            if (pos.size() <= ans) {
                continue; // 无法让 ans 变得更大
            }
            int left = 0;
            for (int right = 0; right < pos.size(); right++) {
                while (pos.get(right) - pos.get(left) > k) { // 要删除的数太多了
                    left++;
                }
                ans = Math.max(ans, right - left + 1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestEqualSubarray(vector<int> &nums, int k) {
        int n = nums.size();
        vector<vector<int>> pos_lists(n + 1);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            pos_lists[x].push_back(i - pos_lists[x].size());
        }

        int ans = 0;
        for (auto& pos : pos_lists) {
            int left = 0;
            for (int right = 0; right < pos.size(); right++) {
                while (pos[right] - pos[left] > k) { // 要删除的数太多了
                    left++;
                }
                ans = max(ans, right - left + 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestEqualSubarray(nums []int, k int) (ans int) {
	posLists := make([][]int, len(nums)+1)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i-len(posLists[x]))
	}

	for _, pos := range posLists {
		if len(pos) <= ans {
			continue // 无法让 ans 变得更大
		}
		left := 0
		for right, p := range pos {
			for p-pos[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}
```

```js [sol-JavaScript]
const longestEqualSubarray = function(nums, k) {
    const n = nums.length;
    const posLists = Array.from({length: n + 1}, () => []);
    for (let i = 0; i < n; i++) {
        const x = nums[i];
        posLists[x].push(i - posLists[x].length);
    }

    let ans = 0;
    for (const pos of posLists) {
        if (pos.length <= ans) {
            continue; // 无法让 ans 变得更大
        }
        let left = 0;
        for (let right = 0; right < pos.length; right++) {
            while (pos[right] - pos[left] > k) { // 要删除的数太多了
                left++;
            }
            ans = Math.max(ans, right - left + 1);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_equal_subarray(nums: Vec<i32>, k: i32) -> i32 {
        let mut pos_lists = vec![vec![]; nums.len() + 1];
        for (i, &x) in nums.iter().enumerate() {
            let mut pos = &mut pos_lists[x as usize];
            pos.push(i - pos.len());
        }

        let mut ans = 0;
        for pos in pos_lists {
            let mut left = 0;
            for (right, &p) in pos.iter().enumerate() {
                while p - pos[left] > k as usize { // 要删除的数太多了
                    left += 1;
                }
                ans = ans.max(right - left + 1);
            }
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 思考题

把「删除最多 $k$ 个数」改成「修改最多 $k$ 个数」，要怎么做？

欢迎在评论区分享你的思路/代码。

### 分类题单

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
