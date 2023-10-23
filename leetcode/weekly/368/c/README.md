统计每个数字的出现次数，记在哈希表 $\textit{cnt}$ 中。

假设可以分成大小为 $k$ 和 $k+1$ 的组，现在需要算出每个 $\textit{cnt}[x]$ 最少可以分成多少组。

举例说明，假设 $\textit{\textit{cnt}}[x]=32$，$k=10$，那么 $32=10+10+10+2$，多出的 $2$ 可以分成两个 $1$，加到两个 $10$ 中，从而得到 $11,11,10$ 这三组。

但如果 $\textit{\textit{cnt}}[x]=34$，那么 $34=10+10+10+4$，多出的 $4$ 无法加到另外三个 $10$ 中。

设 $q=\left\lfloor\dfrac{\textit{cnt}[x]}{k}\right\rfloor$，$r = \textit{cnt}[x] \bmod k$。

如果 $q < r$ 则无法分成 $k$ 和 $k+1$ 组，否则一定可以分组。

**在可以分组的前提下**，分出的 $k+1$ 越多，组数就越少，所以最少可以分出

$$
\left\lceil\dfrac{\textit{cnt}[x]}{k+1}\right\rceil
$$

组。累加组数即为分组个数。

例如 $\textit{cnt}[x] = 9$，如果先分出尽量多的 $k=2$，再分出 $k+1=3$，那么分组方案就是 $9=2+2+2+3$，但是先按照 $k+1=3$ 分，则有 $9=3+3+3$，可以分出更少的组。

从 $\min(\textit{cnt}[x])$ 开始倒着枚举 $k$，只要可以分，就立刻返回答案。

```py [sol-Python3]
class Solution:
    def minGroupsForValidAssignment(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        for k in range(min(cnt.values()), 0, -1):
            ans = 0
            for c in cnt.values():
                q, r = divmod(c, k)
                if q < r:
                    break
                ans += (c + k) // (k + 1)
            else:
                return ans
```

```java [sol-Java]
class Solution {
    public int minGroupsForValidAssignment(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum);
        }
        int k = nums.length;
        for (int c : cnt.values()) {
            k = Math.min(k, c);
        }
        for (; ; k--) {
            int ans = 0;
            for (int c : cnt.values()) {
                if (c / k < c % k) {
                    ans = 0;
                    break;
                }
                ans += (c + k) / (k + 1);
            }
            if (ans > 0) {
                return ans;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minGroupsForValidAssignment(vector<int>& nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }
        int k = min_element(cnt.begin(), cnt.end(), [](const auto& a, const auto& b) {
            return a.second < b.second;
        })->second;
        for (; ; k--) {
            int ans = 0;
            for (auto &[_, c] : cnt) {
                if (c / k < c % k) {
                    ans = 0;
                    break;
                }
                ans += (c + k) / (k + 1);
            }
            if (ans) {
                return ans;
            }
        }
    }
};
```

```go [sol-Go]
func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, c := range cnt {
		k = min(k, c)
	}
	for ; ; k-- {
		ans := 0
		for _, c := range cnt {
			if c/k < c%k {
				ans = 0
				break
			}
			ans += (c + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。设哈希表的大小为 $m$，哈希表中最小的 value 为 $k$，由于所有 value 之和为 $n$，所以 $km\le n$。而循环次数又至多为 $km$，所以时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
