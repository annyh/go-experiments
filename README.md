## hangman

`cd hangman`

Run with 

`go run main.go`

Run tests

`go test`

Generate HTML coverage report

```bash
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```