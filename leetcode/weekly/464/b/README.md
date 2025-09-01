问题相当于每次从 $\textit{nums}$ 中删除恰好 $k$ 个不同元素，判断能否恰好把 $\textit{nums}$ 清空。

既然是 $k$ 个 $k$ 个地删，那么首先要满足 $\textit{nums}$ 的长度 $n$ 是 $k$ 的倍数。

设 $m$ 是出现次数最多的元素的出现次数。

对于 $k=2$ 的情况，我们有如下定理。

**定理**：如果 $m > n-m$，即出现次数最多的元素，比其余元素的个数还要多，那么无法满足题目要求，否则可以满足。

[证明+具体操作方案](https://zhuanlan.zhihu.com/p/1945782212176909162)

推广到一般情况，每次把 $m$ 减少 $1$，其余元素的个数就要减少 $k-1$，所以其余元素的个数必须至少为 $m\cdot(k-1)$，即

$$
m\cdot(k-1) \le n-m
$$

即

$$
m\cdot k \le n
$$

对于 C++ 和 Java，为防止乘法溢出，可以改为等价的

$$
m\le \left\lfloor\dfrac{n}{k}\right\rfloor
$$

具体请看 [视频讲解](https://www.bilibili.com/video/BV1X9eJz2EWE/?t=4m32s)，欢迎点赞关注~

## 哈希表写法

```py [sol-Python3]
class Solution:
    def partitionArray(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        if n % k:
            return False
        mx = max(Counter(nums).values())
        return mx * k <= n
```

```java [sol-Java]
class Solution {
    public boolean partitionArray(int[] nums, int k) {
        int n = nums.length;
        if (n % k > 0) {
            return false;
        }
        Map<Integer, Integer> cnt = new HashMap<>();
        int mx = 0;
        for (int x : nums) {
            int c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
            mx = Math.max(mx, c);
        }
        return mx <= n / k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool partitionArray(vector<int>& nums, int k) {
        int n = nums.size();
        if (n % k) {
            return false;
        }
        unordered_map<int,int> cnt;
        int mx = 0;
        for (int x : nums) {
            mx = max(mx, ++cnt[x]);
        }
        return mx <= n / k;
    }
};
```

```go [sol-Go]
func partitionArray(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	cnt := map[int]int{}
	mx := 0
	for _, x := range nums {
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	return mx <= n/k
}
```

## 数组写法

```py [sol-Python3]
class Solution:
    def partitionArray(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        if n % k:
            return False
        cnt = [0] * (max(nums) + 1)
        for x in nums:
            cnt[x] += 1
        return max(cnt) * k <= n
```

```java [sol-Java]
class Solution {
    public boolean partitionArray(int[] nums, int k) {
        int n = nums.length;
        if (n % k > 0) {
            return false;
        }

        int u = 0;
        for (int x : nums) {
            u = Math.max(u, x);
        }

        int[] cnt = new int[u + 1];
        int mx = 0;
        for (int x : nums) {
            cnt[x]++;
            mx = Math.max(mx, cnt[x]);
        }
        return mx <= n / k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool partitionArray(vector<int>& nums, int k) {
        int n = nums.size();
        if (n % k) {
            return false;
        }
        vector<int> cnt(ranges::max(nums) + 1);
        int mx = 0;
        for (int x : nums) {
            mx = max(mx, ++cnt[x]);
        }
        return mx <= n / k;
    }
};
```

```go [sol-Go]
func partitionArray(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	cnt := make([]int, slices.Max(nums)+1)
	mx := 0
	for _, x := range nums {
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	return mx <= n/k
}
```

## 优化

如果在遍历中发现 $\textit{cnt}[x] > \left\lfloor\dfrac{n}{k}\right\rfloor$，就立刻返回 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def partitionArray(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        if n % k:
            return False
        cnt = [0] * (max(nums) + 1)
        n //= k  # 避免在循环中反复做除法
        for x in nums:
            cnt[x] += 1
            if cnt[x] > n:
                return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean partitionArray(int[] nums, int k) {
        int n = nums.length;
        if (n % k > 0) {
            return false;
        }

        int u = 0;
        for (int x : nums) {
            u = Math.max(u, x);
        }

        int[] cnt = new int[u + 1];
        for (int x : nums) {
            cnt[x]++;
            if (cnt[x] > n / k) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool partitionArray(vector<int>& nums, int k) {
        int n = nums.size();
        if (n % k) {
            return false;
        }
        vector<int> cnt(ranges::max(nums) + 1);
        int mx = 0;
        for (int x : nums) {
            cnt[x]++;
            if (cnt[x] > n / k) {
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func partitionArray(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	cnt := make([]int, slices.Max(nums)+1)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > n/k {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+U)$。

## 专题训练

见下面贪心题单的「**§1.8 相邻不同**」。

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
