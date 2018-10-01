#include <map>
#include <numeric>
#include <vector>
using namespace std;
int gcd(int m, int n) {
  if (m == 0) return n;
  return gcd(n % m, m);
}
class Solution {
 public:
  bool hasGroupsSizeX(vector<int>& deck) {
    std::map<int, size_t> count;
    for (auto&& i : deck) {
      ++count[i];
    }
    vector<int> count_nums;
    for (auto&& p : count) {
      count_nums.push_back(p.second);
    }
    int initial_value = count_nums.back();
    count_nums.pop_back();
    return accumulate(count_nums.begin(), count_nums.end(), initial_value,
                      [](int a, int b) { return gcd(a, b); }) >= 2;
  }
};