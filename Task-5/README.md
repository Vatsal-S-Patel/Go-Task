# Go Tasks

## Task-5 : Web Server and Templates in Golang

### Instructions

1. Explore Web Server and Templates in Golang with `net/http` and `html/template` package.
2. Make a UserForm to make POST request on server, insert data in database and display all users data in a table.

### Solution

- Create a HTTP Server using `net/http` package.
- File Server to work with static files.
- Insert and fetch data into the database using `database/sql` package.
- Define routes and handle it with HandleFunc function of http and parse the data to HTML file using template.
- Display all users data in form of a table.

### How to Run Code ?

1. Clone this Repo
2. Go to `task-5` branch
3. `cd task-5` to switch directory
4. `go mod tidy` to add dependencies used in project
5. `go run server.go` to start server
6. It will start to serve on `localhost:8080`