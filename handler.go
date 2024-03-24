package main

import (
	"encoding/json"
	"net/http"

	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func handleGetAsset(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")

	keyID := os.Getenv("ALPACA_API_KEY_ID")
	secretKey := os.Getenv("ALPACA_SECRET_KEY")
	alpaca.SetBaseUrl("https://broker-api.sandbox.alpaca.markets

	") 

	client := alpaca.NewClient(alpaca.APIKey{ID: keyID, Secret: secretKey})

	asset, err := client.GetAsset(symbol)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching asset")
		return
	}

	respondWithJSON(w, http.StatusOK, asset)
}

func handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Extract order details from request body (assuming it's JSON)
	type OrderRequest struct {
		Symbol      string  `json:"symbol"`
		Qty         int     `json:"qty"`
		Side        string  `json:"side"` // "buy" or "sell"
		Type        string  `json:"type"` // "market", "limit", etc.
		Limit       float64 `json:"limit,omitempty"`
		Stop        float64 `json:"stop,omitempty"`
		TimeInForce string  `json:"time_in_force,omitempty"` // "day", "gtc", etc.
	}

	var orderReq OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	keyID := os.Getenv("ALPACA_API_KEY_ID")
	secretKey := os.Getenv("ALPACA_SECRET_KEY")
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets") // Use paper trading endpoint

	client := alpaca.NewClient(alpaca.APIKey{ID: keyID, Secret: secretKey})

	// Construct the order object
	order := alpaca.PlaceOrderRequest{
		Symbol:      orderReq.Symbol,
		Qty:         orderReq.Qty,
		Side:        orderReq.Side,
		Type:        orderReq.Type,
		LimitPrice:  orderReq.Limit,
		StopPrice:   orderReq.Stop,
		TimeInForce: orderReq.TimeInForce,
	}

	// Place the order
	resp, err := client.PlaceOrder(order)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error placing order")
		return
	}

	// Respond with the order response
	respondWithJSON(w, http.StatusOK, resp)
}
