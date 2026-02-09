package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExchangeData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	// 1. Validate Argument Count
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: convert <amount> <currency_code>")
		fmt.Println("Example: ./convert 100 USD")
		return
	}

	// 2. Parse Amount
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Printf("Invalid amount '%s'. Please provide a number.\n", args[0])
		return
	}

	// 3. Get Exchange Data
	data, err := loadExchangeData("rates.json")
	if err != nil {
		fmt.Printf("System Error: %v\n", err)
		return
	}

	// 4. Validate and Convert
	target := strings.ToUpper(args[1])
	rate, exists := data.Rates[target]
	if !exists {
		fmt.Printf("Currency '%s' not found.\n", target)
		fmt.Printf("Available: %v\n", getKeys(data.Rates))
		return
	}

	// 5. Output Result
	result := amount * rate
	fmt.Println(result)
}

// loadExchangeData returns an error instead of panicking
func loadExchangeData(filename string) (ExchangeData, error) {
	var res ExchangeData
	content, err := os.ReadFile(filename)
	if err != nil {
		return res, fmt.Errorf("could not read file %s: %w", filename, err)
	}

	if err := json.Unmarshal(content, &res); err != nil {
		return res, fmt.Errorf("could not parse JSON: %w", err)
	}

	return res, nil
}

// Helper to show available currencies alphabetically
func getKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
