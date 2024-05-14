// Created by k at 2024/05/14 17:13
// leetgo: dev
// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

#include <bits/stdc++.h>
#include <map>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    char dismantlingAction(string arr)
    {
        map<char, int> m;
        for (int i = 0; i < arr.length(); i++) {
            m[arr[i]] += 1;
        }
        for (int i = 0; i < arr.length(); i++) {
            if (m[arr[i]] == 1) {
                return arr[i];
            }
        }
        return ' ';
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    string arr;
    LeetCodeIO::scan(cin, arr);

    Solution *obj = new Solution();
    auto res      = obj->dismantlingAction(arr);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
