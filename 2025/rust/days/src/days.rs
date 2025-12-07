use std::fs::File;

pub mod day01;
pub mod day02;
pub mod day03;
pub mod day04;

pub fn execute(day: i16, file: &File, part: i16) -> Result<i64, &str> {
    let result = match day {
        1 => match part {
            1 => day01::part1(file),
            2 => day01::part2(file),
            _ => return Err("Invalid part"),
        },
        2 => match part {
            1 => day02::part1(file),
            2 => day02::part2(file),
            _ => return Err("Invalid part"),
        },
        3 => match part {
            1 => day03::part1(file),
            2 => day03::part2(file),
            _ => return Err("Invalid part"),
        },
        4 => match part {
            1 => day04::part1(file),
            2 => day04::part2(file),
            _ => return Err("Invalid part"),
        },
        _ => return Err("Invalid day"),
    };

    Ok(result)
}
