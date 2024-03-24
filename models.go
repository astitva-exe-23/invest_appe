// models.go

package main

type Stock struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Transaction struct {
	ID        int
	StockID   int
	Type      string // "buy" or "sell"
	Quantity  int
	Timestamp string
}
