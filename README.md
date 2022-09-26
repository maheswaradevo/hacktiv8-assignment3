# Hacktiv8-Assignment 3

## Auto Reload Data

Automatically update the data from database every 15 minutes.

### How to run

```bash
go run cmd/main/main.go
```

### Run Database Migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/weather" up
```
