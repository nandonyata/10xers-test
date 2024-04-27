# 10xers-test

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
migrate -database "postgres://postgres:postgres@localhost:5432/booking-to-go?sslmode=disable" -path db/migration up
```

### down migration
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/booking-to-go?sslmode=disable" -path db/migration down
```