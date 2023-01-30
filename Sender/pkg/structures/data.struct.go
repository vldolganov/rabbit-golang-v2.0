package structures

type GlobalData struct {
	Stock Stock `json:"Global Quote"`
}

type Stock struct {
	Symbol string `json:"01. symbol"`
	Open   string `json:"02. open"`
	High   string `json:"03. high"`
	Low    string `json:"04. low"`
	Day    string `json:"07. latest trading day"`
}
