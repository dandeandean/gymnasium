struct TimeMap {
    table: std::collections::HashMap<String, Vec<(String, i32)>>,
}

impl TimeMap {
    fn new() -> Self {
        TimeMap {
            table: std::collections::HashMap::new(),
        }
    }

    fn set(&mut self, key: String, value: String, timestamp: i32) {
        if !self.table.contains_key(&key) {
            self.table.insert(key.clone(), vec![(value, timestamp)]);
        } else {
            self.table.get_mut(&key).unwrap().push((value, timestamp));
        }
    }

    fn get(&self, key: String, timestamp: i32) -> String {
        let mut last = &"".to_string();
        if !self.table.contains_key(&key) {
            return last.to_string();
        }
        let field = self.table.get(&key).unwrap();
        if field[0].1 > timestamp {
            return last.to_string();
        }
        let (mut i, mut j): (usize, usize) = (0, field.len() - 1);
        while i <= j {
            let m = (i + j) / 2;
            if field[m].1 == timestamp {
                return field[m].0.clone();
            }
            if field[m].1 < timestamp {
                last = &field[m].0;
                i = m + 1;
            } else {
                j = m - 1;
            }
        }
        return last.to_string();
    }
}
fn main() {
    let mut tm = TimeMap::new();
    tm.set("love".to_string(), "high".to_string(), 10);
    tm.set("love".to_string(), "low".to_string(), 20);
    dbg!(
        tm.get("love".to_string(), 5), //
        tm.get("love".to_string(), 10),
        tm.get("love".to_string(), 15),
        tm.get("love".to_string(), 20),
        tm.get("love".to_string(), 25),
    );
}
