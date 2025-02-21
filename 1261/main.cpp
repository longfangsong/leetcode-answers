#include <unordered_set>
using namespace std;
struct TreeNode {
  TreeNode *left;
  TreeNode *right;
};

class FindElements {
private:
  unordered_set<int> content;

  static void collect(TreeNode *root, int value, unordered_set<int> &result) {
    if (root) {
      result.insert(value);
      collect(root->left, 2 * value + 1, result);
      collect(root->right, 2 * value + 2, result);
    }
  }

public:
  FindElements(TreeNode *root) {
    FindElements::collect(root, 0, this->content);
  }
  bool find(int target) { return content.contains(target); }
};
