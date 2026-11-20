[视频讲解](https://www.bilibili.com/video/BV1Sm411U7cR/) 第四题，额外讲了另一种做法，以及如何分类思考子序列 DP。

## 提示 1

把数组排序。

为什么？设我们选的元素在排序后为 $b$，那么有 $b[i] + 1 = b[i+1]$，这意味着 $b$ 中元素在操作前，必然有 $b[i] \le b[i+1]$。反证：如果操作前 $b[i] > b[i+1]$，那么操作后 $b[i]$ 至多和 $b[i+1]$ 相等，不会出现 $b[i+1]$ 比 $b[i]$ 多 $1$ 的情况。

所以可以排序。

## 提示 2

排序后，我们选的是 $\textit{nums}$ 中的一个**子序列**。

定义 $f[x]$ 表示子序列的最后一个数是 $x$ 时，子序列的最大长度。

从左到右遍历数组 $x = \textit{nums}[i]$：

- 如果操作，那么 $x+1$ 可以接在末尾为 $x$ 的子序列后面，即 $f[x+1] = f[x] + 1$。
- 如果不操作，那么 $x$ 可以接在末尾为 $x-1$ 的子序列后面，即 $f[x] = f[x-1] + 1$。

比如 $\textit{nums} = [1,2,2]$：

- 遍历到 $\textit{nums}[0]=1$ 时，$f[2]=1,\ f[1]=1$。
- 遍历到 $\textit{nums}[1]=2$ 时，$f[3]=f[2]+1=2,\ f[2]=f[1]+1=2$。注意要先计算 $f[x+1]$ 再计算 $f[x]$（不然这里会算出 $f[3]=3$）。此时 $f[1]$ 还是 $1$。
- 遍历到 $\textit{nums}[2]=2$ 时，$f[3]=f[2]+1=3,\ f[2]=f[1]+1=2$。此时 $f[1]$ 还是 $1$。

最后返回 $f[x]$ 的最大值。

```py [sol-Python3]
class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        f = defaultdict(int)
        for x in nums:
            f[x + 1] = f[x] + 1
            f[x] = f[x - 1] + 1
        return max(f.values())
```

```java [sol-Java]
class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, Integer> f = new HashMap<>();
        for (int x : nums) {
            f.put(x + 1, f.getOrDefault(x, 0) + 1);
            f.put(x, f.getOrDefault(x - 1, 0) + 1);
        }
        return Collections.max(f.values());
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSelectedElements(vector<int>& nums) {
        ranges::sort(nums);
        unordered_map<int, int> f;
        for (int x : nums) {
            f[x + 1] = f[x] + 1;
            f[x] = f[x - 1] + 1;
        }
        return ranges::max(f | views::values);
    }
};
```

```go [sol-Go]
func maxSelectedElements(nums []int) (ans int) {
	slices.Sort(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
		ans = max(ans, f[x], f[x+1])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化

用三个变量 $f_0,f_1,f_2$ 分别表示 $f[x-1],f[x],f[x+1]$。

```py [sol-Python3]
class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        ans = pre = f0 = f1 = f2 = 0
        for x in nums:
            if x == pre + 1:
                # 类似 70. 爬楼梯 的空间优化
                f0 = f1
                f1 = f2
            elif x == pre + 2:
                f0 = f2
                f1 = 0
            elif x > pre + 2:
                f0 = 0
                f1 = 0

            f2 = f1 + 1
            f1 = f0 + 1
            ans = max(ans, f1, f2)
            pre = x
        return ans
```

```java [sol-Java]
class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        int f0 = 0; // f[x-1]
        int f1 = 0; // f[x]
        int f2 = 0; // f[x+1]
        int pre = 0; // pre 初始值随意，不影响计算
        int ans = 0;
        for (int x : nums) {
            if (x == pre + 1) {
                // 类似 70. 爬楼梯 的空间优化
                f0 = f1;
                f1 = f2;
            } else if (x == pre + 2) {
                f0 = f2;
                f1 = 0;
            } else if (x > pre + 2) {
                f0 = 0;
                f1 = 0;
            }

            f2 = f1 + 1;
            f1 = f0 + 1;
            ans = Math.max(ans, Math.max(f1, f2));
            pre = x;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSelectedElements(vector<int>& nums) {
        ranges::sort(nums);
        int f0 = 0, f1 = 0, f2 = 0;
        int pre = 0; // pre 初始值随意，不影响计算
        int ans = 0;
        for (int x : nums) {
            if (x == pre + 1) {
                // 类似 70. 爬楼梯 的空间优化
                f0 = f1;
                f1 = f2;
            } else if (x == pre + 2) {
                f0 = f2;
                f1 = 0;
            } else if (x > pre + 2) {
                f0 = 0;
                f1 = 0;
            }

            f2 = f1 + 1;
            f1 = f0 + 1;
            ans = max(ans, max(f1, f2));
            pre = x;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSelectedElements(nums []int) (ans int) {
	slices.Sort(nums)
	f0, f1, f2 := 0, 0, 0
	pre := 0 // pre 初始值随意，不影响计算
	for _, x := range nums {
		if x == pre+1 {
			// 类似 70. 爬楼梯 的空间优化
			f0 = f1
			f1 = f2
		} else if x == pre+2 {
			f0 = f2
			f1 = 0
		} else if x > pre+2 {
			f0 = 0
			f1 = 0
		}

		f2 = f1 + 1
		f1 = f0 + 1
		ans = max(ans, f1, f2)
		pre = x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
