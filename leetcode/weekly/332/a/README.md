由于每次操作都是取出 $\textit{nums}$ 的第一个元素和最后一个元素串联，所以相向双指针模拟即可。

例如 $x=15$ 和 $y=49$ 串联，结果等于 $x\cdot 10^2 + y=1549$。我们可以把 $y$ 不断除 $10$，来知道 $x$ 要乘上 $10$ 的多少次方。

注意最后只剩一个元素的情况。

```py [sol-Python3]
class Solution:
    def findTheArrayConcVal(self, nums: List[int]) -> int:
        ans = 0
        i = 0
        j = len(nums) - 1
        while i < j:  # 相向双指针
            x = nums[i]
            y = nums[j]
            while y:
                x *= 10
                y //= 10
            ans += x + nums[j]
            i += 1
            j -= 1
        if i == j:
            ans += nums[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long findTheArrayConcVal(int[] nums) {
        long ans = 0;
        int i = 0;
        int j = nums.length - 1;
        while (i < j) {
            int x = nums[i];
            int y = nums[j];
            while (y != 0) {
                x *= 10;
                y /= 10;
            }
            ans += x + nums[j];
            i++;
            j--;
        }
        if (i == j) {
            ans += nums[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long findTheArrayConcVal(vector<int>& nums) {
        long long ans = 0;
        int i = 0;
        int j = nums.size() - 1;
        while (i < j) {
            int x = nums[i];
            int y = nums[j];
            while (y) {
                x *= 10;
                y /= 10;
            }
            ans += x + nums[j];
            i++;
            j--;
        }
        if (i == j) {
            ans += nums[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func findTheArrayConcVal(a []int) (ans int64) {
	for len(a) > 1 {
		x := a[0]
		for y := a[len(a)-1]; y != 0; y /= 10 {
			x *= 10
		}
		ans += int64(x + a[len(a)-1])
		a = a[1 : len(a)-1]
	}
	if len(a) > 0 {
		ans += int64(a[0])
	}
	return
}
```

```js [sol-JavaScript]
var findTheArrayConcVal = function(nums) {
    let ans = 0;
    let i = 0;
    let j = nums.length - 1;
    while (i < j) {
        let x = nums[i];
        let y = nums[j];
        while (y) {
            x *= 10;
            y = Math.floor(y / 10);
        }
        ans += x + nums[j];
        i++;
        j--;
    }
    if (i === j) {
        ans += nums[i];
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_the_array_conc_val(nums: Vec<i32>) -> i64 {
        let mut ans = 0i64;
        let mut i = 0;
        let mut j = nums.len() - 1;
        while i < j {
            let mut x = nums[i];
            let mut y = nums[j];
            while y != 0 {
                x *= 10;
                y /= 10;
            }
            ans += (x + nums[j]) as i64;
            i += 1;
            j -= 1;
        }
        if i == j {
            ans += nums[i] as i64;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
