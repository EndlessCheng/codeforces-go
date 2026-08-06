package copypasta

/*
生成函数相关

https://www2.math.upenn.edu/~wilf/gfology2.pdf 《generatingfunctionology》
- https://zhuanlan.zhihu.com/p/464680129 中文翻译

定义 F_0 = 0, F_1 = 1
F_{n+1} = F_n + F_{n-1}, n >= 1
定义 A(x) = sum_{i=0..} F_i * x^i

把 F 的递推式两边同时乘以 x^n，并对 n >= 1 求和
左边 = (A(x) - x) / x
右边第一项 = A(x)
右边第二项 = A(x) * x
解得 A(x) = x / (1 - x - x^2)

https://codeforces.com/blog/entry/133991?#comment-1198693 卡特兰数推广

https://codeforces.com/problemset/problem/2077/C 2300
https://codeforces.com/problemset/problem/632/E 2400
https://codeforces.com/problemset/problem/1096/G 2400
https://codeforces.com/problemset/problem/1251/F 2500
https://codeforces.com/problemset/problem/1548/C 2500 https://chatgpt.com/c/682db19e-19a4-8011-908b-50e7a797d1df
https://codeforces.com/problemset/problem/1542/E2 2700 https://www.luogu.com.cn/article/k9ipc3ow
https://codeforces.com/problemset/problem/1845/F 2800

*/

/*

https://leetcode.cn/problems/count-sequences-to-k/solutions/3906226/wo-bu-guan-wo-jiu-yao-yong-zu-he-shu-xue-lfy1/
https://chatgpt.com/c/699b9ba1-38cc-8321-9726-cfe44767a061
// C#
// 以 O(N) 线性时间倒推计算多项式 (x + 1 + x^(-1))^n 的各项系数
Z[] calc(int n) {
	if (n == 0) return [1];
	// 数组长度为 2n + 1，最高次幂为 n，最低次幂为 -n
	Z[] res = new Z[n << 1 | 1];
	// 初始化边界条件，对应最高次幂项和次高次幂项的系数
	res[^1] = 1;
	res[^2] = n;
	// 利用已知的组合数学恒等式，从高向低递推所有次幂的系数
	for (int i = n - 1; i > -n; --i) {
		res[n + i - 1] = ((n + i + 1) * res[n + i + 1] + i * res[n + i]) / (n - i + 1);
	}
	return res;
}
*/
