## 观察

为方便描述，把 $0$ 也算作可以得到的数。

假设现在得到了区间 $[0,s-1]$ 中的所有整数，如果此时遍历到整数 $x=\textit{coins}[i]$，那么把 $[0,s-1]$ 中的每个整数都增加 $x$，我们就得到了区间 $[x,s+x-1]$ 中的所有整数。

## 思路

把 $\textit{coins}$ 从小到大排序，遍历 $x=\textit{coins}[i]$。分类讨论，看是否要添加数字：

- 如果 $x \le s$，那么合并 $[0,s-1]$ 和 $[x,s+x-1]$ 这两个区间，我们可以得到 $[0,s+x-1]$ 中的所有整数。
- 如果 $x > s$，或者遍历完了 $\textit{coins}$ 数组，这意味着我们无法得到 $s$，那么就一定要把 $s$ 加到数组中（加一个比 $s$ 还小的数字就没法得到更大的数，不够贪），这样就可以得到了 $[s,2s-1]$ 中的所有整数，再与 $[0,s-1]$ 合并，可以得到 $[0,2s-1]$ 中的所有整数。然后再考虑 $x$ 和 $2s$ 的大小关系，继续分类讨论。

当 $s > \textit{target}$ 时，我们就得到了 $[1,target]$ 中的所有整数，退出循环。

附：[视频讲解](https://www.bilibili.com/video/BV1og4y1Z7SZ/)

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
                s += coins[i++];
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
        ranges::sort(coins);
        int ans = 0, s = 1, i = 0;
        while (s <= target) {
            if (i < coins.size() && coins[i] <= s) {
                s += coins[i++];
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

```js [sol-JavaScript]
var minimumAddedCoins = function(coins, target) {
    coins.sort((a, b) => a - b);
    let ans = 0, s = 1, i = 0;
    while (s <= target) {
        if (i < coins.length && coins[i] <= s) {
            s += coins[i++];
        } else {
            s *= 2; // 必须添加 s
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_added_coins(mut coins: Vec<i32>, target: i32) -> i32 {
        coins.sort_unstable();
        let mut ans = 0;
        let mut s = 1;
        let mut i = 0;
        while s <= target {
            if i < coins.len() && coins[i] <= s {
                s += coins[i];
                i += 1;
            } else {
                s *= 2; // 必须添加 s
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + \log \textit{target})$，其中 $n$ 为 $\textit{coins}$ 的长度。$s$ 至多翻倍 $\mathcal{O}(\log \textit{target})$ 次。瓶颈主要在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 相似题目

- [330. 按要求补齐数组](https://leetcode.cn/problems/patching-array/)
- [1798. 你能构造出连续值的最大数目](https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
