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

## [simple web service](https://goinbigdata.com/how-to-create-a-simple-restful-service-in-go/)

GET http://localhost:8080/hello will return

Hello, Stranger

and GET http://localhost:8080/hello?name=Jack will result in

Hello, Jack

## meme-generator
From  git@github.com:montanaflynn/meme-generator.git 

To run

`go run main.go  --meme raptor --text "not sure if magick or golang" raptor.png
Saved to raptor.png`
