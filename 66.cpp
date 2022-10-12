class Solution {
  public:
    vector <int> plusOne(vector <int> &digits) {
      int p = digits.size() - 1;
      int carry = 1;
      while (p >= 0) {
        digits[p] += carry;
        --carry;
        if (digits[p] == 10) {
          digits[p] = 0;
          ++carry;
          --p;
          continue;
        }
        break;
      }
      if (carry == 1) {
        digits.insert(digits.begin(), carry);
        --carry;
      }
      return digits;
    }
};
