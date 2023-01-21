class Solution {
public:
		vector<int> asteroidCollision(vector<int> &asteroids) {
			vector<int> res;
			for (int x: asteroids) {
				if (res.empty() || x > 0 || (x < 0 && res.back() < 0)) {
					res.push_back(x);
					continue;
				}
				while (1) {
					if (res.empty() || res.back() < 0) {
						res.push_back(x);
						break;
					}
					if (abs(x) == res.back()) {
						res.pop_back();
						break;
					}
					if (abs(x) > res.back()) {
						res.pop_back();
						continue;
					}
					break;
				}
			}
			return res;
		}
};
