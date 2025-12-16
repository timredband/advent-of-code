use regex::Regex;
use std::{
    collections::HashSet,
    fs::File,
    io::{BufRead, BufReader},
};

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
struct Position {
    x: i64,
    y: i64,
    z: i64,
}

#[derive(Debug)]
struct Connection<'a> {
    a: &'a Position,
    b: &'a Position,
    distance: f64,
}

impl<'a> Connection<'a> {
    fn new(a: &'a Position, b: &'a Position) -> Self {
        let radicand = ((a.x - b.x).pow(2)) + ((a.y - b.y).pow(2)) + ((a.z - b.z).pow(2));
        let distance = (radicand as f64).sqrt();
        Self { a, b, distance }
    }
}

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);
    let re = Regex::new(r"(\d+),(\d+),(\d+)").unwrap();

    let mut positions = Vec::new();
    let mut circuits: Vec<HashSet<Position>> = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();

        let captures = re.captures(&line).unwrap();

        let x: i64 = (&captures[1]).parse().unwrap();
        let y: i64 = (&captures[2]).parse().unwrap();
        let z: i64 = (&captures[3]).parse().unwrap();

        let position = Position { x, y, z };
        positions.push(position);

        let mut set = HashSet::new();
        set.insert(position);

        circuits.push(set);
    }

    let mut connections = Vec::new();

    for i in 0..positions.len() {
        for j in i + 1..positions.len() {
            connections.push(Connection::new(&positions[i], &positions[j]));
        }
    }

    connections.sort_by(|a, b| a.distance.total_cmp(&b.distance));

    for index in 0..1000 {
        let connection = &connections[index];

        let mut circuit = HashSet::new();
        let mut indices: Vec<usize> = Vec::new();

        for i in 0..circuits.len() {
            if circuits[i].contains(connection.a) || circuits[i].contains(connection.b) {
                circuits[i].insert(*connection.a);
                circuits[i].insert(*connection.b);
                circuit.extend(circuits[i].clone());
                indices.push(i);
            }
        }

        for i in (0..indices.len()).rev() {
            circuits.swap_remove(indices[i]);
        }

        circuits.push(circuit);
        circuits.retain(|circuit| circuit.len() != 0);
    }

    circuits.sort_by(|a, b| b.len().cmp(&a.len()));

    // for circuit in &circuits {
    //     println!("{} | {:?}", circuit.len(), circuit);
    // }
    //
    // println!("{}", circuits.len());

    let result = &circuits[0].len() * circuits[1].len() * circuits[2].len();
    result as i64
}

pub fn part2(file: &File) -> i64 {
    1
}
