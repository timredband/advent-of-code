use std::fs::File;

pub mod day01;
pub mod day02;

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
        _ => return Err("Invalid day"),
    };

    Ok(result)
}
