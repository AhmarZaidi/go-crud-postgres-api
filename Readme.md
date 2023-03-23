Steps to run:

- Create a `.Env` file in the root of the project
- In the `.Env` file put:
    ```
    PORT=3000
    DB_URL="host=<host> user=<user> password=<password> dbname=<database name> port=<postgres server port> sslmode=disable"
    ```
- Run the postgres service
- From project root run: `go run main.go`
- Test using Postman (or any API platform)