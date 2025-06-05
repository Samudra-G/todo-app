# 📝 Todo CLI App in Go

A clean, modular command-line **Todo application** written in Go — featuring task creation, completion, deletion, persistent storage (JSON), and a beautiful terminal UI using `simpletable`.

![Go Version](https://img.shields.io/badge/Go-1.21+-blue?style=flat-square)
![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)
![Platform](https://img.shields.io/badge/platform-terminal-lightgrey)

---

## 📦 Features

- ✅ Add new tasks via CLI or piped input
- 📋 List all tasks in a styled table (with `CreatedAt`, `CompletedAt`, status)
- ☑️ Mark tasks as complete
- ❌ Delete tasks by index
- 💾 Persistent task storage in `.todo.json`
- 🎨 Colored output using ANSI escape codes
- 🧱 Modular project structure with `internal` packages
- 🛠 Simple `Makefile` for easy builds

---

## 🚀 Getting Started

### 🔧 Prerequisites

- [Go 1.21+](https://golang.org/dl/)

### 📥 Installation

```bash
git clone https://github.com/yourusername/todo-cli-go.git
cd todo-cli-go
make build    # or: go build -o todo ./cmd/todo
```

---

## 🛠 Usage

```bash
./todo [flags] [task description]
```

### 🎯 Flags

| Flag         | Description                                 | Example                                  |
|--------------|---------------------------------------------|------------------------------------------|
| `-add`       | Add a new task                              | `./todo -add "Buy groceries"`            |
| `-list`      | List all tasks                              | `./todo -list`                           |
| `-complete`  | Mark task complete by index (1-based)       | `./todo -complete 2`                     |
| `-del`       | Delete task by index                        | `./todo -del 3`                          |

### 💡 Add via Pipe

```bash
echo "Read 'The Go Programming Language'" | ./todo -add
```

---

## 🗃 Example Output

```bash
./todo -list
```

```
╔═══╤══════════════════════════════════════╤═══════╤═══════════════════════════╤═══════════════════════════╗
║ # │                 Task                 │ Done? │             CreatedAt     │          CompletedAt      ║
╟───┼──────────────────────────────────────┼───────┼────────────────────────────┼──────────────────────────╢
║ 1 │ ✅ Finish building the CLI app       │ Yes   │ 2025-06-05T13:11:59+05:30 │ 2025-06-05T13:13:46+05:30 ║
║ 2 │ Write documentation                 │ No    │ 2025-06-05T13:13:26+05:30 │ 0001-01-01T00:00:00Z       ║
╚═══╧══════════════════════════════════════╧═══════╧═══════════════════════════╧═══════════════════════════╝
```

---

## 📁 Project Structure

```bash
.
├── cmd/
│   └── todo/
│       └── main.go        # CLI entrypoint
├── internal/
│   └── todo/
│       ├── todo.go        # Task logic (add, complete, delete, load, store)
│       └── colors.go      # Color utilities
├── Makefile
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## 🧪 Running Tests

_Tests can be added soon for core task logic in `internal/todo`._

---

## 📜 License

Licensed under the [Apache 2.0 License](./LICENSE).

---

## 🙋‍♂️ Author

**Samudra Mukhar Goswami**

- 💼 [LinkedIn](https://linkedin.com/in/samudramukhar)
- ✍️ [Medium Blog](https://medium.com/@samudramukhar)
- 💻 [GitHub](https://github.com/Samudra-G)

---

## 🌟 Star This Repo

If you found this project useful or educational, a ⭐️ on [GitHub](https://github.com/Samudra-G/todo-cli-go) would be awesome!

---

## 📌 TODOs (for future improvement)

- [ ] Add unit tests for core functions
- [ ] Add task priority and tags
- [ ] Support for editing tasks
- [ ] Auto-backup and config options

---