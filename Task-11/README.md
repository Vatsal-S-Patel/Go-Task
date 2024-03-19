# Go Tasks

## Task-11 : Trains Task in Mux and MongoDB

In this task, i used `gorilla/mux` as router and MongoDB as database.

### Instructions

1. Read CSV file and Dump all that data to MongoDB collection. To do that use flags like `go run main.go -readcsv -csv={filename}`
2. Create APIs for train that include pagination.
3. Also make request using search keywords that contains it in Train no, Name, Source, Destination.


### API endpoints

| Method    | endpoint     | Description   |
| --------- | --------     | ------------- |
| GET       | /api/trains       | Get all trains |
| GET       | /api/trains/?page={page}       | Get all trains with pagination |
| GET       | /api/trains/?search={keywords}       | Get all trains contains keywords in their no, name, source, destination |
| GET       | /api/trains/?search={keywords}&page={page}       | Get all trains contains keywords in their no, name, source, destination with pagination |

### How to Run Code ?

1. Clone this Repo
2. Go to `task-11-train-mongo` branch
3. `cd Task-11/backend` to switch directory
4. `go mod tidy` to add dependencies used in project
5. `go run main.go` to start server
6. It will start to serve on `localhost:PORT`

### Notes
To insert CSV file data into MongoDB run following command in backend directory.

`go run main.go  -readcsv -csv={csvFileName}`

First make .env file in root directory.

Put all your details in same format

```
MONGO_URI   = "YOUR MONGODB CONNECTION URI"
DBNAME      = "YOUR DATABASE NAME"
SERVER_PORT = "YOUR SERVER_PORT"
```