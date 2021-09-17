package main

import (
	"fmt"

	"github.com/sharpe1963/stringmod/quotes"
	"github.com/sharpe1963/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(quotes.GetEmoji("turtle"))
}
