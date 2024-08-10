use chrono::{DateTime, Local};
use prettytable::{row, Row};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize)]
pub struct ToDo {
    number: usize,
    content: String,
    complete: bool,
    data: DateTime<Local>,
}

impl ToDo {
    pub fn new(number: usize, content: &String) -> Self {
        Self {
            number,
            content: content.clone(),
            complete: false,
            data: chrono::Local::now(),
        }
    }

    pub fn get_number(&self) -> usize {
        self.number
    }

    pub fn get_complete(&self) -> bool {
        self.complete
    }

    pub fn complete(&mut self) {
        self.complete = true
    }

    pub fn titles() -> Row {
        row!["Number", "Content", "Complete", "Date"]
    }

    pub fn row(&self) -> Row {
        let complete = match self.complete {
            true => "\u{2714}",
            false => "\u{274C}",
        };


        row![
            self.number,
            self.content,
            complete,
            chrono_humanize::HumanTime::from(self.data)
        ]
    }
}
