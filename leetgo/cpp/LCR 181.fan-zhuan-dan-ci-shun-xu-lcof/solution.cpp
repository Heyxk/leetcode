// Created by k at 2024/05/17 15:56
// leetgo: dev
// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    string reverseMessage(string message)
    {

        int i, j;
        i = message.size() - 1;
        while (i >= 0 && message[i] == ' ') {
            i--;
        }
        j = i;
        while (i >= 0) {
            while (i >= 0 && message[i] != ' ') {
                i--;
            }
            ans.append(message.substr(i + 1, j - i));
            ans.append(" ");
            while (i >= 0 && message[i] == ' ') {
                i--;
            }
            j = i;
        }
        if (ans.size() > 0 && ans.back() == ' ') {
            ans = ans.substr(0, ans.size() - 1);
        }
        return ans;
    }
    string ans;
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    string message;
    LeetCodeIO::scan(cin, message);

    Solution *obj = new Solution();
    auto res      = obj->reverseMessage(message);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
