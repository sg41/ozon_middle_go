package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type ExchangeRates struct {
	rubToDollar  [2]float64
	rubToEuro    [2]float64
	dollarToRub  [2]float64
	dollarToEuro [2]float64
	euroToRub    [2]float64
	euroToDollar [2]float64
}

type Valutes struct {
	rubles  float64
	dollars float64
	euros   float64
}

func maxProfit(rates [3]ExchangeRates, numExchanges int) float64 {
	dp := make([][][]Valutes, numExchanges+1)
	for i := range dp {
		dp[i] = make([][]Valutes, 3)
		for j := range dp[i] {
			dp[i][j] = make([]Valutes, 3)
			for k := range dp[i][j] {
				// Initialize with negative values to indicate not calculated
				dp[i][j][k] = Valutes{-1, -1, -1}
			}
		}
	}

	// Base case: Starting with 1 ruble
	dp[0][0][0] = Valutes{1, 0, 0}

	for i := 1; i <= numExchanges; i++ { // Exchanges
		for j := 0; j < 3; j++ { // Current currency
			for k := 0; k < 3; k++ { // Previous bank
				for b := 0; b < 3; b++ { // Next bank
					if b == k {
						continue // Skip if using the same bank consecutively
					}
					if dp[i-1][j][k].rubles < 0 {
						continue // Skip if previous amount is not calculated (no valid path)
					}

					// Calculate possible exchange amounts
					nextValutes := calculateExchangeAmount(dp[i-1][j][k], j, b, rates)

					// Update dp table if the exchange is beneficial
					if nextValutes.dollars > dp[i][1][b].dollars {
						dp[i][1][b] = nextValutes
					}
				}
			}
		}
	}

	// Find the maximum dollars achievable
	maxDollars := 0.0
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			maxDollars = math.Max(maxDollars, dp[numExchanges][j][k].dollars)
		}
	}
	return maxDollars
}

func calculateExchangeAmount(currentValutes Valutes, currentCurrency, bank int, rates [3]ExchangeRates) Valutes {
	nextValutes := Valutes{0, 0, 0}
	switch currentCurrency {
	case 0: // Rubles
		nextValutes.dollars = currentValutes.rubles * rates[bank].rubToDollar[0] / rates[bank].rubToDollar[1]
		nextValutes.euros = currentValutes.rubles * rates[bank].rubToEuro[0] / rates[bank].rubToEuro[1]
	case 1: // Dollars
		nextValutes.rubles = currentValutes.dollars * rates[bank].dollarToRub[0] / rates[bank].dollarToRub[1]
		nextValutes.euros = currentValutes.dollars * rates[bank].dollarToEuro[0] / rates[bank].dollarToEuro[1]
	case 2: // Euros
		nextValutes.rubles = currentValutes.euros * rates[bank].euroToRub[0] / rates[bank].euroToRub[1]
		nextValutes.dollars = currentValutes.euros * rates[bank].euroToDollar[0] / rates[bank].euroToDollar[1]
	}
	return nextValutes
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var rates [3]ExchangeRates
		for j := 0; j < 3; j++ {
			fmt.Fscan(in, &rates[j].rubToDollar[0], &rates[j].rubToDollar[1])
			fmt.Fscan(in, &rates[j].rubToEuro[0], &rates[j].rubToEuro[1])
			fmt.Fscan(in, &rates[j].dollarToRub[0], &rates[j].dollarToRub[1])
			fmt.Fscan(in, &rates[j].dollarToEuro[0], &rates[j].dollarToEuro[1])
			fmt.Fscan(in, &rates[j].euroToRub[0], &rates[j].euroToRub[1])
			fmt.Fscan(in, &rates[j].euroToDollar[0], &rates[j].euroToDollar[1])
		}

		numExchanges := 3 // Set the number of exchanges here (can be adjusted)
		maxDollars := maxProfit(rates, numExchanges)
		fmt.Fprintf(out, "%.6f\n", maxDollars)
	}
}
