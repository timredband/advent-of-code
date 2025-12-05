use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;
    let mut position = 50;

    for line in reader.lines() {
        let line = line.expect("line exists");
        let direction = &line[0..1];
        let clicks: i32 = (&line[1..]).parse().expect("valid integer");

        position = match direction {
            "L" => position - clicks,
            "R" => position + clicks,
            _ => panic!("invalid direction"),
        }
        .rem_euclid(100);

        if position == 0 {
            result += 1;
        }
    }

    result
}

pub fn part2(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;
    let mut position: i32 = 50;

    for line in reader.lines() {
        let line = line.expect("line exists");
        let direction = &line[0..1];
        let clicks: i32 = (&line[1..]).parse().expect("valid integer");

        if direction == "L" {
            for i in 0..clicks {
                position = position - 1;
                position = position.rem_euclid(100);

                if position == 0 {
                    result += 1;
                    let remaining = clicks - i - 1;
                    result += remaining / 100;
                    position = position - remaining;
                    break;
                }
            }
        }

        if direction == "R" {
            for i in 0..clicks {
                position = position + 1;
                position = position.rem_euclid(100);

                if position == 0 {
                    result += 1;
                    let remaining = clicks - i - 1;
                    result += remaining / 100;
                    position = position + remaining;
                    break;
                }
            }
        }
    }

    result.into()
}
