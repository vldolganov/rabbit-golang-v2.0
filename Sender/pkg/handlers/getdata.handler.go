package handlers

import (
	"io/ioutil"
	"net/http"
)

func StockData(company string) []byte {
	res, _ := http.Get("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + company + "&interval=5min&apikey=ISA1OC5LBZHCMGKI")
	responseData, _ := ioutil.ReadAll(res.Body)
	return responseData
}
