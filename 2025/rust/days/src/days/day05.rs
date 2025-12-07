use std::{
    fs::File,
    io::{BufRead, BufReader},
};

use regex::Regex;

#[derive(Debug)]
struct Range {
    start: u64,
    end: u64,
}

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    let mut ranges = Vec::new();
    let mut ingredients = Vec::new();

    let mut read_ingredients = false;

    for line in reader.lines() {
        let line = line.unwrap();

        if line == "" {
            read_ingredients = true;
            continue;
        }

        if !read_ingredients {
            let re = Regex::new(r"(\d+)\-(\d+)").unwrap();
            let captures = re.captures(&line).unwrap();

            let start: u64 = (&captures[1]).parse().unwrap();
            let end: u64 = (&captures[2]).parse().unwrap();

            ranges.push(Range { start, end });

            continue;
        }

        let ingredient: u64 = line.parse().unwrap();
        ingredients.push(ingredient);
    }

    for ingredient in ingredients.iter() {
        for range in ranges.iter() {
            if *ingredient >= range.start && *ingredient <= range.end {
                result += 1;
                break;
            }
        }
    }

    result
}

pub fn part2(file: &File) -> i64 {
    1
}
