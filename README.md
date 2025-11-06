# Todo CLI

A simple command-line To-Do list manager that allows you to add,
complete, delete, and list tasks, written as a exercise

## Features

* **Add tasks**: Add a task to your ToDo list via command-line arguments or by providing multiple lines via STDIN.
* **Complete tasks**: Mark tasks as completed by specifying their index.
* **Delete tasks**: Remove tasks from the list by specifying their index.
* **Export to file**: Write the list of tasks to a JSON file.
* **List tasks**: Display all tasks in your ToDo list.
* **Pending tasks**: Show only incomplete tasks.
* **Verbose**: Show a detailed list of tasks, including additional information.

## Usage

```bash
$ todo -h
Usage information:
  -add
        Add task to the ToDo list via Args or STDIN
        Multiple lines for multiple tasks
  -complete int
        Item to be completed (provide the task index)
  -delete int
        Item to be deleted (provide the task index)
  -file string
        Write the todo items to a file
  -list
        List all tasks
  -pending
        Show incomplete tasks
  -verbose
        Show a verbose list
```

## Showcase
### Add tasks
![Add tasks](/gifs/add.gif)
### Complete tasks
![Complete tasks](/gifs/complete.gif)
### Track tasks
![Track tasks](/gifs/verbose.gif)

## Installation

1. Manual:

   ```bash
   git clone https://github.com/ukiran03/todo.git
   cd todo
   go build ./cmd/todo
   ```

2. Go install:
   ```bash
   go install github.com/ukiran03/todo/cmd/todo@latest
   ```

## Contributing

It was just a project I wrote as an exercise, but please don't hesitate to correct me.
