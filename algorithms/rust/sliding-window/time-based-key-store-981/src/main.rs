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
        if !self.table.contains_key(&key) {
            return "".to_string();
        }
        let mut last = "".to_string();
        for tup in self.table.get(&key).unwrap() {
            //dbg!(&key, &tup);
            if tup.1 == timestamp {
                return tup.0.clone();
            }
            if tup.1 < timestamp {
                last = tup.0.clone();
            }
            if tup.1 > timestamp {
                break;
            }
        }
        return last;
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
