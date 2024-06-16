package main

import (
	"bufio"
	"fmt"
	"os"
)

type ExchangeRates struct {
	rub_to_dollar  [2]float64
	rub_to_euro    [2]float64
	dollar_to_rub  [2]float64
	dollar_to_euro [2]float64
	euro_to_rub    [2]float64
	euro_to_dollar [2]float64
}

type Valutes struct {
	rubles  float64
	dollars float64
	euros   float64
}

func maxProfit(rates [3]ExchangeRates) float64 {
	dp := make([][]Valutes, 3)
	dp[0] = make([]Valutes, 3)
	dp[1] = make([]Valutes, 3)
	dp[2] = make([]Valutes, 3)
	dp[0][0] = Valutes{1, rates[0].rub_to_dollar[1] / rates[0].rub_to_dollar[0], rates[0].rub_to_euro[1] / rates[0].rub_to_euro[0]}
	dp[1][0] = Valutes{1, rates[1].rub_to_dollar[1] / rates[1].rub_to_dollar[0], rates[1].rub_to_euro[1] / rates[1].rub_to_euro[0]}
	dp[2][0] = Valutes{1, rates[2].rub_to_dollar[1] / rates[2].rub_to_dollar[0], rates[2].rub_to_euro[1] / rates[2].rub_to_euro[0]}

	dp[1][1].dollars = max(dp[0][0].dollars, dp[0][0].euros*rates[1].dollar_to_euro[1]/rates[1].dollar_to_euro[0])
	dp[1][1].euros = max(dp[0][0].euros, dp[0][0].dollars*rates[1].euro_to_dollar[1]/rates[1].euro_to_dollar[0])
	dp[1][1].rubles = max(dp[0][0].rubles, dp[0][0].euros*rates[1].rub_to_euro[1]/rates[1].rub_to_euro[0])
	dp[1][1].rubles = max(dp[1][1].rubles, dp[0][0].dollars*rates[1].rub_to_dollar[1]/rates[1].rub_to_dollar[0])

	dp[2][2].dollars = max(dp[1][1].dollars, dp[1][1].euros*rates[2].dollar_to_euro[1]/rates[2].dollar_to_euro[0])
	dp[2][2].dollars = max(dp[2][2].dollars, dp[1][1].rubles*rates[2].dollar_to_rub[1]/rates[2].dollar_to_rub[0])

	dp[0][1].dollars = max(dp[1][0].dollars, dp[1][0].euros*rates[0].dollar_to_euro[1]/rates[0].dollar_to_euro[0])
	dp[0][1].euros = max(dp[1][0].euros, dp[1][0].dollars*rates[0].euro_to_dollar[1]/rates[0].euro_to_dollar[0])
	dp[0][1].rubles = max(dp[1][0].rubles, dp[1][0].euros*rates[0].rub_to_euro[1]/rates[0].rub_to_euro[0])
	dp[0][1].rubles = max(dp[0][1].rubles, dp[1][0].dollars*rates[0].rub_to_dollar[1]/rates[0].rub_to_dollar[0])

	dp[2][2].dollars = max(dp[2][2].dollars, dp[0][1].dollars)
	dp[2][2].dollars = max(dp[2][2].dollars, dp[0][1].euros*rates[2].dollar_to_euro[1]/rates[2].dollar_to_euro[0])
	dp[2][2].dollars = max(dp[2][2].dollars, dp[0][1].rubles*rates[2].dollar_to_rub[1]/rates[2].dollar_to_rub[0])

	dp[2][1].dollars = max(dp[1][0].dollars, dp[1][0].euros*rates[2].dollar_to_euro[1]/rates[2].dollar_to_euro[0])
	dp[2][1].euros = max(dp[1][0].euros, dp[1][0].dollars*rates[2].euro_to_dollar[1]/rates[2].euro_to_dollar[0])
	dp[2][1].rubles = max(dp[1][0].rubles, dp[1][0].euros*rates[2].rub_to_euro[1]/rates[2].rub_to_euro[0])
	dp[2][1].rubles = max(dp[2][1].rubles, dp[1][0].dollars*rates[2].rub_to_dollar[1]/rates[2].rub_to_dollar[0])

	dp[0][2].dollars = max(dp[0][2].dollars, dp[2][1].euros*rates[0].dollar_to_euro[1]/rates[0].dollar_to_euro[0])
	dp[0][2].dollars = max(dp[0][2].dollars, dp[2][1].rubles*rates[0].dollar_to_rub[1]/rates[0].dollar_to_rub[0])

	dp[1][1].dollars = max(dp[2][0].dollars, dp[0][0].euros*rates[1].dollar_to_euro[1]/rates[1].dollar_to_euro[0])
	dp[1][1].euros = max(dp[2][0].euros, dp[0][0].dollars*rates[1].euro_to_dollar[1]/rates[1].euro_to_dollar[0])
	dp[1][1].rubles = max(dp[2][0].rubles, dp[0][0].euros*rates[1].rub_to_euro[1]/rates[1].rub_to_euro[0])
	dp[1][1].rubles = max(dp[1][1].rubles, dp[0][0].dollars*rates[1].rub_to_dollar[1]/rates[1].rub_to_dollar[0])

	dp[0][2].dollars = max(dp[0][2].dollars, dp[1][1].dollars)
	dp[0][2].dollars = max(dp[0][2].dollars, dp[1][1].euros*rates[0].dollar_to_euro[1]/rates[0].dollar_to_euro[0])
	dp[0][2].dollars = max(dp[0][2].dollars, dp[1][1].rubles*rates[0].dollar_to_rub[1]/rates[0].dollar_to_rub[0])

	// dp[0][1].dollars = max(dp[1][0].dollars, dp[1][0].euros*rates[0].dollar_to_euro[1]/rates[0].dollar_to_euro[0])
	// dp[0][1].euros = max(dp[1][0].euros, dp[1][0].dollars*rates[0].euro_to_dollar[1]/rates[0].euro_to_dollar[0])
	// dp[0][1].rubles = max(dp[1][0].rubles, dp[1][0].euros*rates[0].rub_to_euro[1]/rates[0].rub_to_euro[0])
	// dp[0][1].rubles = max(dp[0][1].rubles, dp[1][0].dollars*rates[0].rub_to_dollar[1]/rates[0].rub_to_dollar[0])

	dp[1][2].dollars = max(dp[0][1].dollars, dp[0][1].euros*rates[1].dollar_to_euro[1]/rates[1].dollar_to_euro[0])
	dp[1][2].dollars = max(dp[1][2].dollars, dp[0][1].rubles*rates[1].dollar_to_rub[1]/rates[1].dollar_to_rub[0])

	dp[1][2].dollars = max(dp[1][2].dollars, dp[2][1].dollars)
	dp[1][2].dollars = max(dp[1][2].dollars, dp[2][1].euros*rates[1].dollar_to_euro[1]/rates[1].dollar_to_euro[0])
	dp[1][2].dollars = max(dp[1][2].dollars, dp[2][1].rubles*rates[1].dollar_to_rub[1]/rates[1].dollar_to_rub[0])

	maxDollars := max(dp[0][2].dollars, dp[1][2].dollars)
	return max(dp[2][2].dollars, maxDollars)
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
			// for k := 0; k < 6; k++ {
			fmt.Fscan(in, &rates[j].rub_to_dollar[0], &rates[j].rub_to_dollar[1])
			fmt.Fscan(in, &rates[j].rub_to_euro[0], &rates[j].rub_to_euro[1])
			fmt.Fscan(in, &rates[j].dollar_to_rub[0], &rates[j].dollar_to_rub[1])
			fmt.Fscan(in, &rates[j].dollar_to_euro[0], &rates[j].dollar_to_euro[1])
			fmt.Fscan(in, &rates[j].euro_to_rub[0], &rates[j].euro_to_rub[1])
			fmt.Fscan(in, &rates[j].euro_to_dollar[0], &rates[j].euro_to_dollar[1])
			// }
		}
		maxDollars := maxProfit(rates)

		fmt.Fprintf(out, "%.6f\n", maxDollars)
	}
}
