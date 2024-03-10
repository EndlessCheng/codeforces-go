分别计算每个 $\textit{answer}[i]$。

对于 $\textit{arr}[i]$，从小到大枚举长度 $\textit{size}$，然后枚举 $\textit{arr}[i]$ 的长为 $\textit{size}$ 的所有子串 $t$，判断 $t$ 是否在其它字符串中，如果不在，就更新 $\textit{answer}[i]$ 的最小值。

```py [sol-Python3]
class Solution:
    def shortestSubstrings(self, arr: List[str]) -> List[str]:
        def check(i: int, sub: str) -> bool:
            for j, s in enumerate(arr):
                if j != i and sub in s:
                    return False
            return True

        ans = []
        for i, s in enumerate(arr):
            m = len(s)
            res = ""
            for size in range(1, m + 1):
                for j in range(size, m + 1):
                    t = s[j - size: j]
                    if (not res or t < res) and check(i, t):
                        res = t
                if res: break
            ans.append(res)
        return ans
```

```java [sol-Java]
class Solution {
    public String[] shortestSubstrings(String[] arr) {
        int n = arr.length;
        String[] ans = new String[n];
        for (int i = 0; i < n; i++) {
            int m = arr[i].length();
            String res = "";
            for (int size = 1; size <= m && res.isEmpty(); size++) {
                for (int j = size; j <= m; j++) {
                    String t = arr[i].substring(j - size, j);
                    if ((res.isEmpty() || t.compareTo(res) < 0) && check(arr, i, t)) {
                        res = t;
                    }
                }
            }
            ans[i] = res;
        }
        return ans;
    }

    private boolean check(String[] arr, int i, String sub) {
        for (int j = 0; j < arr.length; j++) {
            if (j != i && arr[j].contains(sub)) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> shortestSubstrings(vector<string> &arr) {
        int n = arr.size();
        auto check = [&](int i, string &sub) {
            for (int j = 0; j < n; j++) {
                if (j != i && arr[j].find(sub) != string::npos) {
                    return false;
                }
            }
            return true;
        };

        vector<string> ans(n);
        for (int i = 0; i < n; i++) {
            int m = arr[i].size();
            string res;
            for (int size = 1; size <= m && res.empty(); size++) {
                for (int j = size; j <= m; j++) {
                    string t = arr[i].substr(j - size, size);
                    if ((res.empty() || t < res) && check(i, t)) {
                        res = t;
                    }
                }
            }
            ans[i] = res;
        }
        return ans;
    }
};
```

```go [sol-Go]
func shortestSubstrings(arr []string) []string {
	ans := make([]string, len(arr))
	for i, s := range arr {
		m := len(s)
		res := ""
		for size := 1; size <= m && res == ""; size++ {
		next:
			for k := size; k <= m; k++ {
				sub := s[k-size : k]
				if res != "" && sub >= res {
					continue
				}
				for j, t := range arr {
					if j != i && strings.Contains(t, sub) {
						continue next
					}
				}
				res = sub
			}
		}
		ans[i] = res
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m^4)$，其中 $n$ 为 $\textit{arr}$ 的长度，$m$ 为 $\textit{arr}[i]$ 的长度，不超过 $20$。
- 空间复杂度：$\mathcal{O}(m)$ 或 $\mathcal{O}(1)$。忽略返回值的空间。

注：线性做法可以学习「后缀数组」和「后缀自动机」。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
