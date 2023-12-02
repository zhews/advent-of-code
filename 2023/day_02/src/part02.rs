use std::collections::HashMap;

pub fn part_02(input: &String) -> i32 {
    let mut count = 0;
    for game in input.lines().filter(|l| !l.is_empty()) {
        let (_, content) = game.split_once(": ").unwrap();
        let mut max_cubes = HashMap::new();
        for round in content.split("; ") {
            for draw in round.split(", ") {
                let (count_string, color) = draw.split_once(" ").unwrap();
                let count: i32 = count_string.parse().unwrap();
                match max_cubes.get(color) {
                    Some(v) => {
                        if count > *v {
                            max_cubes.insert(color, count);
                        }
                    }
                    None => {
                        max_cubes.insert(color, count);
                        ()
                    }
                }
            }
        }
        let mut power = 1;
        for value in max_cubes.values() {
            power *= value;
        }
        count += power;
    }
    return count;
}
