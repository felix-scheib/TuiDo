use clap::{Parser, Subcommand};

/// CLI ToDo manager
#[derive(Parser)]
#[command(version, about)]
pub struct Args {
    #[command(subcommand)]
    pub command: Commands,
}

#[derive(Subcommand)]
pub enum Commands {
    /// Add a ToDo
    Add {
        /// Content of the ToDo
        content: String,
    },
    /// Complete a ToDo
    Complete {
        // Number of the Todo
        number: usize,
    },
    /// Delete a ToDo
    Delete {
        // Number of the Todo
        number: usize,
    },
    /// List all ToDos
    List {
        /// Display all ToDos
        #[clap(long, short, action)]
        all: bool,

        /// Display complete ToDos
        #[clap(long, short, action)]
        complete: bool,
    },
}
