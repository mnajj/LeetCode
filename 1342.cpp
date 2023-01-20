class Solution {
public:
    int numberOfSteps(int num) {
        int s(0);
        while (num) {
            if (!(num % 2)) num /= 2;
            else --num;
            ++s;
        }
        return s;
    }
};
