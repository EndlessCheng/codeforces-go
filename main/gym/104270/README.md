假设 $d$ 中的每个数都**至少**为 $\textit{low}$，那么 $\textit{low}$ 越大，操作次数就越多，有单调性，可以**二分答案**。

由于必须从最左边开始，所以我们要从左到右计算。

在计算前，还需要说明一个转换关系：我们可以把任何一种或长或短的、来来回回的移动方式，转换成若干组「左右横跳」，也就是在 $1$ 和 $2$ 之间的左右横跳，在 $2$ 和 $3$ 之间的左右横跳，在 $3$ 和 $4$ 之间的左右横跳，……依此类推。如下图：

![](https://cdn.luogu.com.cn/upload/image_hosting/80vfah6j.png)

从第一个数开始。设 $p=a_1$，至少要增加 $k=\left\lceil\dfrac{\textit{low}}{p}\right\rceil$ 次。

第一次操作需要从 $0$ 走到 $1$，后面的 $k-1$ 次增加可以在 $1$ 和 $2$ 之间左右横跳。

所以一共需要

$$
2(k-1)+1 = 2k-1
$$

次操作。

注意这会导致下一个数已经操作了 $k-1$ 次。

如此循环，直到最后一个数。如果循环中发现操作次数已经超过 $m$，退出循环。

注意，如果最后一个数还需要操作的次数 $\le 0$，那么是不需要继续操作的，退出循环。

下面代码采用开区间二分。

- 开区间左端点初始值：$0$。无需操作，一定可以满足要求。
- 开区间右端点初始值：$\left\lceil\dfrac{m}{2}\right\rceil\cdot \min(a)+1$。假设第一个数是最小值，那么它可以通过左右横跳操作 $\left\lceil\dfrac{m}{2}\right\rceil$ 次。结果 $+1$ 之后一定无法满足要求。

[本题视频讲解](https://www.bilibili.com/video/BV1ekN2ebEHx/?t=50m34s)

```cpp
#include<bits/stdc++.h>
using namespace std;

int main() {
    cin.tie(nullptr)->sync_with_stdio(false);
    int T, n;
    long long m;
    cin >> T;
    while (T--) {
        cin >> n >> m;
        vector<int> a(n);
        for (int& x : a) {
            cin >> x;
        }

        auto check = [&](long long low) -> bool {
            long long rem = m, pre = 0;
            for (int i = 0; i < n; i++) {
                long long k = (low - 1) / a[i] + 1 - pre; // 还需要操作的次数
                if (i == n - 1 && k <= 0) { // 最后一个数已经满足要求
                    break;
                }
                k = max(k, 1LL); // 至少要走 1 步
                rem -= k * 2 - 1; // 左右横跳
                if (rem < 0) {
                    return false;
                }
                pre = k - 1; // 右边那个数顺带操作了 k-1 次
            }
            return true;
        };

        long long left = 0;
        long long right = (m + 1) / 2 * ranges::min(a) + 1;
        while (left + 1 < right) {
            long long mid = (left + right) / 2;
            (check(mid) ? left : right) = mid;
        }
        cout << left << '\n';
    }
    return 0;
}
```

**时间复杂度**：$\mathcal{O}(n\log U)$。
