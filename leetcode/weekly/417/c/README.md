问题等价于如下两个问题：

- 每个元音字母至少出现一次，并且**至少**包含 $k$ 个辅音字母的子串个数。记作 $f_k$。
- 每个元音字母至少出现一次，并且**至少**包含 $k+1$ 个辅音字母的子串个数。记作 $f_{k+1}$。

二者相减，所表达的含义就是**恰好**包含 $k$ 个辅音字母了，所以答案为 $f_k - f_{k+1}$。

对于每个问题，由于子串越长，越满足要求，有单调性，所以可以用**滑动窗口**解决。如果你不了解滑动窗口，可以看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

如果你之前没有做过统计子串/子数组个数的滑动窗口，推荐先完成 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/)（[我的题解](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/)），这也是一道至少+统计个数的问题，且比本题要简单许多。

## 答疑

**问**：能不能把 $f_k$ 定义成「至多」？

**答**：至多和前面的「每个元音字母**至少**出现一次」冲突，「至少」要求子串越长越好，而「至多」要求子串越短越好，这样必须分开求解（总共要计算四个滑动窗口），相比下面代码的直接求解要麻烦许多。

**问**：代码中的 `ans += left` 是什么意思？

**答**：滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1TqxCeZEmb/)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def f(self, word: str, k: int) -> int:
        cnt1 = defaultdict(int)  # 每种元音的个数
        ans = cnt2 = left = 0  # cnt2 维护辅音个数
        for b in word:
            if b in "aeiou":
                cnt1[b] += 1
            else:
                cnt2 += 1
            while len(cnt1) == 5 and cnt2 >= k:
                out = word[left]
                if out in "aeiou":
                    cnt1[out] -= 1
                    if cnt1[out] == 0:
                        del cnt1[out]
                else:
                    cnt2 -= 1
                left += 1
            ans += left
        return ans

    def countOfSubstrings(self, word: str, k: int) -> int:
        return self.f(word, k) - self.f(word, k + 1)
```

```java [sol-Java]
class Solution {
    public long countOfSubstrings(String word, int k) {
        char[] s = word.toCharArray();
        return f(s, k) - f(s, k + 1);
    }

