# Go Tasks

## Task-7 : JWT in Go

In this task, REST APIs are made for profiles and task, by using that CRUD operations can perform for profiles and task.

In this task, i used gorm as ORM and mux as router to handle HTTP requests and `golang-jwt` for JWT generation and verification in Go.

### Instructions

Generate and Verify JWT tokens using middlewares in Go

### API endpoints

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| POST      | /profile        | Create profile   |
| GET       | /profile/{id}   | Get profile by id|
| PUT       | /profile/{id}   | Update profile by id|
| DELETE    | /profile/{id}   | Delete profile by id|
| GET    | /profile/task/{id}   | Get all task attached with profile id|

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| POST      | /task        | Create task   |
| GET       | /task/{id}   | Get task by id|
| PUT       | /task/{id}   | Update task by id|
| DELETE    | /task/{id}   | Delete task by id|

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| POST      | /login        | Generate JWT Token with email and password in JSON  |

### How to Run Code ?

1. Clone this Repo
2. Go to `task-7-jwt-auth` branch
3. `cd Task-7` to switch directory
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