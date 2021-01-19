package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type GlobalQuote struct {
	Symbol string `json:"01. symbol"`
	Price string `json:"05. price"`
	ChangePercentage  string `json:"10. change percent"`
}

type Stock struct {
	Symbol string
	CompanyName string
	Price float64
	Percentage float64
}

var requestNumber = 0

func loadFundsQuotesPerformance(fundsMap map[string][]FundStock) map[string][]Stock{
	quotesByFundName := make(map[string][]Stock)

	for _, value := range fundsMap{

		for _, stock := range value{
			quoteToday := findQuoteInfo(stock.Symbol)

			sleepRequest()

			price ,_ := strconv.ParseFloat(quoteToday.Price, 64)
			percentage, _ := strconv.ParseFloat(quoteToday.ChangePercentage[:len(quoteToday.ChangePercentage)-1], 64)

			newStock := Stock{
				Price: math.Round(price * 100)/100,
				Percentage: math.Round(percentage * 100)/100,
				CompanyName: stock.Company,
				Symbol: stock.Symbol,
			}

			quotesByFundName[stock.FundName] = append(quotesByFundName[stock.FundName], newStock)
		}
	}

	return quotesByFundName
}

func findQuoteInfo(symbol string) GlobalQuote{
	resp, err :=http.Get(getApiEndpoint(symbol))

	fmt.Printf(symbol, " ")

	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	globalQuote := make(map[string]GlobalQuote)
	json.Unmarshal(bodyBytes, &globalQuote)

	json.NewDecoder(resp.Body).Decode(globalQuote)

	return globalQuote["Global Quote"]
}

func getApiEndpoint(symbol string) string{
	endpoint := GlobalConfig.Alphavantage.Endpoint
	apikey := GlobalConfig.Alphavantage.Apikey

	return endpoint + "/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + apikey
}

func sleepRequest(){
	requestNumber += 1

	if requestNumber == 5 {
		requestNumber = 0
		time.Sleep(60 * time.Second)
	}
}
