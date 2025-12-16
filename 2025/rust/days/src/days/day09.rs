use regex::Regex;
use std::{
    cmp,
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);
    let re = Regex::new(r"(\d+),(\d+)").unwrap();

    let mut positions = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        let captures = re.captures(&line).unwrap();

        let x: i64 = (&captures[2]).parse().unwrap();
        let y: i64 = (&captures[1]).parse().unwrap();

        positions.push((x, y));
    }

    let mut result = 0;

    for i in 0..positions.len() - 1 {
        for j in i + 1..positions.len() - 1 {
            let x1 = positions[i].0;
            let y1 = positions[i].1;
            let x2 = positions[j].0;
            let y2 = positions[j].1;

            let area = (x1.abs_diff(x2) + 1) * (y1.abs_diff(y2) + 1);
            result = cmp::max(result, area);
        }
    }

    result as i64
}

pub fn part2(file: &File) -> i64 {
    1
}
