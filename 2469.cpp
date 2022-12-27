class Solution {
public:
    vector<double> convertTemperature(double celsius) {
        vector<double> vec(2, 0);
        vec[0] = celsius + 273.15;
        vec[1] = celsius * 1.80 + 32.00;
        return vec;
    }
};
