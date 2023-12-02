use std::collections::HashMap;

pub fn part_01(input: &String) -> i32 {
    let mut count = 0;
    let game_config = HashMap::from([("red", 12), ("green", 13), ("blue", 14)]);
    for game in input.lines().filter(|l| !l.is_empty()) {
        let (game_info, content) = game.split_once(": ").unwrap();
        let game_id: i32 = game_info.split_once(" ").unwrap().1.parse().unwrap();
        let mut game_is_possible = true;
        'rounds: for round in content.split("; ") {
            for draw in round.split(", ") {
                let (count_string, color) = draw.split_once(" ").unwrap();
                let count: i32 = count_string.parse().unwrap();
                if count > *game_config.get(color).unwrap() {
                    game_is_possible = false;
                    break 'rounds;
                }
            }
        }
        if game_is_possible {
            count += game_id;
        }
    }
    return count;
}
