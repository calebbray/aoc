use itertools::Itertools;

fn solve_part_one() -> usize {}

fn parse_input(input: &str) -> usize {
    let parts = input.split("\n\n").collect_tuple().unwrap();
    let (
        seeds,
        seed_to_soil,
        soil_to_fertilier,
        fertilizer_to_water,
        water_to_light,
        light_to_temperature,
        temperature_to_humidity,
        humidity_to_location,
    ) = parts;
}

fn main() {
    let input = parse_input(include_str!("./sample.txt"));

    let part_one = solve_part_one();
    dbg!(part_one);
}
