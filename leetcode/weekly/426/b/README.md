## 分析

设异常值为 $x$，元素和为 $y$，那么其余 $n-2$ 个数的和也是 $y$，所以 $x+2y$ 就是整个 $\textit{nums}$ 数组的元素和 $\textit{total}$，即

$$
x+2y = \textit{total}
$$

也就是说，问题相当于从 $\textit{nums}$ 中选出两个（下标不同的）数 $x$ 和 $y$，满足 $x+2y$ 等于一个定值。你需要计算 $x$ 的最大值是多少。

这是我们最熟悉的 [1. 两数之和](https://leetcode.cn/problems/two-sum/)。

## 方法一：枚举异常值

枚举异常值 $x=\textit{nums}[i]$，那么有

$$
2y = \textit{total}-x
$$

如果 $\textit{total}-x$ 是偶数，且 $y=\dfrac{\textit{total}-x}{2}$ 在（除去 $x$ 的）其余 $n-1$ 个数中，那么 $x$ 就是一个异常值。

用哈希表记录每个数的出现次数，从而加快判断。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1tAzoY1EUN/?t=3m49s)，欢迎点赞关注~

### 写法一

```py [sol-Python3]
class Solution:
    def getLargestOutlier(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        total = sum(nums)

        ans = -inf
        for x in nums:
            cnt[x] -= 1
            if (total - x) % 2 == 0 and cnt[(total - x) // 2] > 0:
                ans = max(ans, x)
            cnt[x] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int getLargestOutlier(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int total = 0;
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
            total += x;
        }

        int ans = Integer.MIN_VALUE;
        for (int x : nums) {
            cnt.merge(x, -1, Integer::sum);
            if ((total - x) % 2 == 0 && cnt.getOrDefault((total - x) / 2, 0) > 0) {
                ans = Math.max(ans, x);
            }
            cnt.merge(x, 1, Integer::sum);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLargestOutlier(vector<int>& nums) {
        unordered_map<int, int> cnt;
        int total = 0;
        for (int x : nums) {
            cnt[x]++;
            total += x;
        }

        int ans = INT_MIN;
        for (int x : nums) {
            cnt[x]--;
            if ((total - x) % 2 == 0 && cnt[(total - x) / 2] > 0) {
                ans = max(ans, x);
            }
            cnt[x]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		cnt[x]--
		if (total-x)%2 == 0 && cnt[(total-x)/2] > 0 {
			ans = max(ans, x)
		}
		cnt[x]++
	}
	return ans
}
```

### 写法二

```py [sol-Python3]
class Solution:
    def getLargestOutlier(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        total = sum(nums)

        ans = -inf
        for x in nums:
            y, rem = divmod(total - x, 2)
            if rem == 0 and y in cnt and (y != x or cnt[y] > 1):
                ans = max(ans, x)
        return ans
```

```java [sol-Java]
class Solution {
    public int getLargestOutlier(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int total = 0;
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
            total += x;
        }

        int ans = Integer.MIN_VALUE;
        for (int x : nums) {
            if ((total - x) % 2 == 0) {
                int y = (total - x) / 2;
                if (cnt.containsKey(y) && (y != x || cnt.get(y) > 1)) {
                    ans = Math.max(ans, x);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLargestOutlier(vector<int>& nums) {
        unordered_map<int, int> cnt;
        int total = 0;
        for (int x : nums) {
            cnt[x]++;
            total += x;
        }

        int ans = INT_MIN;
        for (int x : nums) {
            if ((total - x) % 2 == 0) {
                int y = (total - x) / 2;
                auto it = cnt.find(y);
                if (it != cnt.end() && (y != x || it->second > 1)) {
                    ans = max(ans, x);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		if (total-x)%2 == 0 {
			y := (total - x) / 2
			if cnt[y] > 1 || cnt[y] > 0 && y != x {
				ans = max(ans, x)
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举元素和

枚举 $y=\textit{nums}[i]$，那么异常值等于

$$
\textit{total} - 2y
$$

如果 $\textit{total} - 2y$ 在（去掉 $y$ 之后的）哈希表中，那么 $\textit{total} - 2y$ 就是一个异常值。

```py [sol-Python3]
class Solution:
    def getLargestOutlier(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        total = sum(nums)

        ans = -inf
        for y in nums:
            t = total - y * 2
            if t in cnt and (t != y or cnt[t] > 1):
                ans = max(ans, t)
        return ans
```

```java [sol-Java]
class Solution {
    public int getLargestOutlier(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int total = 0;
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
            total += x;
        }

        int ans = Integer.MIN_VALUE;
        for (int y : nums) {
            int t = total - y * 2;
            if (cnt.containsKey(t) && (t != y || cnt.get(t) > 1)) {
                ans = Math.max(ans, t);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLargestOutlier(vector<int>& nums) {
        unordered_map<int, int> cnt;
        int total = 0;
        for (int x : nums) {
            cnt[x]++;
            total += x;
        }

        int ans = INT_MIN;
        for (int y : nums) {
            int t = total - y * 2;
            auto it = cnt.find(t);
            if (it != cnt.end() && (t != y || it->second > 1)) {
                ans = max(ans, t);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, y := range nums {
		t := total - y*2
		if cnt[t] > 1 || cnt[t] > 0 && t != y {
			ans = max(ans, t)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

如果 $\textit{nums}$ 已经是有序的，你能想出一个 $\mathcal{O}(1)$ 额外空间的做法吗？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
