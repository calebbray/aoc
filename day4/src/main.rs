use std::{collections::HashMap, str::FromStr};

#[derive(Debug)]
struct Card {
    id: usize,
    winners: Vec<usize>,
}

#[derive(Debug)]
struct ParseError;

impl Card {
    fn get_score(&self) -> usize {
        if self.winners.len() == 0 {
            return 0;
        }

        if self.winners.len() == 1 {
            return 1;
        }

        2_usize.pow((self.winners.len() - 1) as u32)
    }
}
impl FromStr for Card {
    type Err = ParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (id, numbers) = s.split_once(':').unwrap_or(("", ""));
        let id = id
            .split_once(' ')
            .unwrap_or(("", ""))
            .1
            .trim()
            .parse::<usize>()
            .unwrap();

        let (goals, mine) = numbers.split_once('|').unwrap_or(("", ""));
        let goals = collect_nubers(goals.trim());
        let mine = collect_nubers(mine.trim());
        let winners = get_winners(&goals, &mine);

        Ok(Card { id, winners })
    }
}

fn collect_nubers(num_str: &str) -> Vec<usize> {
    num_str
        .split(' ')
        .filter_map(|num| {
            if let Ok(parsed) = num.parse() {
                return Some(parsed);
            }
            None
        })
        .collect()
}

fn parse_input() -> Vec<Card> {
    include_str!("./input.txt")
        .lines()
        .map(|line| line.parse().unwrap())
        .collect()
}

fn get_winners(goals: &Vec<usize>, mine: &Vec<usize>) -> Vec<usize> {
    goals
        .iter()
        .filter_map(|g| {
            if mine.contains(g) {
                return Some(*g);
            }
            None
        })
        .collect()
}

fn create_winner_dict(cards: &Vec<Card>) -> HashMap<usize, usize> {
    let mut w_dict = HashMap::new();

    for card in cards {
        w_dict
            .entry(card.id)
            .and_modify(|num| *num += 1)
            .or_insert(1);
        for i in 1..=card.winners.len() {
            let copies = w_dict.get_mut(&card.id).unwrap_or(&mut 0).clone();
            let key = card.id + i;
            w_dict
                .entry(key)
                .and_modify(|n| *n += copies)
                .or_insert(copies);
        }
    }

    w_dict
}

fn solve_part_one(cards: &Vec<Card>) -> usize {
    cards.iter().map(|c| c.get_score()).sum::<usize>()
}

fn solve_part_two(cards: &Vec<Card>) -> usize {
    let w_dict = create_winner_dict(cards);
    w_dict.iter().map(|(_key, value)| *value).sum()
}

fn main() {
    let cards = parse_input();
    let part_one = solve_part_one(&cards);
    let part_two = solve_part_two(&cards);
    dbg!(part_one);
    dbg!(part_two);
}
