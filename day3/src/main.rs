fn parse_input(input: &str) -> (Vec<Symbol>, Vec<Number>) {
    let mut symbols: Vec<Symbol> = vec![];
    let mut numbers: Vec<Number> = vec![];
    input.lines().enumerate().for_each(|(row, line)| {
        let (mut syms, mut nums) = parse_line(line, row as isize);
        symbols.append(&mut syms);
        numbers.append(&mut nums);
    });
    (symbols, numbers)
}

fn parse_line(mut line: &str, row: isize) -> (Vec<Symbol>, Vec<Number>) {
    let mut symbols: Vec<Symbol> = vec![];
    let mut numbers: Vec<Number> = vec![];
    let mut col = 0;
    while !line.is_empty() {
        let next_char = line.chars().next();
        if let Some(c) = next_char {
            if c.is_symbol() {
                symbols.push(Symbol {
                    location: Point { x: col, y: row },
                    symbol: c,
                });
                col += 1;
                line = &line[1..];
            } else if c.is_digit(10) {
                let num_string = parse_number_from_line(line);
                let parsed_number = num_string
                    .parse::<isize>()
                    .expect("This should have been a number");
                numbers.push(Number {
                    value: parsed_number,
                    start: Point { x: col, y: row },
                    end: Point {
                        x: col + num_string.len() as isize - 1,
                        y: row,
                    },
                });
                col += num_string.len() as isize;
                line = &line[num_string.len()..];
            } else {
                line = &line[1..];
                col += 1;
            }
        }
    }
    (symbols, numbers)
}

fn parse_number_from_line(input: &str) -> String {
    input
        .chars()
        .take_while(|c| c.is_digit(10))
        .collect::<String>()
}

fn solve_part_one(input: &str) -> isize {
    let (symbols, numbers) = parse_input(input);
    let score = numbers
        .iter()
        .filter_map(|num| {
            if num.has_adjacent_symbol(&symbols) {
                return Some(num.value);
            }
            None
        })
        .sum::<isize>();

    score
}

fn solve_part_two(input: &str) -> isize {
    let (symbols, numbers) = parse_input(input);
    let gear_symbols: Vec<_> = symbols.iter().filter(|s| s.symbol == '*').collect();

    let gears = gear_symbols
        .iter()
        .filter_map(|s| {
            let touching = s.find_touching_numbers(&numbers);
            if touching.len() >= 2 {
                return Some(touching);
            }
            None
        })
        .collect::<Vec<Vec<isize>>>();

    let gear_ratios: isize = gears.iter().map(|g| g.iter().product::<isize>()).sum();
    gear_ratios
}

trait IsSymbol {
    fn is_symbol(&self) -> bool;
}

impl IsSymbol for char {
    fn is_symbol(&self) -> bool {
        self != &'.' && !self.is_digit(10)
    }
}

#[derive(Debug, Clone, Copy)]
struct Point {
    x: isize,
    y: isize,
}

#[derive(Debug)]
struct Symbol {
    location: Point,
    symbol: char,
}

impl Symbol {
    fn find_left_number(&self, numbers: &Vec<Number>) -> Option<isize> {
        let found = numbers.iter().find(|num| {
            let is_same_y = num.end.y == self.location.y;
            let is_touching_left = num.start.x - self.location.x == 1;
            is_same_y && is_touching_left
        });
        if let Some(f) = found {
            return Some(f.value);
        }
        None
    }

    fn find_right_number(&self, numbers: &Vec<Number>) -> Option<isize> {
        let found = numbers.iter().find(|num| {
            let is_same_y = num.end.y == self.location.y;
            let is_touching_right = self.location.x - num.end.x == 1;
            is_same_y && is_touching_right
        });
        if let Some(f) = found {
            return Some(f.value);
        }
        None
    }

    fn find_upper_number(&self, numbers: &Vec<Number>) -> Vec<isize> {
        numbers
            .iter()
            .filter_map(|num| {
                let is_above = self.location.y - num.start.y == 1;
                let range = num.start.x - 1..=num.end.x + 1;
                let is_overlapping = range.contains(&self.location.x);
                if is_above && is_overlapping {
                    return Some(num.value);
                }
                None
            })
            .collect()
    }

    fn find_lower_number(&self, numbers: &Vec<Number>) -> Vec<isize> {
        numbers
            .iter()
            .filter_map(|num| {
                let is_below = num.start.y - self.location.y == 1;
                let range = num.start.x - 1..=num.end.x + 1;
                let is_overlapping = range.contains(&self.location.x);
                if is_below && is_overlapping {
                    return Some(num.value);
                }
                None
            })
            .collect()
    }

    fn find_touching_numbers(&self, numbers: &Vec<Number>) -> Vec<isize> {
        let mut touching = vec![];
        let left_side = self.find_left_number(numbers);
        let right_side = self.find_right_number(numbers);
        let mut upper_numbers = self.find_upper_number(numbers);
        let mut lower_numbers = self.find_lower_number(numbers);

        if let Some(left) = left_side {
            touching.push(left);
        }
        if let Some(right) = right_side {
            touching.push(right);
        }

        touching.append(&mut upper_numbers);
        touching.append(&mut lower_numbers);

        touching
    }
}

#[derive(Debug, Clone, Copy)]
struct Number {
    value: isize,
    start: Point,
    end: Point,
}

impl Number {
    fn has_adjacent_symbol(&self, symbols: &Vec<Symbol>) -> bool {
        symbols
            .iter()
            .find(|symbol| {
                let is_right = symbol.location.x - self.end.x == 1;
                let is_left = self.start.x - symbol.location.x == 1;
                let is_above = self.start.y - symbol.location.y == 1;
                let is_below = symbol.location.y - self.end.y == 1;

                let is_top_right_corner = is_right && is_above;
                let is_top_left_corner = is_left && is_above;
                let is_bottom_right_corner = is_right && is_below;
                let is_bottom_left_corner = is_left && is_below;

                let is_between =
                    self.start.x <= symbol.location.x && symbol.location.x <= self.end.x;

                let is_same_line = self.start.y == symbol.location.y;

                let is_between_above = is_above && is_between;
                let is_between_below = is_below && is_between;

                let is_right_same_line = is_same_line && is_right;
                let is_left_same_line = is_same_line && is_left;

                is_top_left_corner
                    || is_top_right_corner
                    || is_bottom_left_corner
                    || is_bottom_right_corner
                    || is_between_above
                    || is_between_below
                    || is_right_same_line
                    || is_left_same_line
            })
            .is_some()
    }
}

fn main() {
    let input = include_str!("./input.txt");
    let part_one = solve_part_one(input);
    let part_two = solve_part_two(input);

    dbg!(part_one);
    dbg!(part_two);
}
