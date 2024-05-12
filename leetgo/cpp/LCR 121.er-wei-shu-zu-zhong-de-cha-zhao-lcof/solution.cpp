// Created by k at 2024/05/11 16:16
// leetgo: dev
// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    bool findTargetIn2DPlants(vector<vector<int>> &plants, int target)
    {
        int i = plants.size() - 1;
        int j = 0;
        for (; i >= 0 && j < plants[0].size();) {
            if (plants[i][j] == target) {
                return true;
            }
            else if (plants[i][j] < target) {
                j++;
            }
            else {
                i--;
            }
        }
        return false;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<vector<int>> plants;
    LeetCodeIO::scan(cin, plants);
    int target;
    LeetCodeIO::scan(cin, target);

    Solution *obj = new Solution();
    auto res      = obj->findTargetIn2DPlants(plants, target);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
