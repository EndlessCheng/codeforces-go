由于最大值加一后还是最大值，那么反复利用最大值即可。

设数组的最大值为 $m$，答案就是

$$
m+(m+1)+(m+2)+\cdots + (m+k-1) = \dfrac{(2m+k-1)\cdot k}{2}
$$

```py [sol-Python3]
class Solution:
    def maximizeSum(self, nums: List[int], k: int) -> int:
        return (max(nums) * 2 + k - 1) * k // 2
```

```java [sol-Java]
class Solution {
    public int maximizeSum(int[] nums, int k) {
        int m = 0;
        for (int x : nums) {
            m = Math.max(m, x);
        }
        return (m * 2 + k - 1) * k / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeSum(vector<int> &nums, int k) {
        int m = *max_element(nums.begin(), nums.end());
        return (m * 2 + k - 1) * k / 2;
    }
};
```

```go [sol-Go]
func maximizeSum(nums []int, k int) int {
	return (slices.Max(nums)*2 + k - 1) * k / 2
}
```

```js [sol-JavaScript]
var maximizeSum = function(nums, k) {
    return (Math.max(...nums) * 2 + k - 1) * k / 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximize_sum(nums: Vec<i32>, k: i32) -> i32 {
        let m = *nums.iter().max().unwrap();
        (m * 2 + k - 1) * k / 2
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
