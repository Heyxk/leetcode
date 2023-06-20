// Created by k at 2023/06/20 10:17
// https://leetcode.cn/problems/qiu-12n-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
public:
    int sumNums(int n) {
		char A[n][n+1];
		return sizeof(A)>>1;
    }
};
//ans=1+2+3+...+n
//   =(1+n)*n/2
//   =sizeof(bool a[n][n+1])/2
//   =sizeof(a)>>1


// @lc code=end

int main() {
	ios_base::sync_with_stdio(false);
	stringstream out_stream;

	int n; LeetCodeIO::scan(cin, n);

	Solution *obj = new Solution();

	auto res = obj->sumNums(n);

	LeetCodeIO::print(out_stream, res);
	cout << "
output: " << out_stream.rdbuf() << endl;

	delete obj;
	return 0;
}
