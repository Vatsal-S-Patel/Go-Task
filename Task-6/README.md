# Go Tasks

## Task-6 : Book API using mux and gorm

In this task, REST APIs are made for books, by using that CRUD operations can perform for books.

In this task, i used gorm as ORM and mux as router to handle HTTP requests.

### Instructions

1. For given book struct, create REST APIs to perform CRUD operations for books.
2. Create modular project structure for easy to maintain and readable code.

### API endpoints

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| GET       | /books       | Get all books |
| GET       | /book/{id}   | Get one book by id|
| POST      | /book        | Create book   |
| PUT       | /book/{id}   | Update book by id|
| DELETE    | /book/{id}   | Delete book by id|

### How to Run Code ?

1. Clone this Repo
2. Go to `task-6-book-api` branch
3. `cd Task-6` to switch directory
4. `go mod tidy` to add dependencies used in project
5. `go run main.go` to start server
6. It will start to serve on `localhost:PORT`

### Notes

First make .env file in root directory.

Put all your details in same format


```
USER        = "YOUR DATABASE USER_NAME"
PASSWORD    = "YOUR DATABASE PASSWORD"
HOST        = "YOUR DATABASE HOST_NAME"
DBNAME      = "YOUR DATABASE NAME"
SERVER_PORT = "YOUR SERVER_PORT"
```