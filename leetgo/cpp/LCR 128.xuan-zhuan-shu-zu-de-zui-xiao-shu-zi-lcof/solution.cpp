// Created by k at 2024/05/13 15:32
// leetgo: dev
// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    int stockManagement(vector<int> &stock)
    {
        for (int i = 0; i + 1 < stock.size(); i++) {
            if (stock[i] > stock[i + 1]) {
                return stock[i + 1];
            }
        }
        return stock[0];
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<int> stock;
    LeetCodeIO::scan(cin, stock);

    Solution *obj = new Solution();
    auto res      = obj->inventoryManagement(stock);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
