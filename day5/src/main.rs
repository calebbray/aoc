#[derive(Debug)]
struct Seed(usize);

#[derive(Debug)]
struct Map {
    from: String,
    to: String,
    map: Vec<Range>,
}

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
}

impl Map {
    fn new(map: &str) -> Self {
        let meta = map.lines().take(1).collect::<String>();
        let meta: Vec<&str> = meta.split_ascii_whitespace().collect::<Vec<&str>>()[0]
            .split('-')
            .collect();

        let ranges = map.lines().skip(1).collect::<Vec<&str>>();
        let ranges = ranges.iter().map(|l| Range::new(l)).collect();

        Self {
            from: meta[0].to_string(),
            to: meta[2].to_string(),
            map: ranges,
        }
    }

    fn map_to_next(&self, input: usize) -> usize {
        let found_map = self
            .map
            .iter()
            .find(|range| range.check_is_input_in_map(&input));
        if let Some(found) = found_map {
            let mapped = input - found.source + found.dest;
            return mapped;
        }

        input
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

fn parse_seeds(seeds: String) -> Vec<Seed> {
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

fn parse_input(input: &str) -> anyhow::Result<(Vec<Seed>, Vec<Map>)> {
    let seeds = parse_seeds(input.lines().take(1).collect());
    let maps = parse_maps(input);
    Ok((seeds, maps))
}

fn solve_part_one(seeds: &Vec<Seed>, maps: &Vec<Map>) -> usize {
    seeds.iter().map(|s| s.get_path(&maps)).min().unwrap()
}

fn main() -> anyhow::Result<()> {
    let (seeds, maps) = parse_input(include_str!("./input.txt"))?;

    let part_one = solve_part_one(&seeds, &maps);
    dbg!(part_one);

    Ok(())
}
