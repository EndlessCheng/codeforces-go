统计 $\textit{nums}$ 中的偶数个数 $\textit{cnt}_0$ 和奇数个数 $\textit{cnt}_1$，那么答案就是 $\textit{cnt}_0$ 个 $0$，后跟 $\textit{cnt}_1$ 个 $1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1m39bYiEVV/)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def transformArray(self, nums: List[int]) -> List[int]:
        cnt = Counter(x % 2 for x in nums)
        return [0] * cnt[0] + [1] * cnt[1]
```

```java [sol-Java]
class Solution {
    public int[] transformArray(int[] nums) {
        int[] cnt = new int[2];
        for (int x : nums) {
            cnt[x % 2]++;
        }
        Arrays.fill(nums, 0, cnt[0], 0);
        Arrays.fill(nums, cnt[0], nums.length, 1);
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> transformArray(vector<int>& nums) {
        int cnt[2]{};
        for (int x : nums) {
            cnt[x % 2]++;
        }
        fill(nums.begin(), nums.begin() + cnt[0], 0);
        fill(nums.begin() + cnt[0], nums.end(), 1);
        return nums;
    }
};
```

```go [sol-Go]
func transformArray(nums []int) []int {
	cnt := [2]int{}
	for _, x := range nums {
		cnt[x%2]++
	}
	clear(nums[:cnt[0]]) // 置 0
	for i := cnt[0]; i < len(nums); i++ {
		nums[i] = 1
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 写法二

也可以只统计 $\textit{cnt}_1$，那么 $\textit{cnt}_0 = n-\textit{cnt}_1$。

```py [sol-Python3]
class Solution:
    def transformArray(self, nums: List[int]) -> List[int]:
        cnt1 = sum(x % 2 for x in nums)
        cnt0 = len(nums) - cnt1
        return [0] * cnt0 + [1] * cnt1
```

```java [sol-Java]
class Solution {
    public int[] transformArray(int[] nums) {
        int cnt1 = 0;
        for (int x : nums) {
            cnt1 += x % 2;
        }
        int n = nums.length;
        int cnt0 = n - cnt1;
        Arrays.fill(nums, 0, cnt0, 0);
        Arrays.fill(nums, cnt0, n, 1);
        return nums;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> transformArray(vector<int>& nums) {
        int cnt1 = 0;
        for (int x : nums) {
            cnt1 += x % 2;
        }
        fill(nums.begin(), nums.end() - cnt1, 0);
        fill(nums.end() - cnt1, nums.end(), 1);
        return nums;
    }
};
```

```go [sol-Go]
func transformArray(nums []int) []int {
	cnt1 := 0
	for _, x := range nums {
		cnt1 += x % 2
	}
	n := len(nums)
	cnt0 := n - cnt1
	clear(nums[:cnt0])
	for i := cnt0; i < n; i++ {
		nums[i] = 1
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
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
