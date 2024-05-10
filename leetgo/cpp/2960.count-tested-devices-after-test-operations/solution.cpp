// Created by k at 2024/05/10 17:35
// leetgo: dev
// https://leetcode.cn/problems/count-tested-devices-after-test-operations/

#include "LC_IO.h"
#include <bits/stdc++.h>
using namespace std;

// @lc code=begin

class Solution
{
  public:
    int countTestedDevices(vector<int> &batteryPercentages)
    {
        int ans = 0;
        for (auto i = 0; i < batteryPercentages.size(); i++)
        {
            if ((batteryPercentages[i] - ans) >= 1)
            {
                ans++;
            }
        }

        return ans;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<int> batteryPercentages;
    LeetCodeIO::scan(cin, batteryPercentages);

    Solution *obj = new Solution();
    auto res = obj->countTestedDevices(batteryPercentages);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
