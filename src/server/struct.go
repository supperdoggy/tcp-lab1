package main

type Aneks struct {
	A []Anek `json:"anek"`
}

type Anek struct {
	Id string `json:"id"`
	Cat string `json:"cat"`
	Text string `json:"text"`
}
