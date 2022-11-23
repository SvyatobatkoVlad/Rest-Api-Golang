# Rest-Api-Golang

1) go-gin
2) viper, env
3) postgresql, sqlx,

### Run Postgresql in Docker
``docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres``

### Run migration 
``migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up ``
``migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down ``


### Connect to db
``docker exec -it  {container_id} /bin/bash``


###schema db (Postgres)
![](images/schema_db_postgres.png)