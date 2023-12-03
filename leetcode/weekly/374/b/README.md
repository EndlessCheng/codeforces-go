[本题视频讲解](https://www.bilibili.com/video/BV1og4y1Z7SZ/)

为方便描述，把 $0$ 也算作可以得到的数。

假设现在得到了 $[0,s-1]$ 内的所有整数，如果此时新发现了一个整数 $x$，那么把 $x$ 加到已得到的数字中，就得到了 $[x,s+x-1]$ 内的所有整数。

分类讨论：

- 如果 $x \le s$，那么合并这两个区间，我们可以得到 $[0,s+x-1]$ 内的所有整数。
- 如果 $x > s$，这意味着我们无法得到 $s$，那么就一定要把 $s$ 加到数组中（加一个比 $s$ 还小的数字就没法得到更大的数，不够贪），这样就可以得到 $[s,2s-1]$ 内的所有整数，再与 $[0,s-1]$ 合并，可以得到 $[0,2s-1]$ 内的所有整数。然后再重新考虑 $x$ 和 $s$ 的大小关系，继续分类讨论。

把 $\textit{coins}$ 排序，从小到大考虑 $x=\textit{coins}[i]$。按照上述分类讨论来看是否要添加数字。

```py [sol-Python3]
class Solution:
    def minimumAddedCoins(self, coins: List[int], target: int) -> int:
        coins.sort()
        ans, s, i = 0, 1, 0
        while s <= target:
            if i < len(coins) and coins[i] <= s:
                s += coins[i]
                i += 1
            else:
                s *= 2  # 必须添加 s
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumAddedCoins(int[] coins, int target) {
        Arrays.sort(coins);
        int ans = 0, s = 1, i = 0;
        while (s <= target) {
            if (i < coins.length && coins[i] <= s) {
                s += coins[i];
                i++;
            } else {
                s *= 2; // 必须添加 s
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
    int minimumAddedCoins(vector<int> &coins, int target) {
        sort(coins.begin(), coins.end());
        int ans = 0, s = 1, i = 0;
        while (s <= target) {
            if (i < coins.size() && coins[i] <= s) {
                s += coins[i];
                i++;
            } else {
                s *= 2; // 必须添加 s
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumAddedCoins(coins []int, target int) (ans int) {
	slices.Sort(coins)
	s, i := 1, 0
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s *= 2 // 必须添加 s
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + \log \textit{target})$，其中 $n$ 为 $\textit{coins}$ 的长度。$s$ 至多翻倍 $\mathcal{O}(\log \textit{target})$ 次。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。
