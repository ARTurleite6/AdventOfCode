use std::collections::HashMap;

#[derive(PartialEq, Eq, Hash)]
enum Play {
    Rock,
    Scissors,
    Paper
}

type Round = (Play, Play);

fn main() {
    let rounds: HashMap<Round, u32> = HashMap::from([((Play::Scissors, Play::Scissors), 6)]);
}
