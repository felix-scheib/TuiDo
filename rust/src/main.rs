use clap::Parser;
use todos::ToDos;

mod cli;
mod todos;

const PATH: &str = "./todos.csv";

fn main() {
    let _cli = cli::Args::parse();

    let _ = ToDos::new(PATH);

    /*
    match &cli.command {
        cli::Commands::Add { content } => todo!(),
        cli::Commands::Complete { number } => todo!(),
        cli::Commands::Delete { number } => todo!(),
        cli::Commands::List { all, complete } => todo!(),
    }
    */
}
