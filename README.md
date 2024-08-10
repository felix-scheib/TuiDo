# TuiDo

A simple CLI ToDo application written in Go and Rust.

- `tuido add "Content"` to add a new ToDo
- `tuido complete NUMBER` to complete a ToDo
- `tuido delete NUMBER` to delete a ToDo
- `tuido list` to list all ToDos

## Go
| | Library |
|- | - |
| CLI | [Cobra](https://cobra.dev/) |
| CSV | [GoCSV](https://github.com/gocarina/gocsv) |
| tables | [go-pretty](https://github.com/jedib0t/go-pretty) |
| time formatting | [timediff](https://github.com/mergestat/timediff) |

## Rust
| | Library |
|- | - |
| CLI | [clap](https://github.com/clap-rs/clap) |
| CSV | [csv](https://github.com/BurntSushi/rust-csv) |
| tables | [prettytables-rs](https://github.com/phsym/prettytable-rs) |
| time formatting | [Chrono Humanize](https://gitlab.com/imp/chrono-humanize-rs) |

