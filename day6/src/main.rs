#[derive(Debug)]
struct Race {
    duration: usize,
    record: usize,
}

impl Race {
    fn new(duration: &str, record: &str) -> Self {
        let duration = duration.parse().unwrap_or(0);
        let record = record.parse().unwrap_or(0);
        Race { duration, record }
    }
}

fn solve(races: &Vec<Race>) -> usize {
    races
        .iter()
        .map(|race| {
            let mut count = 0;
            for i in 1..=race.duration {
                if (race.duration - i) * i > race.record {
                    count += 1
                }
            }
            count
        })
        .product()
}

fn parse_race_data(input: &str) -> Vec<Race> {
    let (times, records) = input.split_once("\n").unwrap_or(("", ""));
    let times = times.split_once(":").unwrap_or(("", "")).1;
    let records = records.split_once(":").unwrap_or(("", "")).1;
    times
        .split_ascii_whitespace()
        .zip(records.split_ascii_whitespace())
        .map(|(time, distance)| Race::new(time, distance))
        .collect()
}

fn main() {
    let part_one = solve(&parse_race_data(include_str!("./input.txt")));
    let part_two = solve(&parse_race_data(include_str!("./input2.txt")));
    dbg!(part_one);
    dbg!(part_two);
}
