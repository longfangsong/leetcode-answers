#include <numeric>
#include <vector>

using namespace std;

class NumArray {
 private:
  vector<long long> partial_sums;
  vector<int> m_nums;

 public:
  NumArray(vector<int> nums) : partial_sums({0}), m_nums(nums) {
    std::partial_sum(nums.begin(), nums.end(), back_inserter(partial_sums));
  }

  void update(int i, int val) {
    int delta = val - m_nums[i];
    for (int j = i+1; j < partial_sums.size(); ++j) {
        partial_sums[j] += delta;
    }
  }

  int sumRange(int i, int j) { return partial_sums[j + 1] - partial_sums[i]; }
};