    private long f(char[] word, int k) {
        long ans = 0;
        // 这里用哈希表实现，替换成数组会更快
        HashMap<Character, Integer> cnt1 = new HashMap<>(); // 每种元音的个数
        int cnt2 = 0; // 辅音个数
        int left = 0;
        for (char b : word) {
            if ("aeiou".indexOf(b) >= 0) {
                cnt1.merge(b, 1, Integer::sum); // ++cnt1[b]
            } else {
                cnt2++;
            }
            while (cnt1.size() == 5 && cnt2 >= k) {
                char out = word[left];
                if ("aeiou".indexOf(out) >= 0) {
                    if (cnt1.merge(out, -1, Integer::sum) == 0) { // --cnt1[out] == 0
                        cnt1.remove(out);
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    const string VOWEL = "aeiou";

    long long f(string& word, int k) {
        long long ans = 0;
        // 这里用哈希表实现，替换成数组会更快
        unordered_map<char, int> cnt1; // 每种元音的个数
        int cnt2 = 0; // 辅音个数
        int left = 0;
        for (char b : word) {
            if (VOWEL.find(b) != string::npos) {
                cnt1[b]++;
            } else {
                cnt2++;
            }
            while (cnt1.size() == 5 && cnt2 >= k) {
                char out = word[left];
                if (VOWEL.find(out) != string::npos) {
                    if (--cnt1[out] == 0) {
                        cnt1.erase(out);
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }

public:
    long long countOfSubstrings(string word, int k) {
        return f(word, k) - f(word, k + 1);
    }
};
```

```go [sol-Go]
func f(word string, k int) (ans int64) {
	// 这里用哈希表实现，替换成数组会更快
	cnt1 := map[byte]int{} // 每种元音的个数
	cnt2 := 0 // 辅音个数
	left := 0
	for _, b := range word {
		if strings.ContainsRune("aeiou", b) {
			cnt1[byte(b)]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			out := word[left]
			if strings.ContainsRune("aeiou", rune(out)) {
				cnt1[out]--
				if cnt1[out] == 0 {
					delete(cnt1, out)
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
```

## 优化一

1. 把哈希表改成数组，额外用一个变量 $\textit{size}_1$ 维护元音种类数。
2. 判断元音的代码可以用位运算优化，把 $\texttt{aeiou}$ 视作集合 $\{0,4,8,14,20\}$，根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，这等于 $1065233$。用 `(1065233 >> b & 1) > 0` 可以判断字母 $b$ 是否为元音。

> Python 代码无需优化，保持原样就行。

```java [sol-Java]
class Solution {
    public long countOfSubstrings(String word, int k) {
        char[] s = word.toCharArray();
        return f(s, k) - f(s, k + 1);
    }

    private long f(char[] word, int k) {
        final int VOWEL_MASK = 1065233;
        long ans = 0;
        int[] cnt1 = new int['u' - 'a' + 1];
        int size1 = 0; // 元音种类数
        int cnt2 = 0;
        int left = 0;
        for (char b : word) {
            b -= 'a';
            if ((VOWEL_MASK >> b & 1) > 0) {
                if (cnt1[b]++ == 0) {
                    size1++;
                }
            } else {
                cnt2++;
            }
            while (size1 == 5 && cnt2 >= k) {
                int out = word[left] - 'a';
                if ((VOWEL_MASK >> out & 1) > 0) {
                    if (--cnt1[out] == 0) {
                        size1--;
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int VOWEL_MASK = 1065233;

    long long f(string& word, int k) {
        long long ans = 0;
        int cnt1['u' - 'a' + 1]{};
        int size1 = 0; // 元音种类数
        int cnt2 = 0;
        int left = 0;
        for (char b : word) {
            b -= 'a';
            if (VOWEL_MASK >> b & 1) {
                if (cnt1[b]++ == 0) {
                    size1++;
                }
            } else {
                cnt2++;
            }
            while (size1 == 5 && cnt2 >= k) {
                char out = word[left] - 'a';
                if (VOWEL_MASK >> out & 1) {
                    if (--cnt1[out] == 0) {
                        size1--;
                    }
                } else {
                    cnt2--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }

public:
    long long countOfSubstrings(string word, int k) {
        return f(word, k) - f(word, k + 1);
    }
};
```

```go [sol-Go]
func f(word string, k int) (ans int64) {
	const vowelMask = 1065233
	cnt1 := ['u' - 'a' + 1]int{}
	size1 := 0 // 元音种类数
	cnt2 := 0
	left := 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cnt1[b] == 0 {
				size1++
			}
			cnt1[b]++
		} else {
			cnt2++
		}
		for size1 == 5 && cnt2 >= k {
			out := word[left] - 'a'
			if vowelMask>>out&1 > 0 {
				cnt1[out]--
				if cnt1[out] == 0 {
					size1--
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
```

## 优化二

把两个滑动窗口合并成一个。我一般把这种滑窗叫做**三指针滑窗**。

```py [sol-Python3]
class Solution:
    def countOfSubstrings(self, word: str, k: int) -> int:
        cnt_vowel1 = defaultdict(int)
        cnt_vowel2 = defaultdict(int)
        cnt_consonant1 = cnt_consonant2 = 0
        ans = left1 = left2 = 0
        for b in word:
            if b in "aeiou":
                cnt_vowel1[b] += 1
                cnt_vowel2[b] += 1
            else:
                cnt_consonant1 += 1
                cnt_consonant2 += 1

            while len(cnt_vowel1) == 5 and cnt_consonant1 >= k:
                out = word[left1]
                if out in "aeiou":
                    cnt_vowel1[out] -= 1
                    if cnt_vowel1[out] == 0:
                        del cnt_vowel1[out]
                else:
                    cnt_consonant1 -= 1
                left1 += 1

            while len(cnt_vowel2) == 5 and cnt_consonant2 > k:
                out = word[left2]
                if out in "aeiou":
                    cnt_vowel2[out] -= 1
                    if cnt_vowel2[out] == 0:
                        del cnt_vowel2[out]
                else:
                    cnt_consonant2 -= 1
                left2 += 1

            ans += left1 - left2
        return ans
```

```java [sol-Java]
class Solution {
    public long countOfSubstrings(String word, int k) {
        final int VOWEL_MASK = 1065233;
        char[] s = word.toCharArray();
        long ans = 0;
        int[] cntVowel1 = new int['u' - 'a' + 1], cntVowel2 = new int['u' - 'a' + 1];
        int sizeVowel1 = 0, sizeVowel2 = 0; // 元音种类数
        int cntConsonant1 = 0, cntConsonant2 = 0;
        int left1 = 0, left2 = 0;
        for (char b : s) {
            b -= 'a';
            if ((VOWEL_MASK >> b & 1) > 0) {
                if (cntVowel1[b]++ == 0) {
                    sizeVowel1++;
                }
                if (cntVowel2[b]++ == 0) {
                    sizeVowel2++;
                }
            } else {
                cntConsonant1++;
                cntConsonant2++;
            }

            while (sizeVowel1 == 5 && cntConsonant1 >= k) {
                int out = s[left1] - 'a';
                if ((VOWEL_MASK >> out & 1) > 0) {
                    if (--cntVowel1[out] == 0) {
                        sizeVowel1--;
                    }
                } else {
                    cntConsonant1--;
                }
                left1++;
            }

            while (sizeVowel2 == 5 && cntConsonant2 > k) {
                int out = s[left2] - 'a';
                if ((VOWEL_MASK >> out & 1) > 0) {
                    if (--cntVowel2[out] == 0) {
                        sizeVowel2--;
                    }
                } else {
                    cntConsonant2--;
                }
                left2++;
            }

            ans += left1 - left2;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countOfSubstrings(string word, int k) {
        const int VOWEL_MASK = 1065233;
        long long ans = 0;
        int cnt_vowel1['u' - 'a' + 1]{}, cnt_vowel2['u' - 'a' + 1]{};
        int size_vowel1 = 0, size_vowel2 = 0; // 元音种类数
        int cnt_consonant1 = 0, cnt_consonant2 = 0;
        int left1 = 0, left2 = 0;
        for (int b : word) {
            b -= 'a';
            if (VOWEL_MASK >> b & 1) {
                if (cnt_vowel1[b]++ == 0) {
                    size_vowel1++;
                }
                if (cnt_vowel2[b]++ == 0) {
                    size_vowel2++;
                }
            } else {
                cnt_consonant1++;
                cnt_consonant2++;
            }

            while (size_vowel1 == 5 && cnt_consonant1 >= k) {
                char out = word[left1] - 'a';
                if (VOWEL_MASK >> out & 1) {
                    if (--cnt_vowel1[out] == 0) {
                        size_vowel1--;
                    }
                } else {
                    cnt_consonant1--;
                }
                left1++;
            }

            while (size_vowel2 == 5 && cnt_consonant2 > k) {
                char out = word[left2] - 'a';
                if (VOWEL_MASK >> out & 1) {
                    if (--cnt_vowel2[out] == 0) {
                        size_vowel2--;
                    }
                } else {
                    cnt_consonant2--;
                }
                left2++;
            }

            ans += left1 - left2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countOfSubstrings(word string, k int) (ans int64) {
	const vowelMask = 1065233
	var cntVowel1, cntVowel2 ['u' - 'a' + 1]int
	sizeVowel1, sizeVowel2 := 0, 0 // 元音种类数
	cntConsonant1, cntConsonant2 := 0, 0
	left1, left2 := 0, 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cntVowel1[b] == 0 {
				sizeVowel1++
			}
			cntVowel1[b]++
			if cntVowel2[b] == 0 {
				sizeVowel2++
			}
			cntVowel2[b]++
		} else {
			cntConsonant1++
			cntConsonant2++
		}

		for sizeVowel1 == 5 && cntConsonant1 >= k {
			out := word[left1] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel1[out]--
				if cntVowel1[out] == 0 {
					sizeVowel1--
				}
			} else {
				cntConsonant1--
			}
			left1++
		}

		for sizeVowel2 == 5 && cntConsonant2 > k {
			out := word[left2] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel2[out]--
				if cntVowel2[out] == 0 {
					sizeVowel2--
				}
			} else {
				cntConsonant2--
			}
			left2++
		}

		ans += int64(left1 - left2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$ 或者 $\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=21$。

更多相似题目，见下面滑动窗口题单中的「**§2.3.3 恰好型滑动窗口**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
