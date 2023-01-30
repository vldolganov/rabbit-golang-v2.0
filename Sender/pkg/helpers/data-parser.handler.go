package helpers

import (
	"encoding/json"
	"log"
	"rabbitv2/pkg/handlers"
	"rabbitv2/pkg/structures"
)

func DataParser(company string) []byte {
	var GlobalData structures.GlobalData
	data := handlers.StockData(company)
	err := json.Unmarshal(data, &GlobalData)
	parsedData, err := json.Marshal(&GlobalData)

	if err != nil {
		log.Fatal(err)
	}

	return parsedData
}
