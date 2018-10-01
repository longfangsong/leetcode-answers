#include <algorithm>
#include <climits>
#include <vector>
using namespace std;
class Solution {
 public:
  int partitionDisjoint(vector<int>& A) {
    vector<int> left_max;
    vector<int> right_min;
    int temp = INT_MIN;
    for (auto&& i : A) {
      temp = max(temp, i);
      left_max.push_back(temp);
    }
    temp = INT_MAX;
    for (auto it = A.rbegin(); it != A.rend(); ++it) {
      auto i = *it;
      temp = min(temp, i);
      right_min.push_back(temp);
    }
    auto it1 = left_max.begin();
    auto it2 = right_min.rbegin();
    ++it2;
    for (; it1 != left_max.end() && it2 != right_min.rend(); ++it1, ++it2) {
      if (*it1 <= *it2) {
        return distance(left_max.begin(), it1) + 1;
      }
    }
    return -1;
  }
};