## 方法一：遍历

按照题目要求计算即可。

```py [sol-Python3]
class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        return sum(x * x for i, x in enumerate(nums, 1)
                         if len(nums) % i == 0)
```

```java [sol-Java]
class Solution {
    public int sumOfSquares(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfSquares(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumOfSquares(nums []int) (ans int) {
	for i, x := range nums {
		if len(nums)%(i+1) == 0 {
			ans += x * x
		}
	}
	return
}
```

```js [sol-JavaScript]
var sumOfSquares = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1];
        }
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：枚举因子

根据题意，如果 $i$ 是 $n$ 的因子，那我们也找到了另一个因子 $\dfrac{n}{i}$。所以只需枚举 $\sqrt{n}$ 以内的 $i$。

```py [sol-Python3]
class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(1, isqrt(n) + 1):
            if n % i == 0:
                ans += nums[i - 1] ** 2  # 注意数组的下标还是从 0 开始的
                if i * i < n:  # 避免重复统计
                    ans += nums[n // i - 1] ** 2
        return ans
```

```java [sol-Java]
class Solution {
    public int sumOfSquares(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 1; i * i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1]; // 注意数组的下标还是从 0 开始的
                if (i * i < n) { // 避免重复统计
                    ans += nums[n / i - 1] * nums[n / i - 1];
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
    int sumOfSquares(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 1; i * i <= n; i++) {
            if (n % i == 0) {
                ans += nums[i - 1] * nums[i - 1]; // 注意数组的下标还是从 0 开始的
                if (i * i < n) { // 避免重复统计
                    ans += nums[n / i - 1] * nums[n / i - 1];
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumOfSquares(nums []int) (ans int) {
	n := len(nums)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ans += nums[i-1] * nums[i-1] // 注意数组的下标还是从 0 开始的
			if i*i < n { // 避免重复统计
				ans += nums[n/i-1] * nums[n/i-1]
			}
		}
	}
	return
}
```

```js [sol-JavaScript]
var sumOfSquares = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i * i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1]; // 注意数组的下标还是从 0 开始的
            if (i * i < n) { // 避免重复统计
                ans += nums[n / i - 1] * nums[n / i - 1];
            }
        }
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt{n})$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
