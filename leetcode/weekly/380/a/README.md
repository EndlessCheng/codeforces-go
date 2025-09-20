遍历 $\textit{nums}$，同时用哈希表统计每个元素的出现次数，并维护出现次数的最大值 $\textit{maxCnt}$：

- 如果出现次数 $c > \textit{maxCnt}$，那么更新 $\textit{maxCnt}=c$，答案 $\textit{ans} = c$。
- 如果出现次数 $c = \textit{maxCnt}$，那么答案增加 $c$。

```py [sol-Python3]
class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        cnt = defaultdict(int)
        ans = max_cnt = 0
        for x in nums:
            cnt[x] += 1
            c = cnt[x]
            if c > max_cnt:
                ans = max_cnt = c
            elif c == max_cnt:
                ans += c
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequencyElements(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>(); // 更快的写法见【Java 数组】
        int maxCnt = 0;
        int ans = 0;
        for (int x : nums) {
            int c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
            if (c > maxCnt) {
                ans = maxCnt = c;
            } else if (c == maxCnt) {
                ans += c;
            }
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int maxFrequencyElements(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        
        int[] cnt = new int[mx + 1];
        int maxCnt = 0;
        int ans = 0;
        for (int x : nums) {
            int c = ++cnt[x];
            if (c > maxCnt) {
                ans = maxCnt = c;
            } else if (c == maxCnt) {
                ans += c;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequencyElements(vector<int>& nums) {
        unordered_map<int, int> cnt;
        int ans = 0, max_cnt = 0;
        for (int x : nums) {
            int c = ++cnt[x];
            if (c > max_cnt) {
                ans = max_cnt = c;
            } else if (c == max_cnt) {
                ans += c;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maxFrequencyElements(int* nums, int numsSize) {
    int mx = 0; // 直接初始化 mx = 100 可以做到一次遍历
    for (int i = 0; i < numsSize; i++) {
        mx = MAX(mx, nums[i]);
    }

    int* cnt = calloc(mx + 1, sizeof(int));
    int max_cnt = 0;
    int ans = 0;

    for (int i = 0; i < numsSize; i++) {
        int c = ++cnt[nums[i]];
        if (c > max_cnt) {
            ans = c;
            max_cnt = c;
        } else if (c == max_cnt) {
            ans += c;
        }
    }

    free(cnt);
    return ans;
}
```

```go [sol-Go]
func maxFrequencyElements(nums []int) (ans int) {
	cnt := map[int]int{}
	maxCnt := 0
	for _, x := range nums {
		cnt[x]++
		c := cnt[x]
		if c > maxCnt {
			maxCnt = c
			ans = c
		} else if c == maxCnt {
			ans += c
		}
	}
	return
}
```

```js [sol-JavaScript]
var maxFrequencyElements = function(nums) {
    const cnt = new Map();
    let ans = 0, maxCnt = 0;
    for (const x of nums) {
        const c = (cnt.get(x) ?? 0) + 1;
        cnt.set(x, c);
        if (c > maxCnt) {
            ans = maxCnt = c;
        } else if (c === maxCnt) {
            ans += c;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn max_frequency_elements(nums: Vec<i32>) -> i32 {
        let mut cnt = HashMap::new();
        let mut max_cnt = 0;
        let mut ans = 0;
        for x in nums {
            let e = cnt.entry(x).or_insert(0);
            *e += 1;
            let c = *e;
            if c > max_cnt {
                max_cnt = c;
                ans = c;
            } else if c == max_cnt {
                ans += c;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
