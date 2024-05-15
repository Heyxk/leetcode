// Created by k at 2024/05/15 16:47
// leetgo: dev
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    string pathEncryption(string path)
    {
        for (auto i = 0; i < path.size(); i++) {
            if (path[i] == '.') {
                path[i] = ' ';
            }
        }
        return path;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    string path;
    LeetCodeIO::scan(cin, path);

    Solution *obj = new Solution();
    auto res      = obj->pathEncryption(path);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
