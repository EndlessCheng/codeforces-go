[视频讲解](https://www.bilibili.com/video/BV1Sm411U7cR/) 第四题，额外讲了另一种做法，以及如何分类思考子序列 DP。

## 提示 1

把数组排序。

为什么？设我们选的元素在排序后为 $b$，那么有 $b[i] + 1 = b[i+1]$，这意味着 $b$ 中元素在操作前，必然有 $b[i] \le b[i+1]$。反证：如果操作前 $b[i] > b[i+1]$，那么操作后 $b[i]$ 至多和 $b[i+1]$ 相等，不会出现 $b[i+1]$ 比 $b[i]$ 多 $1$ 的情况。

所以可以排序。

## 提示 2

排序后，我们选的是 $\textit{nums}$ 中的一个**子序列**。

定义 $f[x]$ 表示子序列的最后一个数是 $x$ 时，子序列的最大长度。

从左到右遍历数组 $x = \textit{nums}[i]$：

- 如果操作，那么 $x+1$ 可以接在末尾为 $x$ 的子序列后面，即 $f[x+1] = f[x] + 1$。
- 如果不操作，那么 $x$ 可以接在末尾为 $x-1$ 的子序列后面，即 $f[x] = f[x-1] + 1$。

比如 $\textit{nums} = [1,2,2]$：

- 遍历到 $\textit{nums}[0]=1$ 时，$f[2]=1,\ f[1]=1$。
- 遍历到 $\textit{nums}[1]=2$ 时，$f[3]=f[2]+1=2,\ f[2]=f[1]+1=2$。注意要先计算 $f[x+1]$ 再计算 $f[x]$（不然这里会算出 $f[3]=3$）。此时 $f[1]$ 还是 $1$。
- 遍历到 $\textit{nums}[2]=2$ 时，$f[3]=f[2]+1=3,\ f[2]=f[1]+1=2$。此时 $f[1]$ 还是 $1$。

最后返回 $f[x]$ 的最大值。

```py [sol-Python3]
class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        f = defaultdict(int)
        for x in nums:
            f[x + 1] = f[x] + 1
            f[x] = f[x - 1] + 1
        return max(f.values())
```

```java [sol-Java]
class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, Integer> f = new HashMap<>();
        for (int x : nums) {
            f.put(x + 1, f.getOrDefault(x, 0) + 1);
            f.put(x, f.getOrDefault(x - 1, 0) + 1);
        }
        int ans = 0;
        for (int res : f.values()) {
            ans = Math.max(ans, res);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSelectedElements(vector<int> &nums) {
        ranges::sort(nums);
        unordered_map<int, int> f;
        for (int x : nums) {
            f[x + 1] = f[x] + 1;
            f[x] = f[x - 1] + 1;
        }
        int ans = 0;
        for (auto &[_, res] : f) {
            ans = max(ans, res);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSelectedElements(nums []int) (ans int) {
	slices.Sort(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
	}
	for _, res := range f {
		ans = max(ans, res)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
