use std::{collections::HashMap, str::FromStr};

#[derive(Debug)]
struct Node {
    left: String,
    right: String,
}

#[derive(Debug)]
struct ParseError;

impl FromStr for Node {
    type Err = ParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: Vec<&str> = s.split_whitespace().collect();

        let left = parts[2][1..parts[2].len() - 1].to_string();
        let right = parts[3][0..parts[3].len() - 1].to_string();
        Ok(Node { left, right })
    }
}

fn create_node_hash(input: &str) -> HashMap<String, Node> {
    input
        .lines()
        .skip(2)
        .map(|line| {
            let parts: Vec<&str> = line.split_whitespace().collect();
            let node = line.parse().unwrap();
            (parts[0].to_string(), node)
        })
        .collect()
}

fn traverse_nodes(directions: String, nodes: HashMap<String, Node>, target: &str) -> usize {
    let mut count = 1;
    let mut directions_iterator = directions.chars();
    let mut next_node = nodes.get("AAA");

    loop {
        match next_node {
            Some(node) => {
                let next_dir = directions_iterator.next();
                match next_dir {
                    Some('L') => {
                        println!("Direction is to go Left");
                        if node.left.as_str() == target {
                            return count;
                        } else {
                            next_node = nodes.get(node.left.as_str());
                            count += 1;
                        }
                    }
                    Some('R') => {
                        println!("Direction is to go Right");
                        if node.right.as_str() == target {
                            return count;
                        } else {
                            next_node = nodes.get(node.right.as_str());
                            count += 1;
                        }
                    }
                    Some(_) => unreachable!(),
                    None => {
                        println!("Out of chars! resetting the directions");
                        directions_iterator = directions.chars()
                    }
                }
            }
            None => {
                println!("Did not find a next node");
                break;
            }
        }
    }
    count
}

fn solve_part_one(input: &str, target: &str) -> usize {
    let directions: String = input.lines().take(1).collect();
    let nodes = create_node_hash(input);
    traverse_nodes(directions, nodes, target)
}

fn solve_part_two() -> usize {
    0
}

fn main() {
    let input = include_str!("./input.txt");
    let part_one = solve_part_one(input, "ZZZ");
    dbg!(part_one);
}
