# 1. Start database

- Start database container

```
docker-compose up db --detach
```

- Create simple_vocabulary database

```
make createdb
```

# 2. Migration

We need to install golang-migrate first to do the migration. On Mac, we can use:

```
brew install golang-migrate
```

### Create new migration file

```
migrate create -ext sql -dir database/migration -seq <file_name>
```

Ex: file_name = init_schema

### Run `up` migration file:

```
migrate -path database/migration/ -database "postgresql://<username>:<password>@localhost:5432/simple_vocabulary?sslmode=disable" -verbose up
```

- We can run instead:

```
make migrateUp
```

### Run `down` migration file:

```
migrate -path database/migration/ -database "postgresql://<username>:<password>@localhost:5432/simple_vocabulary?sslmode=disable" -verbose down
```

- We can run instead:

```
make migrateDown
```

- Please check Makefile to see and run the command easily

# 3. Start server

- Run in normal mode:

```
go run main.go
```

- To run with hot-reload, using:

```
nodemon --exec go run main.go --signal SIGTERM
```

- Run with docker (included hot-reload server):

```
docker-compose up
```
