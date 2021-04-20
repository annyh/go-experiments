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

## download_remote_file
go run main.go

TODO
- mass download pictures, given a excel sheet. can use pictures for collage

## pdf_generator
go run main.go

TODO: 
- refactor to produce certificates
- convert PDF to png/jpeg

## web_scraper
go run main.go

logs out the blog post titles. Sort of like RSS

## screenshots
go run main.go

TODO: change to take URL parameter from the command line

## sentiment_analysis
go run main.go

given a string the result is either 0 for negative or 1 for positive. there is no in-between :( Wish machines detected sarcasm; that might be too hard. It's too hard for even some humans.
