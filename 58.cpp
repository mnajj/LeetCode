class Solution {
  public:
    int lengthOfLastWord(string s) {
      string trimed = s;
      int n = 0;
      for (int i = trimed.size() - 1; i > 0; --i) {
        if (trimed[i] == ' ') {
          trimed.pop_back();
          continue;
        }
        break;
      }
      for (int i = trimed.size() - 1; i >= 0; --i) {
        if (trimed[i] == ' ') {
          break;
        }
        ++n;
      }
      return n;
    }
};
