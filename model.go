package main

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

// Se va apela automat la initializarea packetuluyi curent "main".
func init() {
	items = []Item{
		{ID: "1", Name: "Item One"},
		{ID: "2", Name: "Item Two"},
		{ID: "3", Name: "Item Three"},
	}
}
