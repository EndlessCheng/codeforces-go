p = i/n
e[i] = (1+e[i-1])p + (1+e[i+1])(1-p)
     = 1 + e[i-1]*p + e[i+1]*(1-p)
e[i+1] = (e[i] - 1 - e[i-1] * p) / (1 - p)

e[1] 的证明
https://zhuanlan.zhihu.com/p/649553515
