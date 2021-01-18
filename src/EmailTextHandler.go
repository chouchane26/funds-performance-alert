package main

import (
	"strconv"
	"strings"
	"time"
)

type Email struct {
	Subject string
	Text 	string
}

func generateStockPerformanceEmail(fundsQuotesPerformanceMap map[string][]Stock) Email{
	const letter = `Funds Performance for `

	currentDate := time.Now().Format("02 January 2006")

	textHeader := letter + currentDate + "\n" + "\n"

	text := ""

	for k, v := range fundsQuotesPerformanceMap{
		text += k + "\n"

		stockInfo := ""
		for _, stock := range v {
			stockInfo += stock.CompanyName + ": " + strconv.FormatFloat(stock.Percentage, 'f', 2, 64) + " %"+ "\n"
		}
		strings.Join(strings.Fields(stockInfo), " ")

		text += stockInfo + "\n"
	}

	message := textHeader + text

	newEmail := Email{
		Subject: "Your Performance for " + currentDate,
		Text: message,
	}

	return newEmail
}
