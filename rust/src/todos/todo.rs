use chrono::{DateTime, Local};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
pub struct ToDo {
    content: String,
    complete: bool,
    data: DateTime<Local>,
}

impl ToDo {
    pub fn new(content: String) -> Self {
        Self {
            content,
            complete: false,
            data: chrono::Local::now(),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        let mut target = Vec::new();

        {
            let mut writer = csv::Writer::from_writer(&mut target);

            writer.serialize(ToDo::new("content".to_owned()));
        }
        let result: Vec<u8> = Vec::new();
        assert_eq!(target, result);
    }
}

