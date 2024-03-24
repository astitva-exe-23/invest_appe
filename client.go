package main

// StockClient struct for interacting with external stock API
type StockClient struct {
	baseURL string
}

func NewStockClient(baseURL string) *StockClient {
	return &StockClient{baseURL: baseURL}

}

type StockPriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

func (c *StockClient) GetStockPrice(symbol string) (float64, error) {
	url := fmt.Sprintf("%s/stocks/%s/price", c.BaseURL, symbol)

	// Make HTTP GET request to the external API
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful (HTTP status code 200)
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Decode the JSON response body into a StockPriceResponse struct
	var priceResp StockPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return 0, fmt.Errorf("error decoding response body: %v", err)
	}

	// Return the stock price
	return priceResp.Price, nil
}

func main() {
	// Example usage
	client := &StockClient{BaseURL: "https://broker-api.sandbox.alpaca.markets
	"}
	price, err := client.GetStockPrice("AAPL")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current stock price for AAPL:", price)
}
