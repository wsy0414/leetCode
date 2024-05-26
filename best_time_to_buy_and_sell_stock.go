package main

import "fmt"

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func main() {
  testCases := []struct{
    prices []int
    result int
  }{
    {
      prices: []int{7, 1, 5, 3, 6, 4},
      result: 5,
    }, {
      prices: []int{7, 6, 4, 3, 1},
      result: 0,
    },
  }

  for _, testCase := range testCases {
    fmt.Printf("expected result = %d and actual result = %d\n", testCase.result, otherMaxProfit(testCase.prices))
  }
}

func maxProfit(prices []int) int {
  profit := 0
  fmt.Printf("the profit = %d\n", profit)
  fmt.Printf("the prices = %v\n", prices)
  for i := 0; i< len(prices); i++{  
    for j := i + 1; j < len(prices); j++ {
      if prices[j] - prices[i] > profit {
        profit = prices[j] - prices[i]
      }
    }
  }

  return profit
}

func otherMaxProfit(prices []int) int {
  profit := 0
  minPrice := prices[0]

  for i := 1; i < len(prices); i++ {
    // if we find any price which is lower than current minPrice
    // update the minPrice
    if prices[i] < minPrice {
      minPrice = prices[i]
    } else if prices[i] - minPrice > profit {
      profit = prices[i] - minPrice
    }
  }

  return profit
}
