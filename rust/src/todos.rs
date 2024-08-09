use std::{fs::File, path::PathBuf};

use prettytable::Table;
use todo::ToDo;

pub mod todo;

pub struct ToDos {
    path: PathBuf,
    todos: Vec<ToDo>,
    number: usize,
}

impl ToDos {
    pub fn new(path: &str) -> Self {
        let path = PathBuf::from(path);
        let (todos, number) = Self::load_from_file(&path);

        Self {
            path,
            todos,
            number,
        }
    }

    fn load_from_file(path: &PathBuf) -> (Vec<ToDo>, usize) {
        match File::open(&path) {
            Ok(file) => {
                return Self::from_csv(&file);
            }
            Err(e) => {
                eprintln!("Error while reading File: {:#?}", e);
                eprintln!("Creating new File: {:#?}", path);

                if let Err(e) = File::create(path) {
                    panic!("Error while creating File: {:#?}", e);
                }

                return (Vec::new(), 1);
            }
        }
    }

    fn from_csv(file: &File) -> (Vec<ToDo>, usize) {
        let mut reader = csv::Reader::from_reader(file);

        let mut todos = Vec::<ToDo>::new();

        for result in reader.deserialize() {
            match result {
                Ok(t) => todos.push(t),
                Err(e) => eprintln!("Failed to parse ToDo: {:#?}", e),
            }
        }

        let number = match todos.iter().max_by_key(|t| t.get_number()) {
            Some(t) => t.get_number() + 1,
            None => 1,
        };

        (todos, number)
    }

    pub fn add_todo(&mut self, content: &String) {
        self.todos.push(ToDo::new(self.number, content));
    }

    pub fn complete_todo(&mut self, number: usize) {
        if let Some(todo) = self.todos.iter_mut().find(|t| t.get_number() == number) {
            todo.complete()
        }
    }

    pub fn delete_todo(&mut self, number: usize) {
        if let Some(index) = self.todos.iter().position(|t| t.get_number() == number) {
            self.todos.remove(index);

        }
    }

    pub fn to_table(&self, filter: fn(todo: &&ToDo) -> bool) -> Table {
        let mut table = Table::new();

        table.set_titles(ToDo::titles());

        for todo in self.todos.iter().filter(filter) {
            table.add_row(todo.row());
        }

        table
    }
}

impl Drop for ToDos {
    fn drop(&mut self) {
        let mut writer = csv::Writer::from_path(&self.path).unwrap();

        for todo in &self.todos {
            writer.serialize(todo).unwrap();
        }

        writer.flush().unwrap();
    }
}
