use std::{
    cmp,
    fs::File,
    io::{BufRead, BufReader},
};

use regex::Regex;

#[derive(Clone, Copy, Debug, PartialEq)]
struct Range {
    start: u64,
    end: u64,
}

// TODO: look at Cow for no copying
fn merge(a: &Range, b: &Range) -> Vec<Range> {
    if a == b {
        return vec![*a];
    }

    if !(cmp::max(a.start, b.start) <= cmp::min(a.end, b.end)) {
        return vec![*a, *b];
    }

    let range = Range {
        start: cmp::min(a.start, b.start),
        end: cmp::max(a.end, b.end),
    };

    return vec![range];
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

    for ingredient in &ingredients {
        for range in &ranges {
            if *ingredient >= range.start && *ingredient <= range.end {
                result += 1;
                break;
            }
        }
    }

    result
}

pub fn part2(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    let mut ranges = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();

        if line == "" {
            break;
        }

        let re = Regex::new(r"(\d+)\-(\d+)").unwrap();
        let captures = re.captures(&line).unwrap();

        let start: u64 = (&captures[1]).parse().unwrap();
        let end: u64 = (&captures[2]).parse().unwrap();

        ranges.push(Range { start, end });
    }

    loop {
        let len = ranges.len();

        for i in 0..len - 1 {
            if i >= ranges.len() - 1 {
                continue;
            }

            let mut merged = merge(&ranges[i], &ranges[i + 1]);

            if merged.len() == 1 {
                ranges.append(&mut merged);
                ranges.swap_remove(i);
                ranges.swap_remove(i + 1);
            }
        }

        ranges.sort_by(|a, b| a.start.cmp(&b.start));

        if len == ranges.len() {
            break;
        }
    }

    for range in &ranges {
        result += range.end - range.start + 1;
    }

    result.try_into().unwrap()
}
