use std::{
    fs::File,
    io::{BufRead, BufReader},
};

use regex::Regex;

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        let ranges: Vec<&str> = line.split(",").collect();

        for range in ranges {
            let re = Regex::new(r"(\d+)\-(\d+)").unwrap();

            let Some(captures) = re.captures(range) else {
                panic!();
            };

            let start: i64 = (&captures[1]).parse().unwrap();
            let end: i64 = (&captures[2]).parse().unwrap();

            for i in start..=end {
                let s = i.to_string();
                let len = s.len();

                if len % 2 != 0 {
                    continue;
                }

                if s[..len / 2] == s[len / 2..] {
                    result += i;
                }
            }
        }

        break;
    }

    result
}

pub fn part2(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        let ranges: Vec<&str> = line.split(",").collect();

        for range in ranges {
            let re = Regex::new(r"(\d+)\-(\d+)").unwrap();

            let Some(captures) = re.captures(range) else {
                panic!();
            };

            let start: i64 = (&captures[1]).parse().unwrap();
            let end: i64 = (&captures[2]).parse().unwrap();

            for i in start..=end {
                let current = i.to_string();
                let len = current.len();

                for j in 1..=(len / 2) {
                    if len % j != 0 {
                        continue;
                    }

                    let mut current_start = 0;
                    let mut current_end = j;

                    let mut groups = vec![];

                    while current_end < len + 1 {
                        groups.push(&current[current_start..current_end]);
                        current_start = current_end;
                        current_end += j;
                    }

                    let all_equal = groups.iter().all(|&group| group == groups[0]);

                    if all_equal {
                        result += i;
                        break;
                    }
                }
            }
        }

        break;
    }

    result
}
