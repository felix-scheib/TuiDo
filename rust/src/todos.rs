use std::{fs::File, io, path::PathBuf};

use todo::ToDo;

pub mod todo;

pub struct ToDos {
    path: PathBuf,
    todos: Vec<ToDo>,
}

impl ToDos {
    pub fn new(path: &str) -> Self {
        let path = PathBuf::from(path);
        let todos = Self::load_from_file(&path);

        Self { path, todos }
    }

    fn load_from_file(path: &PathBuf) -> Vec<ToDo> {
        match File::open(&path) {
            Ok(file) => {
                let _ = Self::from_csv(&file);
            }
            Err(e) => {
                eprintln!("Error while reading File: {:#?}", e);
                eprintln!("Creating new File: {:#?}", path);

                if let Err(e) = File::create(path) {
                    panic!("Error while creating File: {:#?}", e);
                }
            }
        }
        vec![ToDo::new("content".to_owned())]
    }

    fn from_csv(file: &File) -> Vec<ToDo> {
        let mut reader = csv::Reader::from_reader(file);

        for result in reader.deserialize() {
            // Notice that we need to provide a type hint for automatic
            // deserialization.
            let record: ToDo = result.unwrap();
            println!("{:?}", record);
        }
        Vec::new()
    }

    fn to_csv(&self) {}
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
