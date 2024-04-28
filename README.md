# 10xers-test
## Setting up Migrations
1. Install `golang-migrate` by running the following command:
```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Export an environment variable for convenience:
```shell
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/your_database?sslmode=disable'
```
Replace `password` and `your_database` with your actual database credentials.

3. Run the migrations based on what you need:

- To run migration up: `make migrate-up`
- To run migration down: `make migrate-down`

## Running on Local Machine
1. Create `.env` based on the `.env.example` file. This file should contain all the necessary environment variables for the application to run properly.

<!-- 2. Run the unit tests by executing the command `make test`. This will ensure that all the code is functioning as expected before running the server. -->

2. Run the server using the command `make run`. This will start the server and make it available on your specified port on localhost.

## API Documentation
You can find a Postman exported collection on the project root directory with the name `10xers.postman_collection.json` You can import this collection into Postman to easily test the API endpoints and see examples of the expected request and response formats.

You can also visit this [Postman](https://documenter.getpostman.com/view/24874313/2sA3BuVTtX) URL to view the API Documentation.

<!-- 
### Install Migrate
```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Setup Alias
```shell
alias air='~/go/bin/air' // or use "go run main.go"       
alias migrate='~/go/bin/migrate'
```

### Create Migration
```shell
migrate create -ext sql -dir db/migration table_name
```

### up migration
```shell
migrate -database "postgres://postgres:postgres@localhost:5432/10xers?sslmode=disable" -path db/migration up
```

### down migration
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/10xers?sslmode=disable" -path db/migration down
```

postman docs: https://documenter.getpostman.com/view/24874313/2sA3BuVTtX -->