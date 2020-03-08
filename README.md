# Mux Mongo

A simple REST API with MongoDB made with [Mux HTTP Router](https://github.com/gorilla/mux) and [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

## Getting Started

Set your database URI and name in [database/database.go](database/database.go)

```go
func Connect() *mongo.Database {
    uri := "mongodb://localhost:27017"
    dbname := "dbname"
```

Create a new collection called **`users`**

You can import sample data to your users collection with [sample/db.json](sample/db.json)

Start the service with the `make` command

```sh
$ make
```

And there you go, the server is running on `localhost:8000`

To build the app run

```sh
$ make build
```

## Endpoints

| Method | Route       | Desc                                       |
| ------ | ----------- | ------------------------------------------ |
| GET    | /users      | Fetches all users in the users collection  |
| GET    | /users/{id} | Fetches a single user by the id            |
| POST   | /users      | Adds a new user to the collection          |
| PUT    | /users/{id} | Updates a user in the collection by the id |
| DELETE | /users/{id} | Deletes a user in the collection by the id |
