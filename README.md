
# TodoCLI

This is a command line interface that allows for creating todos.


## Installation
Go is a Prerequisite for build

Clone the project and build the project

```bash
git clone https://github.com/robertgouveia/TodoCLI
cd TodoCLI
sudo go build -o /usr/local/bin/tasks
```

## Uninstalling

Run the below
```bash
sudo rm /usr/local/bin/tasks
```
    
## CLI Commands

#### Help

```bash
tasks -h
```

#### List Todos

```bash
tasks list
```

| ID    | Title    | Done      |
| :---- | :------- | :-------- |
| `int` | `string` | `boolean` |

#### List All Todos (Includes completed)

```bash
tasks list -a
```

| ID    | Title    | Done      |
| :---- | :------- | :-------- |
| `int` | `string` | `boolean` |

#### Add Todo

```bash
tasks add "{title}"
```

#### Complete Todo

```bash
tasks complete {id}
```

#### Remove Todo

```bash
tasks remove {id}
```
## Features

- CRUD Operations
- CLI Standards


## Roadmap

- Adding CSV Storage
- Adding Logs
- Adding Due Dates
