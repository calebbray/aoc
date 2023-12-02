use std::str::FromStr;

use itertools::Itertools;

#[derive(Debug)]
struct Game {
    id: u32,
    turns: Vec<Turn>,
}

impl IsValid for Game {
    fn is_valid(&self) -> bool {
        self.turns.iter().all(|t| t.is_valid())
    }
}

#[derive(Default, Debug)]
struct Turn {
    red: u8,
    green: u8,
    blue: u8,
}

impl Turn {
    fn new(turn: Vec<&str>) -> Self {
        let mut new_turn = Turn::default();
        for cube in turn {
            let (num, color) = cube.split_once(' ').unwrap_or(("", ""));
            let num = num.parse::<u8>().unwrap_or(0);
            match color {
                "green" => new_turn.green = num,
                "red" => new_turn.red = num,
                "blue" => new_turn.blue = num,
                _ => unreachable!("You shouldn't be here"),
            };
        }
        new_turn
    }
}

impl IsValid for Turn {
    fn is_valid(&self) -> bool {
        self.red <= 12 && self.green <= 13 && self.blue <= 14
    }
}

trait IsValid {
    fn is_valid(&self) -> bool;
}

#[derive(Debug)]
struct ParseError;
impl FromStr for Game {
    type Err = ParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (game_id, game): (&str, &str) = s.split(": ").collect_tuple().ok_or(ParseError)?;
        let game_id: &u32 = &game_id[5..].parse().expect("no proper id for given game");

        let turns = game
            .split("; ")
            .map(|turn| turn.split(", ").collect_vec())
            .collect_vec();
        let mut formatted_turns: Vec<Turn> = vec![];

        for turn in turns {
            formatted_turns.push(Turn::new(turn));
        }

        Ok(Game {
            id: *game_id,
            turns: formatted_turns,
        })
    }
}

fn solve_part_one(input: &str) -> u32 {
    input
        .lines()
        .filter_map(|line| {
            let game = line.parse::<Game>().unwrap();
            if game.is_valid() {
                return Some(game.id);
            }
            None
        })
        .sum()
}

fn solve_part_two(_input: &str) -> u32 {
    0
}

fn main() {
    let games = include_str!("./input.txt");
    let part_one = solve_part_one(games);
    let part_two = solve_part_two(games);

    dbg!(part_one);
    dbg!(part_two);
}
