use clap::Parser;
use todos::ToDos;

mod cli;
mod todos;

const PATH: &str = "./todos.csv";

fn main() {
    let cli = cli::Args::parse();

    let mut todos = ToDos::new(PATH);

    match &cli.command {
        cli::Commands::Add { content } => todos.add_todo(content),
        cli::Commands::Complete { number } => todos.complete_todo(*number),
        cli::Commands::Delete { number } => todos.delete_todo(*number),
        cli::Commands::List { all, complete } => {
            if *all {
                todos.to_table(|_| true).printstd();
            } else if *complete {
                todos.to_table(|todo| todo.get_complete()).printstd();
            } else {
                todos.to_table(|todo| !todo.get_complete()).printstd();
            }
        }
    }
}
