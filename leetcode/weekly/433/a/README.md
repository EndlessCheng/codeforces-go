## 方法一：暴力

```py [sol-Python3]
class Solution:
    def subarraySum(self, nums: List[int]) -> int:
        ans = 0
        for i, num in enumerate(nums):
            ans += sum(nums[max(i - num, 0): i + 1])
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def subarraySum(self, nums: List[int]) -> int:
        return sum(sum(nums[max(i - num, 0): i + 1])
                   for i, num in enumerate(nums))
```

```java [sol-Java]
class Solution {
    public int subarraySum(int[] nums) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = Math.max(i - nums[i], 0); j <= i; j++) {
                ans += nums[j];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int subarraySum(vector<int>& nums) {
        int ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = max(i - nums[i], 0); j <= i; j++) {
                ans += nums[j];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func subarraySum(nums []int) (ans int) {
	for i, num := range nums {
		for _, x := range nums[max(i-num, 0) : i+1] {
			ans += x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：前缀和

[原理讲解](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

计算 $\textit{nums}$ 的前缀和，即可 $\mathcal{O}(1)$ 计算任意子数组的元素和。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17RwBeqErJ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def subarraySum(self, nums: List[int]) -> int:
        s = list(accumulate(nums, initial=0))
        ans = 0
        for i, num in enumerate(nums):
            ans += s[i + 1] - s[max(i - num, 0)]
        return ans
```

```java [sol-Java]
class Solution {
    public int subarraySum(int[] nums) {
        int n = nums.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + nums[i];
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += s[i + 1] - s[Math.max(i - nums[i], 0)];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int subarraySum(vector<int>& nums) {
        int n = nums.size();
        vector<int> s(n + 1); // 前缀和
        partial_sum(nums.begin(), nums.end(), s.begin() + 1);

        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += s[i + 1] - s[max(i - nums[i], 0)];
        }
        return ans;
    }
};
```

```go [sol-Go]
func subarraySum(nums []int) (ans int) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	for i, num := range nums {
		ans += s[i+1] - s[max(i-num, 0)]
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法三：差分数组

[原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)（推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看）。

横看成岭侧成峰，用差分数组计算每个 $\textit{nums}[i]$ 要加到答案中多少次。

做法是把区间 $[\max(i-\textit{nums}[i], 0), i]$ 加一，再计算差分数组的前缀和。

```py [sol-Python3]
class Solution:
    def subarraySum(self, nums: List[int]) -> int:
        diff = [0] * (len(nums) + 1)
        for i, num in enumerate(nums):
            diff[max(i - num, 0)] += 1
            diff[i + 1] -= 1

        ans = sd = 0
        for x, d in zip(nums, diff):
            sd += d
            ans += x * sd
        return ans
```

```java [sol-Java]
class Solution {
    public int subarraySum(int[] nums) {
        int n = nums.length;
        int[] diff = new int[n + 1];
        for (int i = 0; i < n; i++) {
            diff[Math.max(i - nums[i], 0)]++;
            diff[i + 1]--;
        }

        int ans = 0, sd = 0;
        for (int i = 0; i < n; i++) {
            sd += diff[i];
            ans += nums[i] * sd;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int subarraySum(vector<int>& nums) {
        int n = nums.size();
        vector<int> diff(n + 1);
        for (int i = 0; i < n; i++) {
            diff[max(i - nums[i], 0)]++;
            diff[i + 1]--;
        }

        int ans = 0, sd = 0;
        for (int i = 0; i < n; i++) {
            sd += diff[i];
            ans += nums[i] * sd;
        }
        return ans;
    }
};
```

```go [sol-Go]
func subarraySum(nums []int) (ans int) {
	diff := make([]int, len(nums)+1)
	for i, num := range nums {
		diff[max(i-num, 0)]++
		diff[i+1]--
	}

	sd := 0
	for i, x := range nums {
		sd += diff[i]
		ans += x * sd
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面数据结构题单中的「**§1.1 前缀和基础**」和「**§2.1 一维差分**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
