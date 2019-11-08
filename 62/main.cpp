class Solution {
public:
  int uniquePaths(int m, int n) {
    return round(tgamma(m + n - 1) / tgamma(m) / tgamma(n));
  }
};