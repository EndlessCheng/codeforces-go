下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

## 前置知识：同向双指针

详见[【基础算法精讲】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

## 思路

把相同元素分组，相同元素的下标记录到哈希表（或者数组）$\textit{pos}$ 中。

遍历 $\textit{pos}$ 中的每个下标列表 $\textit{ps}$，用双指针处理：

如果等值子数组的元素下标是从 $\textit{ps}[\textit{left}]$ 到 $\textit{ps}[\textit{right}]$，那么子数组的长度为

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}] + 1
$$

其中无需删除的元素个数为

$$
\textit{right} - \textit{left} + 1
$$

那么需要删除的元素个数为

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}] - (\textit{right} - \textit{left})
$$

如果上式大于 $k$，则需要移动左指针。

满足条件时，用 $\textit{right}-\textit{left}+1$ 更新答案的最大值。

代码实现时，可以在哈希表中记录 $\textit{ps}[i]-i$，这样删除的元素个数就是

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}]
$$

```py [sol-Python3]
class Solution:
    def longestEqualSubarray(self, nums: List[int], k: int) -> int:
        pos = [[] for _ in range(len(nums) + 1)]
        for i, x in enumerate(nums):
            pos[x].append(i - len(pos[x]))
        ans = 0
        for ps in pos:
            left = 0
            for right, p in enumerate(ps):
                while p - ps[left] > k:  # 要删除的数太多了
                    left += 1
                ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestEqualSubarray(List<Integer> nums, int k) {
        int n = nums.size(), ans = 0;
        List<Integer>[] pos = new ArrayList[n + 1];
        Arrays.setAll(pos, e -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            int x = nums.get(i);
            pos[x].add(i - pos[x].size());
        }
        for (var ps : pos) {
            int left = 0;
            for (int right = 0; right < ps.size(); right++) {
                while (ps.get(right) - ps.get(left) > k) // 要删除的数太多了
                    left++;
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
        int n = nums.size(), ans = 0;
        vector<vector<int>> pos(n + 1);
        for (int i = 0; i < n; i++)
            pos[nums[i]].push_back(i - pos[nums[i]].size());
        for (auto &ps: pos) {
            int left = 0;
            for (int right = 0; right < ps.size(); right++) {
                while (ps[right] - ps[left] > k) // 要删除的数太多了
                    left++;
                ans = max(ans, right - left + 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i-len(pos[x]))
	}
	for _, ps := range pos {
		left := 0
		for right, p := range ps {
			for p-ps[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

把「删除 $k$ 个数」改成「修改 $k$ 个数」要怎么做？
