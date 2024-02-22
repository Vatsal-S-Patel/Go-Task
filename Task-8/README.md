# Go Tasks

## Task-8 : UserAPIs using Fiber and MongoDB 

In this task, REST APIs are made for user, by using that CRUD operations can perform for users.

In this task, i used Fiber as web framework and MongoDB as database.

### Instructions

1. For given user struct, create REST APIs to perform CRUD operations for user.
2. Create modular project structure for easy to maintain and readable code.
2. Use Fiber as web framework and MongoDB as database.

### API endpoints

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| GET       | /api/user       | Get all users |
| POST      | /api/user        | Create user   |
| GET       | /api/user/{id}   | Get one user by id|
| PUT       | /api/user/{id}   | Update user by id|
| DELETE    | /api/user/{id}   | Delete user by id|

### How to Run Code ?

1. Clone this Repo
2. Go to `task-8-fiber-mongo` branch
3. `cd Task-8` to switch directory
4. `go mod tidy` to add dependencies used in project
5. `go run main.go` to start server
6. It will start to serve on `localhost:PORT`

### Notes

First make .env file in root directory.

Put all your details in same format

```
MONGO_URI   = "YOUR MONGODB CONNECTION URI"
DBNAME      = "YOUR DATABASE NAME"
SERVER_PORT = "YOUR SERVER_PORT"
```