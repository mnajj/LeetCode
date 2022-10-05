#include <iostream>
#include <unordered_map>

using namespace std;

class Solution {
public:
		int romanToInt(string s) {
			int sum = 0;
			unordered_map<char, int> map;
			map = {{'I', 1},
						 {'V', 5},
						 {'X', 10},
						 {'L', 50},
						 {'C', 100},
						 {'D', 500},
						 {'M', 1000}
			};
			int i = 0;
			while (i < s.size() - 1) {
				if (s[i] == 'I') {
					if (s[i + 1] == 'V') {
						sum += 4;
						i += 2;
						continue;
					} else if (s[i + 1] == 'X') {
						sum += 9;
						i += 2;
						continue;
					}
				} else if (s[i] == 'X') {
					if (s[i + 1] == 'L') {
						sum += 40;
						i += 2;
						continue;
					} else if (s[i + 1] == 'C') {
						sum += 90;
						i += 2;
						continue;
					}
				} else if (s[i] == 'C') {
					if (s[i + 1] == 'D') {
						sum += 400;
						i += 2;
						continue;
					} else if (s[i + 1] == 'M') {
						sum += 900;
						i += 2;
						continue;
					}
				}
				sum += map[s[i]];
				i++;
			}
			sum += map[s[i]];
			return sum;
		}
};
