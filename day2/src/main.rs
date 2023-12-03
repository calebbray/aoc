use std::str::FromStr;

use itertools::Itertools;

#[derive(Debug)]
struct Game {
    id: u32,
    turns: Vec<Turn>,
}

impl Game {
    fn get_max_red(&self) -> Option<u32> {
        self.turns.iter().map(|t| t.red.unwrap_or(0)).max()
    }
    fn get_max_green(&self) -> Option<u32> {
        self.turns.iter().map(|t| t.green.unwrap_or(0)).max()
    }
    fn get_max_blue(&self) -> Option<u32> {
        self.turns.iter().map(|t| t.blue.unwrap_or(0)).max()
    }
    fn get_minimal_cube_power(&self) -> u32 {
        let red = self.get_max_red().unwrap_or(1);
        let green = self.get_max_green().unwrap_or(1);
        let blue = self.get_max_blue().unwrap_or(1);

        red * green * blue
    }
}

impl IsValid for Game {
    fn is_valid(&self) -> bool {
        self.turns.iter().all(|t| t.is_valid())
    }
}

#[derive(Default, Debug)]
struct Turn {
    red: Option<u32>,
    green: Option<u32>,
    blue: Option<u32>,
}

impl Turn {
    fn new(turn: Vec<&str>) -> Self {
        let mut new_turn = Turn::default();
        for cube in turn {
            let (num, color) = cube.split_once(' ').unwrap_or(("", ""));
            let num = num.parse::<u32>().unwrap();
            match color {
                "green" => new_turn.green = Some(num),
                "red" => new_turn.red = Some(num),
                "blue" => new_turn.blue = Some(num),
                _ => unreachable!("You shouldn't be here"),
            };
        }
        new_turn
    }
}

impl IsValid for Turn {
    fn is_valid(&self) -> bool {
        let mut valid_red = true;
        let mut valid_green = true;
        let mut valid_blue = true;

        if let Some(red) = self.red {
            valid_red = red <= 12;
        }
        if let Some(green) = self.green {
            valid_green = green <= 13;
        }
        if let Some(blue) = self.blue {
            valid_blue = blue <= 14;
        }
        valid_red && valid_green && valid_blue
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

fn solve_part_two(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            let game = line.parse::<Game>().unwrap();
            game.get_minimal_cube_power()
        })
        .sum()
}

fn main() {
    let games = include_str!("./input.txt");
    let part_one = solve_part_one(games);
    let part_two = solve_part_two(games);

    dbg!(part_one);
    dbg!(part_two);
}
