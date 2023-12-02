mod part01;
mod part02;

use part01::part_01;
use part02::part_02;

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("{}", part_01(&input));
    println!("{}", part_02(&input));
}
