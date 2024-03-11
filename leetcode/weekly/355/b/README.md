从暴力枚举到一次遍历（Python/Java/C++/Go/JS/Rust）

请看 [视频讲解](https://www.bilibili.com/video/BV1n8411m7Fs/?t=3m26s) 第二题。

## 暴力思路

从 $i$ 开始，只要 $\textit{nums}[i-1] \le \textit{nums}[i]$，就不断向左合并数字。

例如示例 1 从 $i=3$ 开始，可以一直合并到 $i=0$。

$$
[2,3,\underline{7},\underline{9},3]\rightarrow[2,\underline{3},\underline{16},3]\rightarrow[\underline{2},\underline{19},3]\rightarrow[21,3]
$$

枚举从 $i=0,1,2,\cdots,n-1$ 开始向左合并，取合并后的最大值为答案。

时间复杂度为 $\mathcal{O}(n^2)$，无法通过。

## 优化

如果我们能从 $i=n-1$ 开始向左合并到 $j$，那么后续枚举从 $i=n-2,n-3,\cdots,j$ 开始向左合并，是不会合并出更大的数的，因为 $\textit{nums}$ 中的元素都是正数，这样合并出来的数字只会更小。

此外，这也意味着 

$$
\textit{nums}[j-1] > \textit{nums}[j] + \textit{nums}[j+1] + \cdots + \textit{nums}[n-1]
$$

所以下一个枚举的 $i$ 可以直接从 $j-1$ 开始，这样合并出的数字一定比从 $i=n-1$ 开始更大。

```py [sol-Python3]
class Solution:
    def maxArrayValue(self, nums: List[int]) -> int:
        s = nums[-1]
        for i in range(len(nums) - 2, -1, -1):
            s = s + nums[i] if nums[i] <= s else nums[i]
        return s
```

```java [sol-Java]
class Solution {
    public long maxArrayValue(int[] nums) {
        int n = nums.length;
        long sum = nums[n - 1];
        for (int i = n - 2; i >= 0; i--) {
            sum = nums[i] <= sum ? sum + nums[i] : nums[i];
        }
        return sum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxArrayValue(vector<int> &nums) {
        long long sum = nums.back();
        for (int i = nums.size() - 2; i >= 0; i--) {
            sum = nums[i] <= sum ? sum + nums[i] : nums[i];
        }
        return sum;
    }
};
```

```go [sol-Go]
func maxArrayValue(nums []int) int64 {
	n := len(nums)
	sum := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= sum {
			sum += nums[i] // 继续向左合并
		} else {
			sum = nums[i]
		}
	}
	return int64(sum)
}
```

```js [sol-JavaScript]
var maxArrayValue = function(nums) {
    const n = nums.length;
    let sum = nums[n - 1];
    for (let i = n - 2; i >= 0; i--) {
        sum = nums[i] <= sum ? sum + nums[i] : nums[i];
    }
    return sum;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_array_value(nums: Vec<i32>) -> i64 {
        let n = nums.len();
        let mut sum = nums[n - 1] as i64;
        for i in (0..n - 1).rev() {
            let x = nums[i] as i64;
            sum = if x <= sum { sum + x } else { x };
        }
        sum
    }
}
```

（部分语言）也可以直接在原数组上合并。

```py [sol-Python3]
class Solution:
    def maxArrayValue(self, nums: List[int]) -> int:
        for i in range(len(nums) - 1, 0, -1):
            if nums[i - 1] <= nums[i]:
                nums[i - 1] += nums[i]  # 把合并值向左传
        return nums[0]
```

```go [sol-Go]
func maxArrayValue(nums []int) int64 {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] <= nums[i] {
			nums[i-1] += nums[i] // 把合并值向左传
		}
	}
	return int64(nums[0])
}
```

```js [sol-JavaScript]
var maxArrayValue = function(nums) {
    for (let i = nums.length - 1; i > 0; i--) {
        if (nums[i - 1] <= nums[i]) {
            nums[i - 1] += nums[i]; // 把合并值向左传
        }
    }
    return nums[0];
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
