class Solution {
public:
    int mostWordsFound(vector<string>& sentences) {
        int a[101] = {0};
        int max = 0;
        for (string s : sentences) {
            int w = 1;
            for (char c : s) {
                if (c == ' ') ++w;
            }
            if (max < w) max = w;
        }
        return max;
    }
};
