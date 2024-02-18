## 方法一：字符串+哈希表

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        for s in map(str, arr1):
            for i in range(1, len(s) + 1):
                st.add(s[:i])

        ans = 0
        for s in map(str, arr2):
            for i in range(1, len(s) + 1):
                if s[:i] not in st:
                    break
                ans = max(ans, i)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<String> st = new HashSet<>();
        for (int x : arr1) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) {
                st.add(s.substring(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            String s = Integer.toString(x);
            for (int i = 1; i <= s.length(); i++) {
                if (!st.contains(s.substring(0, i))) {
                    break;
                }
                ans = Math.max(ans, i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int> &arr1, vector<int> &arr2) {
        unordered_set<string> st;
        for (int x : arr1) {
            string s = to_string(x);
            for (int i = 1; i <= s.length(); i++) {
                st.insert(s.substr(0, i));
            }
        }

        int ans = 0;
        for (int x : arr2) {
            string s = to_string(x);
            for (int i = 1; i <= s.length(); i++) {
                if (!st.contains(s.substr(0, i))) {
                    break;
                }
                ans = max(ans, i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := map[string]bool{}
	for _, x := range arr1 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s); i++ {
			has[s[:i]] = true
		}
	}

	for _, x := range arr2 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s) && has[s[:i]]; i++ {
			ans = max(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log^2 U)$，其中 $n$ 为 $\textit{arr}_1$ 的长度，$m$ 为 $\textit{arr}_2$ 的长度，$U$ 为数组元素的最大值。
- 空间复杂度：$\mathcal{O}(n\log^2 U)$。

## 方法二：数位数组+哈希表

```py [sol-Python3]
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        st = set()
        a = []
        for x in arr1:
            a.clear()
            while x:
                a.append(x % 10)
                x //= 10
            x = 0
            for i in range(len(a) - 1, -1, -1):
                x = x * 10 + a[i]
                st.add(x)

        ans = 0
        for x in arr2:
            a.clear()
            while x:
                a.append(x % 10)
                x //= 10
            x = 0
            for i in range(len(a) - 1, -1, -1):
                x = x * 10 + a[i]
                if x not in st:
                    break
                ans = max(ans, len(a) - i)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<Integer> st = new HashSet<>();
        int[] a = new int[9];
        for (int x : arr1) {
            int index = 0;
            for (; x > 0; x /= 10) {
                a[index++] = x % 10;
            }
            x = 0;
            for (int i = index - 1; i >= 0; i--) {
                x = x * 10 + a[i];
                st.add(x);
            }
        }

        int ans = 0;
        for (int x : arr2) {
            int index = 0;
            for (; x > 0; x /= 10) {
                a[index++] = x % 10;
            }
            x = 0;
            for (int i = index - 1; i >= 0; i--) {
                x = x * 10 + a[i];
                if (!st.contains(x)) {
                    break;
                }
                ans = Math.max(ans, index - i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestCommonPrefix(vector<int> &arr1, vector<int> &arr2) {
        unordered_set<int> st;
        vector<int> a;
        for (int x : arr1) {
            a.clear();
            for (; x; x /= 10) {
                a.push_back(x % 10);
            }
            x = 0;
            for (int i = a.size() - 1; i >= 0; i--) {
                x = x * 10 + a[i];
                st.insert(x);
            }
        }

        int ans = 0;
        for (int x : arr2) {
            a.clear();
            for (; x; x /= 10) {
                a.push_back(x % 10);
            }
            x = 0;
            for (int i = a.size() - 1; i >= 0; i--) {
                x = x * 10 + a[i];
                if (!st.contains(x)) {
                    break;
                }
                ans = max(ans, (int) a.size() - i);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
	has := map[int]bool{}
	a := []int{}
	for _, v := range arr1 {
		a = a[:0]
		for x := v; x > 0; x /= 10 {
			a = append(a, x%10)
		}
		v = 0
		for i := len(a) - 1; i >= 0; i-- {
			v = v*10 + a[i]
			has[v] = true
		}
	}

	for _, v := range arr2 {
		a = a[:0]
		for x := v; x > 0; x /= 10 {
			a = append(a, x%10)
		}
		v = 0
		for i := len(a) - 1; i >= 0; i-- {
			v = v*10 + a[i]
			if !has[v] {
				break
			}
			ans = max(ans, len(a)-i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log U)$，其中 $n$ 为 $\textit{arr}_1$ 的长度，$m$ 为 $\textit{arr}_2$ 的长度，$U$ 为数组元素的最大值。
- 空间复杂度：$\mathcal{O}(n\log U)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
