#include <queue>
#include <utility>
using namespace std;
class RecentCounter {
 private:
  queue<int> data;

 public:
  RecentCounter():data() {}

  int ping(int t) {
      while(!data.empty() && data.front()+3000 < t) {
        data.pop();
      }
      data.push(t);
      return data.size();
  }
};

int main() {
    auto data = {39187,45399};
    return 0;
}