## 方法一：两次遍历

遍历 $\textit{nums}$，统计 $\textit{nums}$ 中的非零元素个数，记作 $\textit{cnt}$。

那么 $\textit{nums}$ 的前 $\textit{cnt}$ 个数中的 $0$，要和 $\textit{nums}$ 的后 $n-\textit{cnt}$ 个数中的非零元素交换。答案为 $\textit{nums}$ 的前 $\textit{cnt}$ 个数中的 $0$ 的个数。

[本题视频讲解](https://www.bilibili.com/video/BV1iuG76VEXy/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumSwaps(self, nums: list[int]) -> int:
        cnt = sum(x != 0 for x in nums) 
        return sum(x == 0 for x in nums[:cnt]) 
```

```java [sol-Java]
class Solution {
    public int minimumSwaps(int[] nums) {
        int cnt = 0;
        for (int x : nums) {
            if (x != 0) {
                cnt++;
            }
        }

        int ans = 0;
        for (int i = 0; i < cnt; i++) {
            if (nums[i] == 0) {
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
    int minimumSwaps(vector<int>& nums) {
        int cnt = 0;
        for (int x : nums) {
            cnt += x != 0;
        }

        int ans = 0;
        for (int i = 0; i < cnt; i++) {
            ans += nums[i] == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumSwaps(nums []int) (ans int) {
	cnt := 0
	for _, x := range nums {
		if x != 0 {
			cnt++
		}
	}

	for _, x := range nums[:cnt] {
		if x == 0 {
			ans++
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：相向双指针

如果 $\textit{nums}[0] \ne 0$，无需交换，问题变成剩余 $n-1$ 个数的子问题。

如果 $\textit{nums}[n-1] = 0$，无需交换，问题变成剩余 $n-1$ 个数的子问题。

如果 $\textit{nums}[0] = 0$ 且 $\textit{nums}[n-1] \ne 0$，必须交换，问题变成剩余 $n-2$ 个数的子问题。

用两个指针 $\ell$ 和 $r$ 表示 $\textit{nums}$ 的剩余元素为子数组 $[\ell, r]$。

```py [sol-Python3]
class Solution:
    def minimumSwaps(self, nums: list[int]) -> int:
        ans = 0
        l, r = 0, len(nums) - 1
        while l < r:
            if nums[l] != 0:
                l += 1
            elif nums[r] == 0:
                r -= 1
            else:
                # 交换 nums[l] 和 nums[r]
                ans += 1
                l += 1
                r -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumSwaps(int[] nums) {
        int ans = 0;
        int l = 0;
        int r = nums.length - 1;
        while (l < r) {
            if (nums[l] != 0) {
                l++;
            } else if (nums[r] == 0) {
                r--;
            } else {
                // 交换 nums[l] 和 nums[r]
                ans++;
                l++;
                r--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSwaps(vector<int>& nums) {
        int ans = 0;
        int l = 0, r = nums.size() - 1;
        while (l < r) {
            if (nums[l] != 0) {
                l++;
            } else if (nums[r] == 0) {
                r--;
            } else {
                // 交换 nums[l] 和 nums[r]
                ans++;
                l++;
                r--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumSwaps(nums []int) (ans int) {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] != 0 {
			l++
		} else if nums[r] == 0 {
			r--
		} else {
			// 交换 nums[l] 和 nums[r]
			ans++
			l++
			r--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面双指针题单的「**§3.2 相向双指针**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
