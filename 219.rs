use std::collections::HashMap;

impl Solution {
    pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
        let mut map: HashMap<i32, Vec<i32>> = HashMap::new();
        for (i, x) in nums.iter().enumerate() {
            let v = map.entry(*x).or_insert(vec![]);
            v.push(i as i32);
        }
        for (_, v) in map {
            for i in 0..v.len() {
                for j in i + 1..v.len() {
                    if (v[i] - v[j]).abs() <= k {
                        return true;
                    }
                }
            }
        }
        false
    }
}
