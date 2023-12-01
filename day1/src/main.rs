fn main() {
    let numbers = include_str!("./input.txt")
        .lines()
        .map(|line| {
            line.chars()
                .filter_map(|c| c.to_string().parse::<usize>().ok())
                .collect::<Vec<usize>>()
        })
        .collect::<Vec<_>>();
    let _answer = numbers
        .iter()
        .map(|line| combine_numbers(get_first_and_last_int(line)))
        .sum::<usize>();

    let part2 = include_str!("./input.txt")
        .lines()
        .map(|line| {
            let first_int = find_from_start(line).unwrap();
            let last_int = find_from_end(line).unwrap();
            combine_numbers((first_int, last_int))
        })
        .sum::<usize>();
    dbg!(part2);
}

fn get_first_and_last_int(nums: &Vec<usize>) -> (usize, usize) {
    let first = nums[0];
    let last = *nums.last().unwrap();
    (first, last)
}

fn combine_numbers((first, last): (usize, usize)) -> usize {
    format!("{first}{last}").parse::<usize>().unwrap()
}

const CASES: [(&str, usize); 20] = [
    ("zero", 0),
    ("one", 1),
    ("two", 2),
    ("three", 3),
    ("four", 4),
    ("five", 5),
    ("six", 6),
    ("seven", 7),
    ("eight", 8),
    ("nine", 9),
    ("0", 0),
    ("1", 1),
    ("2", 2),
    ("3", 3),
    ("4", 4),
    ("5", 5),
    ("6", 6),
    ("7", 7),
    ("8", 8),
    ("9", 9),
];

fn find_from_start(mut line: &str) -> Option<usize> {
    while !line.is_empty() {
        let found_num = CASES.iter().find(|(str, _)| line.starts_with(str));

        if let Some((_, num)) = found_num {
            return Some(*num);
        }

        line = &line[1..];
    }

    None
}

fn find_from_end(mut line: &str) -> Option<usize> {
    while !line.is_empty() {
        let found_num = CASES.iter().find(|(str, _)| line.ends_with(str));

        if let Some((_, num)) = found_num {
            return Some(*num);
        }

        line = &line[..line.len() - 1];
    }
    None
}
