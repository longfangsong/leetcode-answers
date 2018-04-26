#include <algorithm>
#include <climits>
#include <vector>

using std::max_element;
using std::vector;
#define NO_VALUE -1
class Solution {
 private:
  vector<vector<int>> buffer;

 public:
  int longestIncreasingPathFrom(const vector<vector<int>>& matrix, int x,
                                int y) {
    if (buffer[x][y] != NO_VALUE) {
      return buffer[x][y];
    }
    vector<int> connected_grids_values;
    if (x > 0 && matrix[x - 1][y] > matrix[x][y]) {
      connected_grids_values.push_back(
          longestIncreasingPathFrom(matrix, x - 1, y));
    }
    if (x < matrix.size() - 1 && matrix[x + 1][y] > matrix[x][y]) {
      connected_grids_values.push_back(
          longestIncreasingPathFrom(matrix, x + 1, y));
    }
    if (y > 0 && matrix[x][y - 1] > matrix[x][y]) {
      connected_grids_values.push_back(
          longestIncreasingPathFrom(matrix, x, y - 1));
    }
    if (y < matrix.front().size() - 1 && matrix[x][y + 1] > matrix[x][y]) {
      connected_grids_values.push_back(
          longestIncreasingPathFrom(matrix, x, y + 1));
    }
    if (connected_grids_values.empty()) {
      buffer[x][y] = 0;
    } else {
      buffer[x][y] = *(max_element(connected_grids_values.begin(),
                                   connected_grids_values.end())) +
                     1;
    }
    return buffer[x][y];
  }
  int longestIncreasingPath(vector<vector<int>>& matrix) {
    if (matrix.size() == 0) {
      return 0;
    }
    vector<int> results;
    buffer =
        vector<vector<int>>(matrix.size(), vector<int>(matrix[0].size(), -1));
    for (int row_index = 0; row_index < matrix.size(); ++row_index) {
      for (int col_index = 0; col_index < matrix[row_index].size();
           ++col_index) {
        results.push_back(
            longestIncreasingPathFrom(matrix, row_index, col_index));
      }
    }
    return *max_element(results.begin(), results.end()) + 1;
  }
};