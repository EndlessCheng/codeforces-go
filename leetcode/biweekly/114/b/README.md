下午两点[【b站@灵茶山艾府】](https://b23.tv/JMcHRRp)直播讲题，欢迎关注！

---

先统计每个元素的出现次数。

考虑出现了 $c$ 次的数：

- 如果 $c=1$，无法操作，返回 $-1$。
- 如果 $c$ 是 $3$ 的倍数，那么可以用 $\dfrac{c}{3}$ 次操作删除。
- 如果 $c$ 模 $3$ 为 $1$，那么 $c=(c-4)+4$，其中 $c-4$ 是 $3$ 的倍数，剩余的 $4$ 可以用两次操作完成。
- 如果 $c$ 模 $3$ 为 $2$，那么 $c=(c-2)+2$，其中 $c-2$ 是 $3$ 的倍数，剩余的 $2$ 可以用一次操作完成。
- 总的来说，都需要 $\left\lceil\dfrac{c}{3}\right\rceil=\left\lfloor\dfrac{c+2}{3}\right\rfloor$ 次操作完成。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        if 1 in cnt.values():
            return -1
        return sum((c + 2) // 3 for c in cnt.values())
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        var cnt = new HashMap<Integer, Integer>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum);
        }
        int ans = 0;
        for (int c : cnt.values()) {
            if (c == 1) {
                return -1;
            }
            ans += (c + 2) / 3;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }
        int ans = 0;
        for (auto &[_, c] : cnt) {
            if (c == 1) {
                return -1;
            }
            ans += (c + 2) / 3;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range cnt {
		if c == 1 {
			return -1
		}
		ans += (c + 2) / 3
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相同题目

- [2244. 完成所有任务需要的最少轮数](https://leetcode.cn/problems/minimum-rounds-to-complete-all-tasks/)
