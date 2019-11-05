#include <vector>

using namespace std;

class Solution {
private:
  void backtrack(vector<vector<int>> &result, vector<int> &current_result,
                 const vector<int> &candidates, int current_result_value,
                 const int target, size_t start_from) {
    cout << "sf: " << start_from << endl;
    for (size_t index = start_from; index < candidates.size(); ++index) {
      auto i = candidates[index];
      if (current_result_value + i < target) {
        current_result.push_back(i);
        backtrack(result, current_result, candidates, current_result_value + i,
                  target, index);
        current_result.pop_back();
      } else if (current_result_value + i == target) {
        current_result.push_back(i);
        result.push_back(current_result);
        for (auto &&e : current_result) {
          cout << e << ',';
        }
        cout << endl;
        current_result.pop_back();
        break;
      } else {
        break;
      }
    }
  }

public:
  vector<vector<int>> combinationSum(vector<int> &candidates, int target) {
    sort(candidates.begin(), candidates.end());
    auto last = unique(candidates.begin(), candidates.end());
    candidates.erase(last, candidates.end());
    vector<vector<int>> result;
    vector<int> current_result;
    int current_result_value = 0;
    backtrack(result, current_result, candidates, current_result_value, target,
              0);
    return result;
  }
};