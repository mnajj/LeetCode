class Solution {
public:
    int numJewelsInStones(string jewels, string stones) {
        int arr[255] = {0};
        int sum = 0;
        for (char c : jewels) ++arr[c];
        for (char c : stones) {
            if (arr[c]) ++sum;
        }
        return sum;
    }
};
