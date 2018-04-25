#include <algorithm>
#include <iterator>
#include <string>
#include <utility>
#include <vector>

using std::any_of;
using std::back_inserter;
using std::for_each;
using std::make_pair;
using std::pair;
using std::string;
using std::transform;
using std::vector;

typedef pair<int, int> Coordinate;
class Solution {
 private:
  vector<vector<Coordinate>> m_result;

 public:
  bool canQueenControlGrid(Coordinate queenPosition, Coordinate target) {
    return queenPosition.first == target.first ||
           queenPosition.second == target.second ||
           queenPosition.first - queenPosition.second ==
               target.first - target.second ||
           queenPosition.first + queenPosition.second ==
               target.first + target.second;
  }

  bool canAddQueenAt(const vector<Coordinate>& already_putted_queens,
                     Coordinate coordinate) {
    return !any_of(already_putted_queens.begin(), already_putted_queens.end(),
                   [&](const Coordinate& anotherQueen) {
                     return canQueenControlGrid(anotherQueen, coordinate);
                   });
  }

  void NQueens(vector<Coordinate> already_putted_queens, size_t needQueens) {
    if (already_putted_queens.size() == needQueens) {
      m_result.push_back(already_putted_queens);
      return;
    }
    int y;
    if (already_putted_queens.empty()) {
      y = 0;
    } else {
      y = already_putted_queens.back().second + 1;
    }
    for (int x = 0; x < needQueens; ++x) {
      if (canAddQueenAt(already_putted_queens, make_pair(x, y))) {
        already_putted_queens.push_back(make_pair(x, y));
        NQueens(already_putted_queens, needQueens);
        already_putted_queens.pop_back();
      }
    }
  }

  vector<vector<string>> solveNQueens(int n) {
    NQueens(vector<Coordinate>(), n);
    vector<vector<string>> result;
    transform(m_result.begin(), m_result.end(), back_inserter(result),
              [=](const vector<Coordinate>& queens) {
                vector<string> r(n, string(n, '.'));
                for_each(queens.begin(), queens.end(),
                         [&](const Coordinate& coord) {
                           r[coord.second][coord.first] = 'Q';
                         });
                return r;
              });
    return result;
  }
};