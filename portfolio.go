package main

import (
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func main() {
	// Set your Alpaca API credentials
	alpaca.SetBaseUrl("https://broker-api.sandbox.alpaca.markets

	") // Use paper trading endpoint
	alpaca.SetKeyID("your_alpaca_api_key_id")
	alpaca.SetSecretKey("your_alpaca_secret_key")

	// Initialize Alpaca client
	client := alpaca.NewClient(common.Credentials())

	// Example: Get account information
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal("Error fetching account:", err)
	}
	log.Printf("Account: %+v\n", account)

	// Example: Get stock information
	stockSymbol := "AAPL"
	asset, err := client.GetAsset(stockSymbol)
	if err != nil {
		log.Fatal("Error fetching asset:", err)
	}
	log.Printf("Asset: %+v\n", asset)

	// Example: Place a market order to buy 10 shares of AAPL
	qty := 10
	order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
		AccountID:   account.ID,
		Symbol:      stockSymbol,
		Qty:         qty,
		Side:        "buy",
		Type:        "market",
		TimeInForce: "gtc",
	})
	if err != nil {
		log.Fatal("Error placing order:", err)
	}
	log.Printf("Order placed: %+v\n", order)
}
