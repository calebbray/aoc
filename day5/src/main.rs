use std::collections::BTreeSet;

#[derive(Debug)]
struct Seed(usize);

#[derive(Debug)]
struct Seed2 {
    start: usize,
    length: usize,
}

#[derive(Debug)]
struct Map(Vec<Range>);

#[derive(Debug)]
struct Range {
    dest: usize,
    source: usize,
    length: usize,
}

impl Range {
    fn new(range: &str) -> Self {
        let parts: Vec<usize> = range
            .split_ascii_whitespace()
            .map(|n| n.parse().unwrap())
            .collect();
        Self {
            dest: parts[0],
            source: parts[1],
            length: parts[2],
        }
    }

    fn check_is_input_in_map(&self, input: &usize) -> bool {
        let range = self.source..self.source + self.length;
        range.contains(input)
    }

    fn convert(&self, input: &usize) -> Option<usize> {
        if self.check_is_input_in_map(input) {
            return Some(input - self.source + self.dest);
        }
        None
    }
}

impl Map {
    fn new(map: &str) -> Self {
        let ranges = map.lines().skip(1).collect::<Vec<&str>>();
        let ranges = ranges.iter().map(|l| Range::new(l)).collect();

        Self(ranges)
    }

    fn map_to_next(&self, input: usize) -> usize {
        let found_map = self
            .0
            .iter()
            .find(|range| range.check_is_input_in_map(&input));
        if let Some(found) = found_map {
            let mapped = input - found.source + found.dest;
            return mapped;
        }

        input
    }

    fn convert(&self, input: &usize) -> usize {
        match self.0.iter().map(|r| r.convert(input)).find_map(|r| r) {
            Some(dest) => dest,
            None => *input,
        }
    }

    fn convert_range(&self, seed_range: &Seed2) -> Vec<Seed2> {
        let mut chunks = BTreeSet::new();
        let seed_range_end = seed_range.start + seed_range.length;

        for range in &self.0 {
            let range_end = range.source + range.length;

            if seed_range_end < range.source || seed_range.start > range_end {
                continue;
            }

            if range.source > seed_range.start {
                chunks.insert(range.source);
            }

            if range_end < seed_range_end {
                chunks.insert(range_end);
            }
        }

        chunks.insert(seed_range_end);

        let mut output = Vec::new();
        let mut current = seed_range.start;

        for position in chunks {
            output.push(Seed2 {
                start: self.convert(&current),
                length: position - current,
            });
            current = position;
        }

        output
    }
}

impl Seed {
    fn get_path(&self, maps: &Vec<Map>) -> usize {
        let Seed(n) = self;
        let mut last_path = *n;
        for map in maps {
            last_path = map.map_to_next(last_path);
        }
        last_path
    }
}

fn parse_seed_ranges(seeds: &String) -> Vec<Seed2> {
    seeds
        .split_ascii_whitespace()
        .skip(1)
        .collect::<Vec<&str>>()
        .chunks(2)
        .map(|range| {
            let start = range[0].parse::<usize>().unwrap();
            let length = range[1].parse::<usize>().unwrap();
            Seed2 { start, length }
        })
        .collect()
}

fn parse_seeds(seeds: &String) -> Vec<Seed> {
    seeds
        .split_ascii_whitespace()
        .skip(1)
        .collect::<Vec<&str>>()
        .iter()
        .map(|s| Seed(s.parse().unwrap()))
        .collect()
}

fn parse_maps(input: &str) -> Vec<Map> {
    let map_sections: Vec<&str> = input.split("\n\n").skip(1).collect();
    map_sections
        .iter()
        .map(|section| Map::new(section))
        .collect()
}

fn parse_input(input: &str) -> anyhow::Result<(Vec<Seed>, Vec<Map>, Vec<Seed2>)> {
    let seed_line: String = input.lines().take(1).collect();
    let seeds = parse_seeds(&seed_line);
    let maps = parse_maps(input);
    let seed_ranges = parse_seed_ranges(&seed_line);
    Ok((seeds, maps, seed_ranges))
}

fn solve_part_one(seeds: &Vec<Seed>, maps: &Vec<Map>) -> usize {
    seeds.iter().map(|s| s.get_path(&maps)).min().unwrap()
}

fn solve_part_two(seeds: Vec<Seed2>, maps: &Vec<Map>) -> usize {
    let mut next = Vec::new();
    let mut current = seeds;

    for map in maps {
        for range in current {
            next.extend(map.convert_range(&range));
        }

        current = next;
        next = Vec::new();
    }

    current.iter().map(|range| range.start).min().unwrap()
}

fn main() -> anyhow::Result<()> {
    let (seeds, maps, seed_ranges) = parse_input(include_str!("./input.txt"))?;

    let part_one = solve_part_one(&seeds, &maps);
    let part_two = solve_part_two(seed_ranges, &maps);
    dbg!(part_one);
    dbg!(part_two);

    Ok(())
}
