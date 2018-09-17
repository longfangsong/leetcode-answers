#include <algorithm>
#include <deque>
#include <iostream>
#include <numeric>
#include <vector>

using namespace std;

class Solution {
 public:
  int shortestSubarray(vector<int>& A, int K) {
    vector<long long> prefix_sum{0};
    std::partial_sum(A.begin(), A.end(), back_inserter(prefix_sum));
    for (auto n : prefix_sum) {
      cout << n << ',';
    }
    cout << endl;
    deque<size_t> mono_queue;
    size_t ans = A.size() + 1;
    for (size_t now_index = 0; now_index <= A.size(); ++now_index) {
      while (!mono_queue.empty() &&
             prefix_sum[now_index] <= prefix_sum[mono_queue.back()]) {
        mono_queue.pop_back();
      }
      while (!mono_queue.empty() &&
             prefix_sum[now_index] >= prefix_sum[mono_queue.front()] + K) {
        ans = min(ans, now_index - mono_queue.front());
        mono_queue.pop_front();
      }
      mono_queue.push_back(now_index);
    }
    return ans < A.size() + 1 ? ans : -1;
  }
};
