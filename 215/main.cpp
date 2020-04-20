#include <algorithm>
#include <vector>

using namespace std;

class Solution {
public:
  int findKthLargest(vector<int> &nums, int k) {
    nth_element(nums.begin(), nums.begin() + k - 1, nums.end(),
                std::greater<int>());
    return nums[k - 1];
  }
};