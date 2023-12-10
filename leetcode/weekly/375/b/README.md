下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

---

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般涉及到取模的题目，会用到如下两个恒等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。

对于本题，我们可以用**快速幂**计算答案（暴力 for 循环计算也可以），具体请看 [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)。

```py [sol-Python3]
class Solution:
    def getGoodIndices(self, variables: List[List[int]], target: int) -> List[int]:
        return [i for i, (a, b, c, m) in enumerate(variables)
                if pow(pow(a, b, 10), c, m) == target]
```

```java [sol-Java]
public class Solution {
    public List<Integer> getGoodIndices(int[][] variables, int target) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < variables.length; i++) {
            int[] v = variables[i];
            if (pow(pow(v[0], v[1], 10), v[2], v[3]) == target) {
                ans.add(i);
            }
        }
        return ans;
    }

    private long pow(long x, int n, int mod) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0)
                res = res * x % mod;
            x = x * x % mod;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    long long pow(long long x, int n, int m) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) res = res * x % m;
            x = x * x % m;
        }
        return res;
    }

public:
    vector<int> getGoodIndices(vector<vector<int>> &variables, int target) {
        vector<int> ans;
        for (int i = 0; i < variables.size(); i++) {
            auto &v = variables[i];
            if (pow(pow(v[0], v[1], 10), v[2], v[3]) == target) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func getGoodIndices(variables [][]int, target int) (ans []int) {
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return
}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{variables}$ 的长度，$U$ 为 $b_i$ 和 $c_i$ 的最大值，本题为 $10^3$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
