use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line.expect("line exists");
    }

    1
}

pub fn part2(file: &File) -> i64 {
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line.expect("line exists");
    }

    1
}
