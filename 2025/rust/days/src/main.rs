use std::fs::File;

use clap::Parser;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long)]
    day: String,

    #[arg(short, long)]
    file: String,

    #[arg(short, long)]
    part: String,
}

pub mod days;

fn main() {
    let args = Args::parse();

    let day: i16 = args.day.parse().expect("day is not an integer");
    let file = File::open(args.file).expect("File should exist");
    let part: i16 = args.part.parse().expect("part is not an integer");

    let result = days::execute(day, &file, part).expect("result exists");
    println!("{result:?}")
}
