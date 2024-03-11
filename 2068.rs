use std::collections::HashMap;

impl Solution {
    pub fn check_almost_equivalent(word1: String, word2: String) -> bool {
        let mut map: HashMap<char, i32> = HashMap::new();
        for c in word1.chars() {
            *map.entry(c).or_insert(0) += 1;
        }
        for c in word2.chars() {
            *map.entry(c).or_insert(0) -= 1;
        }
        for (_, v) in &map {
            if *v > 3 || *v < -3 {
                return false;
            }
        }
        true
    }
}
