package main

import (
	"log"

	"github.com/cdipaolo/sentiment"
)

func main() {

	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	var analysis *sentiment.Analysis
	var text string

	// Negative Example
	text = "Fat cash cow"
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}

	text = "It is windy outside"
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}

	// Positive Example
	text = "It is a beautiful day outside"
	analysis = model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		log.Printf("%s - Score of %d = Positive Sentiment\n", text, analysis.Score)
	} else {
		log.Printf("%s - Score of %d = Negative Sentiment\n", text, analysis.Score)
	}
}
