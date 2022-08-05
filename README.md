# Halo Lab
   
### Depends:
        Redis
        Postgres

#### For building the project, you can use the following commands:

```bash
make build
```

##### OR

```bash
go build -o halo cmd/test/main.go
```

#### Setup environment variables:

```bash
export PG_USER=postgres
export PG_PASSWORD=SecretPassword
export PG_DB=localhost
export PG_PORT=5432
export PG_HOST=dvdrental
export URL=http://localhost
export PORT=8080
export REDIS=redis://localhost:6379

```
#### Import database from link
```bash
https://www.postgresqltutorial.com/wp-content/uploads/2019/05/dvdrental.zip
```

#### Run the project:

```bash
make run
```
#### or 
```bash
 go run cmd/test/main.go
```
