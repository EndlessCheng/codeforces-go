设 $x=\textit{nums}[i]$。

遍历 $\textit{nums}$，同时维护 $x$ 的最高位的出现次数 $\textit{cnt}$。枚举 $[1,9]$ 内的数字 $y$，如果与 $x\bmod 10$ 互质则答案加上 $\textit{cnt}[y]$。

具体请看[【周赛 351 视频讲解】](https://www.bilibili.com/video/BV1du41187ZN/)的第一题，欢迎点赞！

```py [sol-Python3]
class Solution:
    def countBeautifulPairs(self, nums: List[int]) -> int:
        ans, cnt = 0, [0] * 10
        for x in nums:
            for y in range(1, 10):
                if cnt[y] and gcd(x % 10, y) == 1:
                    ans += cnt[y]
            while x >= 10: x //= 10  # 这里需要 O(log x) 的时间
            cnt[x] += 1  # 统计最高位的出现次数
        return ans
```

```java [sol-Java]
class Solution {
    public int countBeautifulPairs(int[] nums) {
        int ans = 0;
        var cnt = new int[10];
        for (int x : nums) {
            for (int y = 1; y < 10; y++)
                if (cnt[y] > 0 && gcd(x % 10, y) == 1)
                    ans += cnt[y];
            while (x >= 10) x /= 10; // 这里需要 O(log x) 的时间
            cnt[x]++; // 统计最高位的出现次数
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countBeautifulPairs(vector<int> &nums) {
        int ans = 0, cnt[10]{};
        for (int x: nums) {
            for (int y = 1; y < 10; y++)
                if (cnt[y] && gcd(x % 10, y) == 1)
                    ans += cnt[y];
            while (x >= 10) x /= 10; // 这里需要 O(log x) 的时间
            cnt[x]++; // 统计最高位的出现次数
        }
        return ans;
    }
};
```

```go [sol-Go]
func countBeautifulPairs(nums []int) (ans int) {
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 { // 这里需要 O(log x) 的时间
			x /= 10
		}
		cnt[x]++ // 统计最高位的出现次数
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(k+\log U))$，其中 $n$ 为 $\textit{nums}$ 的长度，$k=10$，$U=\max(\textit{nums})$。GCD 的时间视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(k)$。
