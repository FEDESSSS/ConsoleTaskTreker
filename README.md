# Task Tracker CLI

A simple command-line tool for managing tasks, written in Go.

This project is the solution for [This Project](https://roadmap.sh/projects/task-tracker)

## Features
- Add, update, delete tasks
- Mark tasks as todo, in-progress, or done
- List all tasks
- Persistent storage in `task.json`


### Available Commands

- `add <task_description>`: Add a new task with the given description
- `list`: List all tasks
- `update <task_id> <new_description>`: Update the description of a task
- `delete <task_id>`: Delete a task by its ID
- `mark-todo <task_id>`: Mark a task as "todo"
- `mark-in-progress <task_id>`: Mark a task as "in-progress"
- `mark-done <task_id>`: Mark a task as "done"

## Examples

Add a new task:
```sh
task-tracker-cli add "Write documentation"
```

List all tasks:
```sh
task-tracker-cli list
```

Update a task:
```sh
task-tracker-cli update 1 "Write detailed documentation"
```

Delete a task:
```sh
task-tracker-cli delete 1
```

Mark a task as in progress:
```sh
task-tracker-cli mark-in-progress 1
```

Mark a task as done:
```sh
task-tracker-cli mark-done 1
```

Mark a task as done (revert):
```sh
task-tracker-cli mark-todo 1
```

## License
See [LICENSE](../LICENSE).
