[视频讲解](https://www.bilibili.com/video/BV1DM4y1x7bR/)

## 算法一：遍历

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
        int ans = 0, n = nums.length;
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
    int sumOfSquares(vector<int> &nums) {
        int ans = 0, n = nums.size();
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
var sumOfSquares = function (nums) {
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

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 算法二：枚举因子

根据题意，$i$ 是 $n$ 的因子，此时 $\dfrac{n}{i}$ 也是 $n$ 的因子，那么只需要枚举 $\sqrt{n}$ 以内的 $i$，就可以得到大于 $\sqrt{n}$ 的另一个因子了。

```py [sol-Python3]
class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        ans, n = 0, len(nums)
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
        int ans = 0, n = nums.length;
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
    int sumOfSquares(vector<int> &nums) {
        int ans = 0, n = nums.size();
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
var sumOfSquares = function (nums) {
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

- 时间复杂度：$\mathcal{O}(\sqrt{n})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
