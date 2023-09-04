请看 [视频讲解](https://www.bilibili.com/video/BV1Nj411178Z/)。

```py [sol-Python3]
class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        ans = 0
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            ans += n % 2 == 0 and sum(map(int, s[:n // 2])) == sum(map(int, s[n // 2:]))
        return ans
```

```java [sol-Java]
class Solution {
    public int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            char[] s = Integer.toString(i).toCharArray();
            int n = s.length;
            if (n % 2 > 0) {
                continue;
            }
            int sum = 0;
            for (int j = 0; j < n / 2; j++) {
                sum += s[j];
            }
            for (int j = n / 2; j < n; j++) {
                sum -= s[j];
            }
            if (sum == 0) {
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
    int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            auto s = to_string(i);
            int n = s.length();
            if (n % 2 == 0 && accumulate(s.begin(), s.begin() + n / 2, 0) == accumulate(s.begin() + n / 2, s.end(), 0)) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSymmetricIntegers(low int, high int) (ans int) {
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n%2 > 0 {
			continue
		}
		sum := 0
		for _, c := range s[:n/2] {
			sum += int(c)
		}
		for _, c := range s[n/2:] {
			sum -= int(c)
		}
		if sum == 0 {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var countSymmetricIntegers = function (low, high) {
    let ans = 0;
    for (let i = low; i <= high; i++) {
        const s = i.toString();
        const n = s.length;
        if (n % 2) {
            continue;
        }
        const m = n >> 1;
        let sum = 0;
        for (let j = 0; j < m; j++) {
            sum += s.charCodeAt(j);
        }
        for (let j = m; j < n; j++) {
            sum -= s.charCodeAt(j);
        }
        if (sum === 0) {
            ans++;
        }
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((\textit{high} - \textit{low})\log \textit{high})$。
- 空间复杂度：$\mathcal{O}(\log \textit{high})$。

## 思考题

你能用 [数位 DP](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m05s) 解决本题吗？

时间复杂度可以做到 $\mathcal{O}(\log^2 \textit{high})$。
